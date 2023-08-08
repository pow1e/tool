package tests

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestBody(t *testing.T) {

	req, err := http.NewRequest("GET", "http://8.134.126.146:6379", nil)
	if err != nil {
		t.Fatal(err)
	}

	var (
		UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
		Accept    = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	)
	req.Header.Set("User-agent", UserAgent)
	req.Header.Set("Accept", Accept)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	client := http.DefaultClient
	resp, err := client.Do(req)

	if errors.Is(err, io.EOF) {
		t.Fatal(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if errors.Is(err, io.EOF) {
		t.Log(err)
	}
	title := gettitle(body)

	t.Log(string(body))
	t.Log(resp.Status)
	t.Log(resp.StatusCode)
	t.Log(title)
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
