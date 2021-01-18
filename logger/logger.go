package logger

import(
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	log *zap.Logger
)

func GetLogger() (*zap.Logger){
	return log
}


func init() {

	logConfig:=zap.Config{
		OutputPaths: []string{"stdout"},
		Level: zap.NewAtomicLevelAt(getLevel()),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err!= nil{
		panic(err);
	}
}

func getLevel() zapcore.Level{

	if os.Getenv("LOG_LEVEL") == "INFO" {
		return zapcore.InfoLevel
	}
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		return zapcore.DebugLevel
	}
	if os.Getenv("LOG_LEVEL") == "ERROR" {
		return zapcore.ErrorLevel
	}
	return zapcore.InfoLevel
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
