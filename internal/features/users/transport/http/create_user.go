package users_transport_http

import (
	"net/http"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
)

type CreateUserResponse UserDTO

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100" example:"Ivan Ivanov"`
	Password    string  `json:"password" validate:"required,min=6,max=100" example:"qwerty123"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164" example:"+79998887766"`
	Email       string  `json:"email" validate:"required,email" example:"somemail@mail.ru"`
}

// CreateUser godoc
// @Summary creates user
// Description creates user in db
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "CreateUser request body"
// @Success 201 {object} CreateUserResponse "Successfully created user"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/ [post]
func (t *UserTransport) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	ctx := r.Context()

	logger := core_logger.FromContext(ctx)
	logger.Debug("started creating user")

	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	if err := core_http_request.DecodeAndValidate(r, &createUserRequest); err != nil {
		responseHandler.RespondError("validate request", err)
		return
	}

	userDomain := domainFromDTO(createUserRequest)
	userDomain, err := t.UserService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.RespondError("failed to create user", err)
		return
	}

	response := CreateUserResponse(dtoFromDomain(userDomain))
	responseHandler.RespondJSON(http.StatusCreated, response)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(
		dto.FullName,
		dto.Password,
		dto.PhoneNumber,
		dto.Email,
	)
}
