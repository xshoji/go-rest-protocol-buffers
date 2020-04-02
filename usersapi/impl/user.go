package proto_impl

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	usersapi "github.com/xshoji/go-rest-protocol-buffers/usersapi/proto"
	"sync"
)

type UserService struct{}

var UserRepositoryMap = make(map[int32]*usersapi.User)
var idSequence int32 = 1
var mu sync.Mutex

func createResponse(status int64, message string) *usersapi.UserServiceResponse {
	return &usersapi.UserServiceResponse{
		Status:  status,
		Message: message,
	}
}

func createOkResponse() *usersapi.UserServiceResponse {
	return createResponse(200, "OK")
}

func createNotFoundResponse() *usersapi.UserServiceResponse {
	return createResponse(500, "Not found")
}

func (*UserService) Create(context context.Context, user *usersapi.User) (*usersapi.UserServiceResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	user.Id = idSequence
	UserRepositoryMap[idSequence] = user
	idSequence++
	response := createOkResponse()
	response.Users = []*usersapi.User{user}
	return response, nil
}

func (*UserService) Read(ctx context.Context, in *usersapi.UserServiceSelector) (*usersapi.UserServiceResponse, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok != true {
		return createNotFoundResponse(), nil
	}
	response := createOkResponse()
	response.Users = []*usersapi.User{foundUser}
	return response, nil
}

func (*UserService) ReadAll(ctx context.Context, in *empty.Empty) (*usersapi.UserServiceResponse, error) {
	users := make([]*usersapi.User, 0, len(UserRepositoryMap))
	for _, v := range UserRepositoryMap {
		users = append(users, v)
	}
	response := createOkResponse()
	response.Users = users

	return response, nil
}

func (*UserService) Update(context context.Context, user *usersapi.User) (*usersapi.UserServiceResponse, error) {
	foundUser, ok := UserRepositoryMap[user.Id]
	if ok != true {
		return createNotFoundResponse(), nil
	}

	if user.Name != "" {
		foundUser.Name = user.Name
	}
	if user.GetNicknameOptional() != nil {
		foundUser.NicknameOptional = user.GetNicknameOptional()
	}
	if user.GetAgeOptional() != nil {
		foundUser.AgeOptional = user.GetAgeOptional()
	}
	if user.GetBirth() != nil {
		foundUser.Birth = user.GetBirth()
	}
	if user.GetAddresss() != nil {
		foundUser.Addresss = user.GetAddresss()
	}
	fmt.Println(foundUser)
	fmt.Println(user)
	response := createOkResponse()
	response.Users = []*usersapi.User{foundUser}
	return response, nil
}

func (*UserService) Delete(context context.Context, in *usersapi.UserServiceSelector) (*usersapi.UserServiceResponse, error) {
	foundUser, ok := UserRepositoryMap[in.Id]
	if ok {
		delete(UserRepositoryMap, in.Id)
	} else {
		return createNotFoundResponse(), nil
	}

	response := createOkResponse()
	response.Users = []*usersapi.User{foundUser}
	return response, nil
}
