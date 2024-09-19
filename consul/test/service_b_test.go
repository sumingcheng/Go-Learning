package test

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestService_b(t *testing.T) {
	// 直接发起对服务 A 的 GET 请求
	resp, err := http.Get("http://172.16.50.251:8081/call-a")
	if err != nil {
		t.Fatal("创建请求失败:", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		t.Errorf("期望的状态码 200，但是获取的是 %d", resp.StatusCode)
	}

	// 读取响应体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("读取响应体失败:", err)
	}
	body := string(bodyBytes)

	// 打印响应体内容
	t.Logf("响应体内容: %s", body)
}
