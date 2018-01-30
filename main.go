package main

import (
	"github.com/jrgirvan/devx/handlers"

	"flag"
	"github.com/jrgirvan/devx/config"
	"github.com/jrgirvan/devx/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	appVersion  string
	showVersion bool
)

func main() {

	configure()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbHandle := db.GetDB()

	e.File("/", "public/index.html")
	e.Static("/static", "public")
	e.GET("/talks", handlers.GetTalks(dbHandle))
	e.PUT("/talks/:id", handlers.PutTalk(dbHandle))
	// e.PUT("/talks", handlers.CreateTalk(dbHandle))
	e.POST("/talks", handlers.CreateTalk(dbHandle))
	e.DELETE("/talks/:id", handlers.DeleteTalk(dbHandle))

	e.Run(standard.New(":80"))

	defer db.CloseDB()
}

func configure() {
	configFile := flag.String("config-file", "./default.config", "Path to the configuration file")
	flag.BoolVar(&showVersion, "version", false, "Shows the version")
	flag.Parse()

	if showVersion {
		log.Debug("Version: %s\n", appVersion)
		return
	}

	config, err := config.ReadConfig(*configFile)
	if err != nil {
		log.Debug(err)
	}

	db.Init(config.DBDriver, config.DBConfig)

}
