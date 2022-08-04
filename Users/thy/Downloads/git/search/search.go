package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/config"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/middleware"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	// 全局中间件
	//server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		logx.Info("global middleware")
	//		next(w, r)
	//	}
	//})
	//handler.RegisterHandlers(server, ctx)
	server.Use(middleware.Middleware)
	server.Use(middleware.MiddlewareWithAnotherService(new(middleware.AnotherService)))
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
