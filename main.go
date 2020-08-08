package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// App ...
type App struct {
	Config *Config
	DB     *sql.DB
}

// NewApp ...
func NewApp() *App {
	config, err := NewConfig("./config.yml")
	if err != nil {
		log.Fatal("Cannot load config file")
		return nil
	}

	pgInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBName,
		config.Postgres.SSLMode,
	)
	fmt.Printf("--- connecting Postgres %s\n", pgInfo)
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}

	return &App{
		Config: config,
		DB:     db,
	}
}

// NewServer ...
func (app *App) NewServer() *gin.Engine {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*")
	e.Static("/html", "./public/tmp")
	e.Static("/js", "./public/js")

	e.GET("/addItem", RenderAddItemHandler)
	api := e.Group("api")
	{
		api.POST("/urls", app.CreateURLHandler)
	}

	return e
}

func main() {
	app := NewApp()
	server := app.NewServer()

	server.Run(fmt.Sprintf(":%v", app.Config.Server.Port))
}
