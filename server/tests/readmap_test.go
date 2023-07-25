package tests

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
)

const (
	blockSize = 4096 // 分块大小，可根据实际情况调整
)

func TestMd5(t *testing.T) {
	mergedFilename := "D:\\Language\\Go\\Go_code\\src\\tools\\static\\dict\\md5.txt"

	rainbowTable, err := buildRainbowTable(mergedFilename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rainbow table built successfully!")
	// 可以在这里使用生成的彩虹表进行密码破解等操作
	table := LookupRainbowTable("c21cb1f8d8ce7b5a586bd338e9e7c1a5", rainbowTable)
	if table == "" {
		println("失败")
		return
	}
	println("成功", table)

}
func buildRainbowTable(mergedFilename string) (map[string]string, error) {
	mergedFile, err := os.Open(mergedFilename)
	if err != nil {
		return nil, err
	}
	defer mergedFile.Close()

	rainbowTable := make(map[string]string)

	scanner := bufio.NewScanner(mergedFile)

	for scanner.Scan() {
		line := scanner.Text()
		lastColonIndex := strings.LastIndex(line, ":")
		if lastColonIndex == -1 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		p1 := line[:lastColonIndex]
		c1 := line[lastColonIndex+1:]

		rainbowTable[c1] = p1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rainbowTable, nil
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
