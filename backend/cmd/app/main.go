package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	product_routes "github.com/jibaru/localshops/internal/products/infrastructure/routes"
	"github.com/jibaru/localshops/internal/shared/infrastructure/logger"
	shop_routes "github.com/jibaru/localshops/internal/shops/infrastructure/routes"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func main() {
	// Log configuration
	logHandler := logger.NewLogHandler(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slog.New(logHandler))

	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env", "error", err)
		return
	}

	c := DatabaseConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})
	if err != nil {
		slog.Error("can not init db", "error", err)
		return
	}

	// Gin configuration
	gin.SetMode(gin.ReleaseMode)

	// Routes configuration
	r := gin.New()
	r.Use(logger.LogMiddleware)
	r.Use(gin.Recovery())

	shop_routes.Set(r, db)
	product_routes.Set(r, db)

	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	slog.Info("running app", "address", address)

	if err := r.Run(address); err != nil {
		slog.Error("error running app", "error", err)
	}
}
