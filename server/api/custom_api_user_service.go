package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

// UserApiService is a service that implements the logic for the UserApiServicer
// This service should implement the business logic for every endpoint for the UserApi API.
// Include any external packages or services that will be required by this service.
type UserApiService struct {
}

// NewUserApiService creates a default api service
func NewUserApiService() openapi.UserApiServicer {
	return &UserApiService{}
}

type userRepository struct {
	recordList []openapi.GetUserOutput
}

func (ur userRepository) findOne(key int64) (openapi.GetUserOutput, bool) {

	for _, value := range ur.recordList {
		log.Println(value)
		if value.Id == key {
			return value, true
		}
	}
	return openapi.GetUserOutput{}, false
}

var ur = userRepository{
	recordList: []openapi.GetUserOutput{
		{
			Id:           1001,
			EmailAddress: "tanaka@example.co.jp",
			LastName:     "tanaka",
			FirstName:    "taro",
			Birthday:     "2021/01/01",
			Address:      "aaa",
		},
		{
			Id:           1002,
			EmailAddress: "yamada@example.co.jp",
			LastName:     "yamada",
			FirstName:    "hanako",
			Birthday:     "2021/01/01",
			Address:      "bbb",
		},
	},
}

// CreateUser -
func (s *UserApiService) CreateUser(ctx context.Context, createUserInput openapi.CreateUserInput) (openapi.ImplResponse, error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CreateUserOutput{}) or use other options such as http.Ok ...
	//return Response(200, CreateUserOutput{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("CreateUser method not implemented")

}

// GetUser -
func (s *UserApiService) GetUser(ctx context.Context, id int64) (openapi.ImplResponse, error) {
	record, isSuccess := ur.findOne(id)
	log.Println(record)
	if !isSuccess {
		return openapi.Response(http.StatusNotFound, nil), nil
	}
	return openapi.Response(http.StatusOK, record), nil
}

// GetUserList -
func (s *UserApiService) GetUserList(ctx context.Context) (openapi.ImplResponse, error) {
	// TODO - update GetUserList with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, []map[string]interface{}{}), nil

	// return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetUserList method not implemented")
	return openapi.Response(http.StatusOK, ur.recordList), nil
}
