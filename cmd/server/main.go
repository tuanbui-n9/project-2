package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	config "github.com/tuanbui-n9/project-2/configs"
	router "github.com/tuanbui-n9/project-2/internal/app/routes"
	"github.com/tuanbui-n9/project-2/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := database.Connect()
	r := router.SetupRouter(db)
	trustedProxies := strings.Split(cfg.TRUSTED_PROXIES, ",")
	r.SetTrustedProxies(trustedProxies)

	r.Run(":" + cfg.PORT)
}
