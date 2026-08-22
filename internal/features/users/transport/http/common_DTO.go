package users_transport_http

import "github.com/qury53/todo-app2/internal/core/domain"

type UserDTO struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	Password    string  `json:"password"`
	PhoneNumber *string `json:"phone_number"`
	Email       string  `json:"email"`
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
