package main

import (
	"app/core"
	"app/handlers"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {

	appEnv := os.Getenv("APP_ENV")

	path, _ := os.Getwd()
	configurator := core.Configurator{}

	if appEnv == "" {
		_ = os.Setenv("APP_ENV", "dev")
		appEnv = os.Getenv("APP_ENV")
		configurator.Read(path+"/app/config/main."+appEnv+".json", &core.App.Config)
	}else{
		configurator.Read(path+"/config/main."+appEnv+".json", &core.App.Config)
	}




	//core.App.Pgx = core.NewPgx(core.App.Config["dsn"])
	core.App.Gorm = core.NewGorm(core.App.Config["dsn"])
}

func main() {

	r := gin.Default()

	r.GET("/api/reverse", handlers.ReverseHandler)

	_ = r.Run(":" + core.App.Config["server_port"])
}
