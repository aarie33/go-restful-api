package simple

// Struct Field Provider
type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{
			Name: "Simple Bray",
		},
	}
}
