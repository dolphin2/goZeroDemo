package middleware

import (
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type (
	AnotherService struct{}

	Request struct {
		User string `form:"user"`
	}
)

func (s *AnotherService) GetToken() string {
	return stringx.Rand()
}

//
//func NewExampleMiddleware() *ExampleMiddleware {
//	return &ExampleMiddleware{}
//}
//
//func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
//	logx.Info("example middle")
//	return func(w http.ResponseWriter, r *http.Request) {
//		// TODO generate middleware implement function, delete after code implementation
//
//		// Passthrough to next handler if need
//		next(w, r)
//	}
//}

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "static-middleware")
		next(w, r)
	}
}

func MiddlewareWithAnotherService(s *AnotherService) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Middleware", s.GetToken())
			next(w, r)
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "hello, "+req.User)
}
