package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	ZapCfg    *zap.Config
	ZapLogger *zap.Logger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) SetZapConfig(zCfg *zap.Config) {
	l.ZapCfg = zCfg
}

func (l *Logger) SetZapLogger(zLogger *zap.Logger) {
	l.ZapLogger = zLogger
}

func (l *Logger) GetLogger() *zap.Logger {
	logFile, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(logFile),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	logger := zap.New(core)
	// defer logger.Sync()
	// defer logFile.Close()
	logger.Sync()
	return logger
}

// type ZapLogger *zap.Logger

// var zapLogger *zap.Logger

// func (z ZapLogger) Info() string {}

// func GetLogger(name string) *zap.Logger {
// 	if zapLogger == nil {
// 		config := zap.Config{
// 			Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel), // Mức độ log
// 			Encoding:         "json",                                  // Định dạng log
// 			OutputPaths:      []string{"stdout"},                      // Đầu ra của log
// 			ErrorOutputPaths: []string{"stderr"},                      // Đầu ra lỗi
// 		}
// 		logger, err := config.Build()
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer logger.Sync() // Đảm bảo log được ghi trước khi chương trình kết thúc
// 		zapLogger = logger
// 	}

// 	return zapLogger
// }
