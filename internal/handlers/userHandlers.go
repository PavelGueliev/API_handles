package handlers

import (
	"1/internal/userService"
	"1/internal/web/users"
	"golang.org/x/net/context"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := handler.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, us := range allUsers {
		user := users.User{
			Email:    &us.Email,
			Id:       &us.ID,
			Password: &us.Password,
		}
		response = append(response, user)
	}
	return response, nil
}
func (handler *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}
	createdUser, err := handler.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, err
}

func (handler *UserHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	id := uint(request.Id)
	err := handler.Service.DeleteUser(id)
	if err != nil {
		return users.DeleteUsers404Response{}, err
	}

	return users.DeleteUsers204Response{}, err
}

func (handler *UserHandler) PatchUsers(_ context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	id := uint(request.Id)
	userRequest := request.Body
	userToUpdate := userService.User{
		Password: *userRequest.Password,
		Email:    *userRequest.Email,
	}
	updatedUser, err := handler.Service.UpdateUser(id, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsers200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, err

}
