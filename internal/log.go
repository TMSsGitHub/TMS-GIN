package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Logger *zap.Logger

func getLogFilePath(logDir, level string) string {
	now := time.Now()

	dir := now.Format("2006/01/02")
	filename := fmt.Sprintf("%s_%s.log", dir, level)

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return ""
	}
	return filepath.Join(logDir, filename)
}

func newCustomLogWriter(logDir, level string) (zapcore.WriteSyncer, error) {
	//lumberjackLogger := &lumberjack.Logger{
	//	Filename:   getLogFilePath(logDir, level),
	//	Compress:   true,
	//	MaxAge:     28,
	//	MaxSize:    100,
	//	MaxBackups: 10,
	//	LocalTime:  true,
	//}
	//return zapcore.AddSync(lumberjackLogger), nil

	writer, err := rotatelogs.New(
		getLogFilePath(logDir, level),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return zapcore.AddSync(writer), err
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	infoFileWriteSyncer, err := newCustomLogWriter("./logs", "info")
	errorFileWriteSyncer, err := newCustomLogWriter("./logs", "error")

	// 仅Info级别日志处理器

	infoLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.InfoLevel
	})
	if err != nil {
		return nil, err
	}
	infoCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), infoFileWriteSyncer), // 同时输出到控制台与文件中
		infoLevelEnabler,
	)
	errorCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), errorFileWriteSyncer),
		zap.ErrorLevel,
	)
	core := zapcore.NewTee(infoCore, errorCore)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)), nil
}

func InitLogger() {
	var err error
	Logger, err = newZapLogger()
	if err != nil {
		panic(err)
	}
}

//
//func ZapMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//		path := c.Request.URL.Path
//		c.Next()
//		latency := time.Since(start)
//		clientIP := c.ClientIP()
//		method := c.Request.Method
//		statusCode := c.Writer.Status()
//		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
//
//		Logger.Info(
//			path,
//			zap.Int("status", statusCode),
//			zap.String("path", path),
//			zap.String("ip", clientIP),
//			zap.String("method", method),
//			zap.String("errorMessage", errorMessage),
//			zap.Duration("latency", latency),
//		)
//	}
//}
