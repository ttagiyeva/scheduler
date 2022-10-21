package log

import (
	"context"

	"github.com/ttagiyeva/scheduler/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//NewZapLogger constructs a logger instance
func NewZapLogger(lc fx.Lifecycle, conf *config.Config) *zap.SugaredLogger {
	cfg := zap.Config{
		Encoding:    conf.LoggerConfig.Encoding,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "levelKey",
			MessageKey:  "messageKey",
			FunctionKey: "functionKey",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"app": "scheduler"},
	}

	cfg.Level.UnmarshalText([]byte(conf.LoggerConfig.Level))

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return logger.Sync()
		},
	})

	return logger.Sugar()
}
