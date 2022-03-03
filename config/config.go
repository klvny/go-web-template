package config

import "github.com/spf13/viper"

type AppConfig struct {
	App struct {
		Port string `json:"port"`
	} `json:"app"`
	Database struct {
		File string `json:"file"`
	} `json:"database"`
}

func (c *AppConfig) Bind(v *viper.Viper) {
	c.App.Port = v.GetString("app.port")
	c.Database.File = v.GetString("database.file")
}
