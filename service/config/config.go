package config

import (
	"github.com/micro/cli/v2"
	proto "github.com/micro/go-micro/v2/config/source/service/proto"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	"github.com/micro/micro/v2/service"
	"github.com/micro/micro/v2/service/config/handler"
)

const (
	name = "go.micro.config"
)

var (
	// Flags specific to the config service
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "watch_topic",
			EnvVars: []string{"MICRO_CONFIG_WATCH_TOPIC"},
			Usage:   "watch the change event.",
		},
	}
)

// Run micro config
func Run(c *cli.Context) error {
	if len(c.String("watch_topic")) > 0 {
		handler.WatchTopic = c.String("watch_topic")
	}

	srv := service.New(service.Name(name))

	srv.Options().Store.Init(store.Table("config"))
	h := &handler.Config{Store: srv.Options().Store}

	proto.RegisterConfigHandler(srv.Server(), h)
	service.RegisterSubscriber(handler.WatchTopic, srv.Server(), handler.Watcher)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	return nil
}
