package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// doRequest 发起一个HTTP请求，并在上下文取消时停止
func doRequest(ctx context.Context, url string) error {
	// http.NewRequestWithContext 创建一个新的请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	// 使用http.DefaultClient发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码是否为200
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败: 状态码 %d", resp.StatusCode)
	}
	return nil
}

func main() {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // 确保所有退出路径都取消上下文

	// 请求的URL列表
	urls := []string{
		"http://example.com/request1",
		"http://example.com/request2",
		"http://example.com/request3",
	}

	// 顺序执行每个请求
	for _, url := range urls {
		err := doRequest(ctx, url)
		if err != nil {
			fmt.Println("在执行请求时发生错误:", err)
			break
		}
		fmt.Println("成功完成请求:", url)
	}

	// 检查上下文是否被取消
	if ctx.Err() == context.Canceled {
		fmt.Println("操作被取消")
	} else if ctx.Err() == context.DeadlineExceeded {

		fmt.Println("操作超时")
	}

	// 如果没有错误，说明所有请求都成功执行
	if ctx.Err() == nil {
		fmt.Println("所有请求成功完成")
	}
}
