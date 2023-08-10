package system

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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
	ScanUrl                string   `json:"Url"`    // 扫描url
	Status                 string   `json:"Status"` // http协议+状态码
	Title                  string   `json:"Title"`  // 标题
	Data                   string   `json:"Data"`
	ContentType            string   `json:"Content-Type"`  // 内容格式
	LastModified           string   `json:"Last-Modified"` // 上次修改时间
	X_Content_Type_Options string   `json:"X-Content-Type-Options"`
	X_Xss_Protection       string   `json:"X-Xss-Protection"`
	Server                 string   `json:"Server"`
	TLS                    string   `json:"TLS"`
	TransferEncoding       []string `json:"Transfer-Encoding"`
	X_Powered_By           string   `json:"X-Powered-By"`
}

func BasePortScanner(open bool, ip string, port int) *PortScanResp {
	return &PortScanResp{
		IsOpen: open,
		Port:   port,
		PortInfo: &PortInfo{
			ScanUrl: fmt.Sprintf("http://%s:%d", ip, port),
		},
	}
}

// BuildInfoScan 收集信息
func (p *PortScanResp) BuildInfoScan() *PortScanResp {
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
		p.PortInfo.Status = resp.Proto + " " + resp.Status
		p.PortInfo.Data = resp.Header.Get("Date")
		p.PortInfo.ContentType = resp.Header.Get("Content-Type")
		p.PortInfo.LastModified = resp.Header.Get("Last-Modified")
		p.PortInfo.X_Content_Type_Options = resp.Header.Get("X-Content-Type-Options")
		p.PortInfo.X_Xss_Protection = resp.Header.Get("X-Xss-Protection")
		p.PortInfo.Server = resp.Header.Get("Server")
		p.PortInfo.TLS = resp.Header.Get("TLS")
		p.PortInfo.TransferEncoding = resp.TransferEncoding
		p.PortInfo.X_Powered_By = resp.Header.Get("X-Powered-By")
	} else {
		lock.Lock()
		p.PortInfo.Status = strconv.Itoa(http.StatusBadGateway)
		lock.Unlock()
	}
	return p
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
