package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jibaru/localshops/internal/products/application"
	"github.com/jibaru/localshops/internal/products/infrastructure/handlers"
	"github.com/jibaru/localshops/internal/products/infrastructure/repositories/orm"
)

func Set(r *gin.Engine, db *gorm.DB) {
	productRepo := orm.NewProductRepo(db)
	createProductServ := application.NewCreateServ(productRepo)
	createProductHandler := handlers.NewCreateHandler(createProductServ)

	r.POST("/products", createProductHandler.Handle)
}
