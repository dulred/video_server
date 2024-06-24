package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

// NewMiddleWareHandler 返回一个实现了 http.Handler 接口的指针
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &middleWareHandler{r: r}
	return m
}

// 实现 ServeHTTP 方法
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 这里可以添加中间件逻辑，比如检查 session
	validateUserSession(r)
	// 调用 httprouter 的 ServeHTTP 方法
	m.r.ServeHTTP(w, r)
}

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandles()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
