package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// doRequest 发起HTTP请求，并在上下文取消时停止
func doRequest(ctx context.Context, url string, wg *sync.WaitGroup, cancelFunc context.CancelFunc) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		cancelFunc() // 出现创建请求的错误，取消后续所有请求
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		cancelFunc() // 请求失败，取消后续所有请求
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败: 状态码 %d\n", resp.StatusCode)
		cancelFunc() // HTTP状态不是200，取消后续所有请求
		return
	}

	fmt.Println("成功完成请求:", url)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urls := []string{
		"http://example.com/request1",
		"http://example.com/request2",
		"http://example.com/request3",
		"http://example.com/request4",
		"http://example.com/request5",
	}

	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go doRequest(ctx, url, &wg, cancel)

		// 检查第二个请求后是否有取消信号
		if i == 1 {
			time.Sleep(2 * time.Second) // 等待第二个请求的结果
			if ctx.Err() != nil {
				break // 如果有取消信号，则不再发起新的请求
			}
		}
	}

	// 等待所有已经开始的goroutine
	wg.Wait()

	// 检查上下文是否被取消
	if ctx.Err() == context.Canceled {
		fmt.Println("操作被取消")
	} else {
		fmt.Println("所有请求成功完成")
	}
}
