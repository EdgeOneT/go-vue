// log日志

package log

import (
	"github.com/sirupsen/logrus"
	"go-vue/common/config"
	"os"
	"path/filepath"
)

var log *logrus.Logger

var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

// 初始化
func init() {
	setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
}

// Log 方法调用
func Log() *logrus.Logger {
	// 文件输出
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		// 控制台输出
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2008"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}
