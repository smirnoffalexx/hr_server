package app

import (
	"fmt"
	"hr-server/config"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func FormatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func InitLogger(cfg *config.Config) error {
	logrus.SetReportCaller(true)

	logrusLvl, err := logrus.ParseLevel(cfg.Logger.LOGLVL)
	if err != nil {
		return fmt.Errorf("can't load \"LOGL\": %s", err)
	}

	txtFormatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		ForceColors:            true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", FormatFilePath(f.File), f.Line)
		},
	}

	logrus.SetLevel(logrusLvl)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(txtFormatter)

	return nil
}
