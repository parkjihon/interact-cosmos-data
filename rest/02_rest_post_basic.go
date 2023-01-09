package rest

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// http://golang.site/go/article/103-HTTP-POST-%ED%98%B8%EC%B6%9C
func QueryRest02() error {
	// 간단한 http.PostForm 예제
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	return nil
}
