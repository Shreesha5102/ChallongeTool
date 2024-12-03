package app

import (
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/configs"
	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	log    = logger.GetLogger()
)

func startRestServer() {
	mapRestUrls()
	log.Info("Starting REST server on port: %s", configs.GetHTTPPort())
	err := router.Run(configs.GetHTTPPort())
	if err != nil {
		log.Error("Server failed to start due to %+v", err)
		return
	}
}

func StartApplication() {
	log.Infof("Starting gin server")
	startRestServer()
}
