package ginrest

import (
	"go-clean-arch/config"
	appRest "go-clean-arch/core/app/delivery/rest"
	appUcase "go-clean-arch/core/app/usecases"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIService struct{}

func NewAPIService() *APIService {
	return &APIService{}
}

// Start api service init and start
func (api *APIService) Start(mongoClient *mongo.Client) {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// Force log's color
	gin.ForceConsoleColor()
	// Disable log's color
	// gin.DisableConsoleColor()

	/* usecase, delivery 注入router */
	appUsecase := appUcase.NewAppUsecase()

	appRest.NewAppHandler(r, appUsecase)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    config.Cfgs.Port,
		Handler: r,
	}

	api.gracefulShutdown(server)
	glog.Infof("Start API Service: 127.0.0.1%s", config.Cfgs.Port)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			glog.Info("Server closed under request")
		} else {
			glog.Fatalf("Server closed unexpect: %v", err)
		}
	}
	glog.Info("Server exiting")
	os.Exit(1)
}

func (a *APIService) gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		glog.Info("receive interrupt signal")
		if err := server.Close(); err != nil {
			glog.Fatal("Server Close:", err)
		}
	}()
}
