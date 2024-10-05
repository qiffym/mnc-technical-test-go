package transaction

import (
	"mncPaymentAPI/internal/adapter/Repository"
	"mncPaymentAPI/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Router struct {
		rq *RequestHandler
	}
)

func TrxRoute(
	db *gorm.DB,
) *Router {
	return &Router{rq: &RequestHandler{
		ctrl: &Controller{
			Uc: UseCase{
				TrxRepo: Repository.NewTrxRepo(db),
			},
		},
	},
	}
}

func (r Router) Route(router *gin.RouterGroup) {
	trx := router.Group("/transaction")

	trx.POST(
		"/send",
		middleware.Authentication(),
		r.rq.SendTrx,
	)
	trx.GET(
		"/log",
		r.rq.GetLogTrx,
	)
}
