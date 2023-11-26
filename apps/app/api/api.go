package main

import (
	"flag"
	"fmt"

	"github.com/POABOB/go-discord/apps/app/api/internal/config"
	"github.com/POABOB/go-discord/apps/app/api/internal/handler"
	"github.com/POABOB/go-discord/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting servers at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
