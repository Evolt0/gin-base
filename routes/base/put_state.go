package base

import (
	"net/http"
	"reflect"

	"github.com/Parker-Yang/gin-base/global"
	"github.com/Parker-Yang/gin-base/models"
	"github.com/Parker-Yang/gin-base/modules/base"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Put(ctx *gin.Context) {
	data := &models.PutState{}
	err := ctx.BindJSON(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	logrus.Debugln(reflect.TypeOf(data))
	controller := base.NewController(global.Conf.Client)
	state, err := controller.PutState(*data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, state)
}
