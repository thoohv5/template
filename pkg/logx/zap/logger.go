package zap

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"

	los "github.com/thoohv5/template/pkg/logx/standard"
	"github.com/thoohv5/template/pkg/util"
)

const (
	OutSep = ","
)

type log struct {
	zlogger *zap.Logger
}

func New(config *los.Config) los.ILogger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	ws := make([]zapcore.WriteSyncer, 0)
	for _, out := range strings.Split(config.Out, OutSep) {
		switch out {
		case "std":
			ws = append(ws, zapcore.AddSync(os.Stdout))
		default:
			if strings.HasPrefix(out, ".") {
				out = util.AbPath(out)
			}
			hook := lumberjack.Logger{
				Filename:   out,  // 日志文件路径
				MaxSize:    128,  // 每个日志文件保存的最大尺寸 单位：M
				MaxBackups: 30,   // 日志文件最多保存多少个备份
				MaxAge:     7,    // 文件最多保存多少天
				Compress:   true, // 是否压缩
			}
			ws = append(ws, zapcore.AddSync(&hook))
		}
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(ws...),
		zap.NewAtomicLevelAt(parseLevel(config.Level)),
	)

	caller := zap.AddCaller()

	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	// log := zap.New(core, caller, zap.Development(), filed)
	logger := zap.New(core, caller, zap.Development())

	return &log{
		zlogger: logger,
	}
}

func (log *log) Debugc(ctx context.Context, msg string, values ...interface{}) {
	log.zlogger.Debug(fmt.Sprintf(msg, values...))
}

func (log *log) Infoc(ctx context.Context, msg string, values ...interface{}) {
	log.zlogger.Info(fmt.Sprintf(msg, values...))
}

func (log *log) Warnc(ctx context.Context, msg string, values ...interface{}) {
	log.zlogger.Warn(fmt.Sprintf(msg, values...))
}

func (log *log) Errorc(ctx context.Context, msg string, values ...interface{}) {
	log.zlogger.Error(fmt.Sprintf(msg, values...))
}

func (log *log) Debug(msg string, values ...interface{}) {
	log.Debugc(context.Background(), msg, values...)
}

func (log *log) Info(msg string, values ...interface{}) {
	log.Infoc(context.Background(), msg, values...)
}

func (log *log) Warn(msg string, values ...interface{}) {
	log.Warnc(context.Background(), msg, values...)
}

func (log *log) Error(msg string, values ...interface{}) {
	log.Errorc(context.Background(), msg, values...)
}

// 日志类别: debug, warn, info，error
func parseLevel(level string) zapcore.Level {
	zl := zapcore.DebugLevel
	switch level {
	case "debug":
		zl = zapcore.DebugLevel
	case "warn":
		zl = zapcore.WarnLevel
	case "info":
		zl = zapcore.InfoLevel
	case "error":
		zl = zapcore.ErrorLevel
	}
	return zl
}
