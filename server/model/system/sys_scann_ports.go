package system

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	Accept    = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
)

type PortScanResp struct {
	Port     int       `json:"port"`
	IsOpen   bool      `json:"isOpen"`
	PortInfo *PortInfo `json:"Info"`
}

type PortInfo struct {
	ScanUrl string `json:"url"`
	Title   string `json:"title"`
	Code    int    `json:"code"`
}

func BasePortScanner(ip string, port int) *PortScanResp {
	return &PortScanResp{
		IsOpen: true,
		Port:   port,
		PortInfo: &PortInfo{
			ScanUrl: fmt.Sprintf("http://%s:%d", ip, port),
		},
	}
}

// BuildInfoScan 收集信息
func (p *PortScanResp) BuildInfoScan() {
	var lock sync.Mutex

	req, _ := http.NewRequest("GET", p.PortInfo.ScanUrl, nil)

	req.Header.Set("User-agent", UserAgent)
	req.Header.Set("Accept", Accept)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	client := http.DefaultClient
	resp, _ := client.Do(req)

	if resp != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		lock.Lock()
		p.PortInfo.Title = gettitle(data)
		p.PortInfo.Code = resp.StatusCode
		lock.Unlock()
	} else {
		lock.Lock()
		p.PortInfo.Code = http.StatusBadGateway
		lock.Unlock()
	}
}

func gettitle(body []byte) (title string) {
	re := regexp.MustCompile("(?ims)<title>(.*?)</title>")
	find := re.FindSubmatch(body)
	if len(find) > 1 {
		title = string(find[1])
		title = strings.TrimSpace(title)
		title = strings.Replace(title, "\n", "", -1)
		title = strings.Replace(title, "\r", "", -1)
		title = strings.Replace(title, "&nbsp;", " ", -1)
		if len(title) > 100 {
			title = title[:100]
		}
	}
	if title == "" {
		title = "None"
	}
	return
}
