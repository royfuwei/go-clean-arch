package usecase

import "go-clean-arch/domain"

type appUsecase struct{}

func NewAppUsecase() domain.AppUsecase {
	return &appUsecase{}
}

func (a *appUsecase) GetApp() *domain.App {
	return &domain.App{
		AppName: "go-clean-arch",
	}
}
