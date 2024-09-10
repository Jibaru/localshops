package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jibaru/localshops/internal/products/application"
)

type CreateHandler struct {
	serv application.CreateServ
}

func NewCreateHandler(serv application.CreateServ) *CreateHandler {
	return &CreateHandler{
		serv: serv,
	}
}

func (h *CreateHandler) Handle(ctx *gin.Context) {
	var req application.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.serv.Execute(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
