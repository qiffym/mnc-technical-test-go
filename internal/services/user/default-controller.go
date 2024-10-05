package user

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
		Register(
			ctx context.Context,
			payload RegisterUser,
		) (*dto.Response, error)

		GetAll(
			ctx context.Context,
		) (*dto.Response, error)

		LoginUser(
			ctx context.Context,
			email, password string,
		) (SuccessLoginUser, error)

		Remove(
			ctx context.Context,
			id uint,
		) (SuccessRemoveUser, error)

		UpdateUser(
			ctx context.Context,
			id uint,
			updateData map[string]interface{},
		) (UpdateUserPayload, error)
	}
)

func (ctrl Controller) Register(
	ctx context.Context,
	payload RegisterUser,
) (*dto.Response, error) {
	start := time.Now()
	result, err := ctrl.Uc.RegisterUser(ctx, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		result,
		"Register is success",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) GetAll(
	ctx context.Context,
) (*dto.Response, error) {
	start := time.Now()
	res, err := ctrl.Uc.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dto.NewSuccessResponse(
		res,
		"list of user",
		fmt.Sprint(time.Since(start).Milliseconds(), " ms."),
	), nil
}

func (ctrl Controller) LoginUser(
	ctx context.Context,
	email, password string,
) (SuccessLoginUser, error) {
	user, tokenString, err := ctrl.Uc.LoginUser(ctx, email, password)
	if err != nil {
		return SuccessLoginUser{}, err
	}
	response := SuccessLoginUser{
		Response: dto.ResponseMeta{
			Success: true,
			Message: "Login succeed",
		},
		Email:       user.Email,
		AccessToken: tokenString,
	}

	return response, nil
}

func (ctrl Controller) Remove(
	ctx context.Context,
	id uint,
) (SuccessRemoveUser, error) {
	user, err := ctrl.Uc.DeleteUser(ctx, id)
	if err != nil {
		return SuccessRemoveUser{}, err
	}

	if user == nil {
		return SuccessRemoveUser{}, err
	}

	response := SuccessRemoveUser{
		Response: dto.ResponseMeta{
			Success: true,
			Message: "Removed",
		},
		ID: user.ID,
	}

	return response, nil
}

func (ctrl Controller) UpdateUser(
	ctx context.Context,
	id uint,
	updateData map[string]interface{},
) (UpdateUserPayload, error) {
	user, err := ctrl.Uc.UpdateUser(ctx, id, updateData)
	if err != nil {
		return UpdateUserPayload{}, err
	}

	return UpdateUserPayload{
		ID: user.ID,
	}, nil
}
