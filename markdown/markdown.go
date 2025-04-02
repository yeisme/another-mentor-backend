package main

import (
	"flag"

	"markdown/internal/config"
	"markdown/internal/data"
	"markdown/internal/handler"
	"markdown/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/markdown.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 加载日志配置
	logx.MustSetup(c.Log)
	logx.Info("configuration loaded successfully")

	// 初始化与存储相关的服务
	data.InitDataService(c)
	
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logx.Info("starting server at ", c.Host, ":", c.Port)
	server.Start()
}
