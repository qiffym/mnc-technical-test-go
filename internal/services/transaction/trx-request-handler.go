package transaction

import (
	"mncPaymentAPI/internal/adapter/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RequestHandler struct {
		ctrl ControllerInterface
	}
)

func (rh RequestHandler) SendTrx(ctx *gin.Context) {
	var payload = SendTrx{}
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}
	res, err := rh.ctrl.SendTrx(ctx, payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
	}

	ctx.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetLogTrx(ctx *gin.Context) {
	res, err := rh.ctrl.GetLogTrx(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
	}
	ctx.JSON(http.StatusOK, res)
}
