package app

type AppService struct {
}

var appServiceInstance *AppService

func NewService() *AppService {
	if appServiceInstance == nil {
		appServiceInstance = &AppService{}
	}
	return appServiceInstance
}

func (appService *AppService) Ping() (string, error) {
	return "pong", nil
}
