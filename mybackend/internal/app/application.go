package app

var (
	router = gin.Default()
	logger = logg.GetLogger()
)

func startRestServer() {
	mapRestUrls()
	logger.Info("Starting REST server")
	err := router.RunTLS(":"+config.GetHTTPSPort(), constant.CertFile, constant.KeyFile)
	if err != nil {
		logger.Error("Server failed to start due to %+v", err)
		return
	}
}
