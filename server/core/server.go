package core

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"tools/global"
	"tools/router"
)

// 定义sever接口 实现ListenAndServe方法
type server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func RunServer() {
	fmt.Println(`
 _________    ________      ________      ___          
|\___   ___\ |\   __  \    |\   __  \    |\  \         
\|___ \  \_| \ \  \|\  \   \ \  \|\  \   \ \  \        
     \ \  \   \ \  \\\  \   \ \  \\\  \   \ \  \       
      \ \  \   \ \  \\\  \   \ \  \\\  \   \ \  \____  
       \ \__\   \ \_______\   \ \_______\   \ \_______\
        \|__|    \|_______|    \|_______|    \|_______|
                                                       
                                                       
                                                       `)
	routers := router.Routers() // 初始化路由
	address := fmt.Sprintf(":%s", global.Config.System.Addr)
	srv := initServer(address, routers)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(err.Error())
		}
	}()
	global.Log.Info("server run success on ", zap.String("address", address))

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	global.Log.Info("server exiting ...")
}
