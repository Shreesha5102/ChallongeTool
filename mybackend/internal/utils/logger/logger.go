package logger

import (
	"fmt"
	"os"

	"github.com/Shreesha5102/ChallongeTool/mybackend/internal/utils/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger() *zap.SugaredLogger {
	// Create zap config
	logOutputPath := os.Getenv("LOG_FILE")
	if logOutputPath == "" {
		logOutputPath = constants.LogOutputPath
	}

	config := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			FunctionKey: "function",
		},
		OutputPaths:      []string{logOutputPath},
		ErrorOutputPaths: []string{"stderr", logOutputPath},
	}

	loggerBuild, err := config.Build()

	if err != nil {
		fmt.Printf("Error building logger: %v", err)
		return zap.NewExample().Sugar()
	}

	defer func() {
		if err := loggerBuild.Sync(); err != nil {
			// Handle error, perhaps with another logger or stdlib
			fmt.Printf("Logger Sync failed: %v", err)
		}
	}()

	logger := loggerBuild.Sugar()
	return logger
}

func GetLogger() zap.SugaredLogger {
	return *createLogger()
}
