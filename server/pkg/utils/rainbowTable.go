package utils

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"tools/global"
)

// BuildRainbowTable 构建彩虹表
func BuildRainbowTable() error {
	start := time.Now()
	// 读取明文和对应的 MD5 密文文件
	curDir, _ := os.Getwd()
	md5Path := filepath.Join(curDir, global.Config.RainbowRainbow.Path)

	// 打开明文文件
	file, err := os.Open(md5Path)
	if err != nil {
		global.Log.Error(fmt.Sprintf("Error opening file: %s", err.Error()))
		return err
	}

	defer file.Close()

	// 创建彩虹表
	rainbowTable := make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lastColonIndex := strings.LastIndex(line, ":")
		if lastColonIndex == -1 {
			return err
		}

		p1 := line[:lastColonIndex]
		c1 := line[lastColonIndex+1:]

		rainbowTable[c1] = p1
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	end := time.Now()
	elapsed := end.Sub(start) // 计算用时多久

	global.Log.Info("构建彩虹表用时" + elapsed.String() + "秒")
	global.RainBowTable = &rainbowTable

	return nil
}

func WatchTextFile() error {
	curDir, _ := os.Getwd()
	filePath := filepath.Join(curDir, global.Config.RainbowRainbow.Path)
	// 创建文件监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("创建文件监视器失败：%w", err)
	}
	defer watcher.Close()

	// 添加要监视的文件
	err = watcher.Add(filePath)
	if err != nil {
		return fmt.Errorf("添加文件到监视器失败：%w", err)
	}

	// 启动协程处理文件事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("文件 %s 发生更改\n", event.Name)
					err := BuildRainbowTable()
					if err != nil {
						panic(err.Error())
						return
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("文件监视错误：%v\n", err)
			}
		}
	}()

	// 等待信号
	done := make(chan bool)
	<-done

	return nil
}

// LookupRainbowTable 在彩虹表中查找对应的明文
func LookupRainbowTable(md5 string, rainbowTable map[string]string) string {
	// 创建互斥锁
	var mutex sync.Mutex

	// 并发安全地读取map
	mutex.Lock()
	plaintext, found := rainbowTable[md5]
	mutex.Unlock()

	if found {
		return plaintext
	}
	return ""
}
