package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jibaru/localshops/internal/shops/application"
)

type GetAllHandler struct {
	serv application.GetAllServ
}

func NewGetAllHandler(serv application.GetAllServ) *GetAllHandler {
	return &GetAllHandler{
		serv: serv,
	}
}

func (h *GetAllHandler) Handle(ctx *gin.Context) {
	resp, err := h.serv.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
