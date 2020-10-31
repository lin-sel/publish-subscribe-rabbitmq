package log

import (
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

// NewLogger Create New Logger Instance
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	// logger.SetLevel(logrus.ErrorLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		DisableColors:   false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			d := formatPackageAndFunctionName(f.Function)
			return "", colorSet(color.FgLightYellow).Sprintf("[%s:%d]",
				formatFilePath(f.File), f.Line) + colorSet(color.FgLightMagenta).Sprintf("[%s]",
				d[0]) + colorSet(color.Cyan).Sprintf("[%s]", d[1])
		},
	})
	return logger
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func formatPackageAndFunctionName(input string) []string {
	var output []string
	if len(input) == 0 {
		return output
	}
	arr := strings.Split(input, ".")
	packagename := strings.Split(arr[1], "/")
	output = append(output, packagename[len(packagename)-1], arr[len(arr)-1])
	return output
}

func colorSet(c ...color.Color) color.Style {
	return color.New(c...)
}
