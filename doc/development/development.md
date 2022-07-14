## 開發主要使用套件

開發階段相關指令如下：

```sh
go get -u github.com/gin-gonic/gin

go install github.com/cosmtrek/air
air -d

# gin-swagger
# https://github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag

swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

# open: http://127.0.0.1:5003/swagger/index.html#

# configs
go get github.com/spf13/pflag
go get github.com/spf13/viper
go get github.com/golang/glog

# db
go get go.mongodb.org/mongo-driver/mongo

# go get go get github.com/dgrijalva/jwt-go
# go get golang.org/x/oauth2

# https://github.com/satori/go.uuid
go get github.com/satori/go.uuid

```


```
# run unit tests
go test ./...
```