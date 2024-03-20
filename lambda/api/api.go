package api

import (
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error unmarshalling request %v", err),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.Username == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body:       "username and password cannot be empty",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	userExists, err := api.dbStore.DoesUserExists(registerUser.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error checking if user exists %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if userExists {
		return events.APIGatewayProxyResponse{
			Body:       "user already exists",
			StatusCode: http.StatusConflict,
		}, nil
	}

	user, err := types.NewUser(registerUser)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error creating user %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	err = api.dbStore.InsertUser(user)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error inserting user %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "user registered successfully",
		StatusCode: http.StatusCreated,
	}, nil
}

func (api ApiHandler) LoginUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := json.Unmarshal([]byte(request.Body), &loginRequest)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error unmarshalling request %v", err),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	loginUser, err := api.dbStore.GetUser(loginRequest.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error getting user %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if loginUser.Username == "" || loginUser.PasswordHash == "" {
		return events.APIGatewayProxyResponse{
			Body:       "username and password cannot be empty",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	user, err := api.dbStore.GetUser(loginUser.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("error getting user %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if !types.ValidatePassword(user.PasswordHash, loginRequest.Password) {
		return events.APIGatewayProxyResponse{
			Body:       "invalid username or password",
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	accessToken := types.CreateToken(user)
	successMsg := fmt.Sprintf(`{"access_token": "%s"}`, accessToken)

	return events.APIGatewayProxyResponse{
		Body:       successMsg,
		StatusCode: http.StatusOK,
	}, nil

}
