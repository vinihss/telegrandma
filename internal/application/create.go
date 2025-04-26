package application

func NewApplication(name string, description string) *Application {

	var config, err = getDefaultConfig()
	var application = &Application{
		Name:        name,
		Description: description,
	}
	if err == nil {
		application.AppConfig = *config
	}
	return application

}
