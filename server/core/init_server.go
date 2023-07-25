package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(addr string, router *gin.Engine) server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    1000 * time.Second, // 读超时
		WriteTimeout:   1000 * time.Second, // 写超时
		MaxHeaderBytes: 1 << 20,
	}
}
