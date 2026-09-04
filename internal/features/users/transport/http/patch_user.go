package users_transport_http

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/qury53/todo-app2/internal/core/domain"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
	core_logger "github.com/qury53/todo-app2/internal/core/logger"
	core_http_request "github.com/qury53/todo-app2/internal/core/transport/http/request"
	transport_http_response "github.com/qury53/todo-app2/internal/core/transport/http/response"
	core_http_types "github.com/qury53/todo-app2/internal/core/transport/http/types"
	core_http_utils "github.com/qury53/todo-app2/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" validate:"required,min=3,max=100" example:"Ivanov Ivan Ivanovich"`
	Password    core_http_types.Nullable[string] `json:"password" validate:"required,min=6,max=100" example:"124ojolsdjgOejoi"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" validate:"omitempty,e164" example:"+79998887766"`
	Email       core_http_types.Nullable[string] `json:"email" validate:"required,email" example:"some_mail@mail.ru"`
}

type PatchUserResponse UserDTO

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("fullname cannot be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value)) // гпт рекомендует использовать utf8.RuneCountInString() тк он быстрее
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("fullname length must be between 3 and 100: %w", core_errors.ErrInvalidArgument)
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			if !core_http_request.IsE164(*r.PhoneNumber.Value) {
				return fmt.Errorf("phone number does not match e164 format: %w", core_errors.ErrInvalidArgument)
			}
		}
	}

	if r.Email.Set {
		_, err := mail.ParseAddress(*r.Email.Value)
		if err != nil {
			return fmt.Errorf("invalid email format: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

// PatchUser godoc
// @Summary Patches user
// Description Patches user in db
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param patch body PatchUserRequest true "PatchRequestBody"
// @Success 200 {object} PatchUserResponse "Patched user response"
// @Failure 400 {object} transport_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} transport_http_response.ErrorResponse "Conflict"
// @Failure 404 {object} transport_http_response.ErrorResponse "Not found"
// @Failure 500 {object} transport_http_response.ErrorResponse "Internal server error"
// @Router /users/{id} [patch]
func (t *UserTransport) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := transport_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.RespondError("failed to get userID", err)
		return
	}

	var patchRequest PatchUserRequest

	if err := core_http_request.DecodeAndValidate(r, &patchRequest); err != nil {
		responseHandler.RespondError("failed to decode request: %w", err)
		return
	}

	if err := patchRequest.Validate(); err != nil {
		responseHandler.RespondError("failed to validate request: %w", err)
		return
	}

	userPatch := userPatchFromRequest(patchRequest)
	user, err := t.UserService.PatchUser(ctx, userID, userPatch)

	if err != nil {
		responseHandler.RespondError("failed to patch user", err)
		return
	}

	responseHandler.RespondJSON(http.StatusOK, PatchUserResponse(dtoFromDomain(user)))
}

func userPatchFromRequest(patchRequest PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		patchRequest.FullName.ToDomain(),
		patchRequest.Password.ToDomain(),
		patchRequest.PhoneNumber.ToDomain(),
		patchRequest.Email.ToDomain(),
	)
}
