package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Use Postgres Driver for DB connection
)

// App ...
type App struct {
	Config *Config
	DB     *sql.DB
}

// NewApp ...
func NewApp(cfgFile string) *App {
	config, err := NewConfig(cfgFile)
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
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		log.Fatalf("Cannot connect to Postgres %v", err)
		return nil
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
