package main

//import (
//	"flag"
//	"fmt"
//	"goZeroDemo/Users/thy/Downloads/git/gozerodemo"
//
//	"goZeroDemo/Users/thy/Downloads/git/internal/config"
//	"goZeroDemo/Users/thy/Downloads/git/internal/server"
//	"goZeroDemo/Users/thy/Downloads/git/internal/svc"
//
//	"github.com/zeromicro/go-zero/core/conf"
//	"github.com/zeromicro/go-zero/core/service"
//	"github.com/zeromicro/go-zero/zrpc"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//)
//
//var configFile = flag.String("f", "etc/gozerodemo.yaml", "the config file")
//
//func main() {
//	flag.Parse()
//
//	var c config.Config
//	conf.MustLoad(*configFile, &c)
//	ctx := svc.NewServiceContext(c)
//	svr := server.NewGoZeroDemoServer(ctx)
//
//	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
//		goZeroDemo.RegisterGoZeroDemoServer(grpcServer, svr)
//
//		if c.Mode == service.DevMode || c.Mode == service.TestMode {
//			reflection.Register(grpcServer)
//		}
//	})
//	defer s.Stop()
//
//	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
//	s.Start()
//}
