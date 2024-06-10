package config

type App struct {
	Host   string `yml:"host"`
	Port   string `yml:"port"`
	Secret string `yml:"secret"`
}
