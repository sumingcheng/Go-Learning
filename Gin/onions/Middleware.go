package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(
	f http.HandlerFunc,
	middlewares ...Middleware,
) http.HandlerFunc {
	if len(middlewares) == 0 {
		return f
	}

	// 将中间件逆序堆叠，形成洋葱模型
	wrapped := f
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}
	return wrapped
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		// 在请求处理之前执行
		log.Printf("Request URI: %s", r.RequestURI)
		log.Printf("Request Method: %s", r.Method)
		// 调用下一个中间件或最终的handler
		next(w, r)
		// 在请求处理之后执行
		log.Println("Request finished")
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		// 认证逻辑
		authHeader := r.Header.Get("Authorization")
		if authHeader != "valid-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// 认证通过后执行下一个中间件或handler
		next(w, r)
	}
}
