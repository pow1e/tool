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

	req, err := http.NewRequest("GET", "http://8.134.126.146:9000", nil)
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
	// map[Accept-Ranges:[bytes] Cache-Control:[no-cache, no-store, must-revalidate] Content-Type:[text/html; charset=utf-8] Date:[Tue, 08 Aug 2023 08:11:47 GMT] Last-Modified:[Mon, 22 May 2023 08:32:44 GMT] Vary:[Accept-Encoding] X-Content-Type-Options:[nosniff] X-Xss-Protection:[1; mode=block]]
	t.Log(resp.Header.Get("Content-Type"))
	t.Log(resp.Header.Get("Last-Modified"))
	t.Log(resp.Header)
	t.Log(resp.Status)
	t.Log(resp.Request)
	t.Log(resp.ContentLength)
	t.Log(resp.TLS)
	t.Log(resp.TransferEncoding)
	t.Log(resp.Proto)
	t.Log(resp.StatusCode)
	t.Log(title)
	// &http.Response{Status:"200 OK", StatusCode:200, Proto:"HTTP/1.1", ProtoMajor:1, ProtoMinor:1, Header:http.Header{"Accept-Ranges":[]string{"bytes"},
	// "Cache-Control":[]string{"no-cache, no-store, must-revalidate"},
	// "Content-Type":[]string{"text/html; charset=utf-8"},
	// "Date":[]string{"Tue, 08 Aug 2023 12:07:49 GMT"},
	// "Last-Modified":[]string{"Mon, 22 May 2023 08:32:44 GMT"},
	// "Vary":[]string{"Accept-Encoding"},
	// "X-Content-Type-Options":[]string{"nosniff"},
	// "X-Xss-Protection":[]string{"1; mode=block"}},
	// Body:(*http.gzipReader)(0xc000184040), ContentLength:-1, TransferEncoding:[]string{"chunked"},
	// Close:false, Uncompressed:true,
	// Trailer:http.Header(nil),
	// Request:(*http.Request)(0xc00013c100),
	// TLS:(*tls.ConnectionState)(nil)}
	t.Logf("%#v\n", resp.Header.Get("Date"))
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
