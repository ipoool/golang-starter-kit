package commands

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/ipoool/golang-starter-kit/drivers"
	"github.com/ipoool/golang-starter-kit/helpers"
	"github.com/ipoool/golang-starter-kit/routes"
	"github.com/rs/cors"
	config "github.com/spf13/viper"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var info = `
HTTP serve running on :
- Config : %s
- Port   : %s
`

func setupConfiguration(path string) error {
	if len(path) == 0 {
		config.AddConfigPath("./")
		if err := config.ReadInConfig(); err != nil {
			return fmt.Errorf("Failed to read the config file: %s", err)
		}
		return nil
	}

	ext := filepath.Ext(path)
	ext = strings.TrimPrefix(ext, ".")
	config.SetConfigType(ext)
	file, err := os.Open(helpers.AbsolutePath(path))
	if err != nil {
		return err
	}
	defer file.Close()
	if err := config.ReadConfig(file); err != nil {
		return fmt.Errorf("Failed to read the config file: %s", err)
	}
	return nil
}

// ServeHTTP - Serve the action HTTP
func ServeHTTP(c *cli.Context) error {
	// setup configuration file
	if err := setupConfiguration(c.String("config")); err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf(info, c.String("config"), config.GetString("app.port")))
	// initialize routers
	var route = routes.Route{
		// setup sql connection
		DB: &drivers.SQL{
			Host:     config.GetString("db.host"),
			Port:     config.GetInt("db.port"),
			Username: config.GetString("db.username"),
			Password: config.GetString("db.password"),
			Charset:  config.GetString("db.charset"),
		},
		// setup redis connection
		Redis: &drivers.Redis{
			Host:     config.GetString("redis.host"),
			Port:     config.GetInt("redis.port"),
			Password: config.GetString("redis.password"),
			DB:       config.GetInt("redis.db"),
		},
	}

	route.Init()
	routers := route.Handle()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(routers)

	// Setup cors
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{config.GetString("app.cors")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            config.GetBool("app.debug"),
	}).Handler(n)

	// Serve http
	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.GetString("app.port")),
		Handler:      cors,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		fmt.Println("Shutting down the server...")
		if err := serve.Shutdown(context.Background()); err != nil {
			fmt.Println("Could not shutdown: ", err.Error())
		}
	}()

	if err := serve.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
