package main

import (
	"context"
	"flag"
	"github.com/josearomeroj/xm-exercise/pkg/logging"
	"github.com/josearomeroj/xm-exercise/pkg/service"
	"time"
)

func init() {
	loc, _ := time.LoadLocation("Africa/Cairo")
	time.Local = loc
}

var (
	configFile = flag.String("config", "example_config.yaml", "configuration file")
)

var (
	mainLogger    = logging.NewLogger("main")
	serviceLogger = logging.NewLogger("service")
	dbLogger      = logging.NewLogger("db")
)

func main() {
	ctx := context.Background()

	flag.Parse()

	mainLogger.Infof("Loading config from %s file.", *configFile)
	conf, err := service.LoadConfig(*configFile)
	if err != nil {
		mainLogger.Fatalf("Could not load config from %s, error: %s", *configFile, err)
	}

	service.StartService(ctx, conf, serviceLogger, dbLogger, mainLogger, false)
}
