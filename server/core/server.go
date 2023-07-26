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

var prevSrv server // 使用你定义的 server 接口作为类型来保存上一次的服务器实例

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

	// 先关闭上一次的服务器（如果存在）
	if prevSrv != nil {
		fmt.Println("Shutting down the previous server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := prevSrv.Shutdown(ctx); err != nil {
			// 处理错误，可以记录日志等
			fmt.Println("Previous server shutdown error:", err)
		}
		fmt.Println("Previous server shutdown complete.")
	}

	// 创建新的服务器实例
	srv := initServer(address, routers)

	go func() {
		// 启动新的服务器
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Error(err.Error())
		}
	}()
	global.Log.Info("Server run success on ", zap.String("address", address))

	// 将当前服务器实例赋值给 prevSrv，以备下次关闭使用
	prevSrv = srv

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	global.Log.Info("Server exiting ...")
}
