package transaction

import (
	"context"
	"fmt"
	"log"
	"mncPaymentAPI/internal/adapter/dto"
	"time"
)

type (
	Controller struct {
		Uc UsecaseInterface
	}

	ControllerInterface interface {
		SendTrx(
			ctx context.Context,
			payload SendTrx,
		) (*dto.Response, error)

		GetLogTrx(
			ctx context.Context,
		) (*dto.Response, error)
	}
)

func (ctrl Controller) SendTrx(
	ctx context.Context,
	payload SendTrx,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.Uc.SendTrx(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Transaction is Success",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) GetLogTrx(
	ctx context.Context,
) (*dto.Response, error) {
	start := time.Now()
	res, err := ctrl.Uc.GetLogTrx(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		res,
		"list of log transaction",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}
