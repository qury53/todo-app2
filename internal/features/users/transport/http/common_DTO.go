package users_transport_http

import "github.com/qury53/todo-app2/internal/core/domain"

type UserDTO struct {
	ID          int     `json:"id" example:"1"`
	Version     int     `json:"version" example:"2"`
	FullName    string  `json:"full_name" example:"Ivan Ivanov"`
	Password    string  `json:"password" example:"qwerty123"`
	PhoneNumber *string `json:"phone_number" example:"+79998887766"`
	Email       string  `json:"email" example:"somemail@mail.ru"`
}

func dtoFromDomain(domain domain.User) UserDTO {
	return UserDTO{
		ID:          domain.ID,
		Version:     domain.Version,
		FullName:    domain.FullName,
		Password:    domain.Password,
		PhoneNumber: domain.PhoneNumber,
		Email:       domain.Email,
	}
}

func dtosFromDomains(domains []domain.User) []UserDTO {
	dtos := make([]UserDTO, len(domains))
	for i, domain := range domains {
		dtos[i] = dtoFromDomain(domain)
	}

	return dtos
}
