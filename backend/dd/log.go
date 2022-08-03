package dd

import (
	"os"
	"time"

	"animeii.tech/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	level zapcore.Level
	opts  []zap.Option
)

func InitLog() {
	// 创建目录
	createRootDir()
	// 设置日志等级
	setLogLevel()
	if Cf.Log.ShowLine {
		opts = append(opts, zap.AddCaller())
	}
	Log = zap.New(getZapCore(), opts...)

}

// 创建目录
func createRootDir() {
	if ok, _ := utils.PathExists(Cf.Log.RootDir); !ok {
		_ = os.Mkdir(Cf.Log.RootDir, os.ModePerm)
	}
}

// 设置日志等级
func setLogLevel() {
	switch Cf.Log.Level {
	case "debug":
		level = zap.DebugLevel
		opts = append(opts, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		opts = append(opts, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}

	// 格式化json
	if Cf.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewCore(encoder, getLogWriter(), level)
}

// 写入文件
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   Cf.Log.RootDir + "/" + Cf.Log.Filename,
		MaxSize:    Cf.Log.MaxSize,
		MaxBackups: Cf.Log.MaxBackups,
		MaxAge:     Cf.Log.MaxAge,
		Compress:   Cf.Log.Compress,
	}

	return zapcore.AddSync(file)
}
