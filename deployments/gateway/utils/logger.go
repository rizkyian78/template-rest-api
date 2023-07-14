package utils

import (
	"fmt"
	"os"
	"time"

	joonix "github.com/joonix/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerRequest(param gin.LogFormatterParams) string {

	// your custom format
	return fmt.Sprintf("[%s] \n - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func LoggerProcess() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(joonix.NewFormatter())

	logger.SetOutput(os.Stdout)

	return logger
}
