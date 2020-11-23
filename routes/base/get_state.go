package base

import (
	"net/http"

	"github.com/Parker-Yang/gin-base/global"
	"github.com/Parker-Yang/gin-base/models"
	"github.com/Parker-Yang/gin-base/modules/base"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	key := ctx.Param("key")
	controller := base.NewController(global.Conf.Client)
	State := &models.GetState{
		Key: key,
	}
	state, err := controller.GetState(State)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, state)
}
