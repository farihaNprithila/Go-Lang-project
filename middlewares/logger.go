package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func SetupLogOutput() {
	absolutePath, _ := os.Getwd()
	path := filepath.Join(absolutePath, "log", "golang.log")

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("File not found")
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
