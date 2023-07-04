package main

import (
	"context"

	"github.com/gogf/gf/contrib/registry/polaris/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
)

func main() {
	conf := config.NewDefaultConfiguration([]string{"192.168.110.30:8091"})
	conf.Consumer.LocalCache.SetPersistDir("./manifest/logs/polaris/backup")
	if err := api.SetLoggersDir("./manifest/logs/polaris/log"); err != nil {
		g.Log().Fatal(context.Background(), err)
	}

	// TTL egt 2*time.Second
	gsvc.SetRegistry(polaris.NewWithConfig(conf, polaris.WithTTL(10)))

	s := g.Server(`hello-world.svc`)
	s.BindHandler("/", func(r *ghttp.Request) {
		g.Log().Info(r.Context(), `request received three`)
		r.Response.Write(`Hello world three`)
	})
	s.Run()
}
