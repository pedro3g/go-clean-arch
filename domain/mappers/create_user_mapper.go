package mappers

import (
	"github.com/pedro3g/go-clean-arch/domain/entities"
	"github.com/pedro3g/go-clean-arch/shared/inputs"
)

func CreateUserMapper(i inputs.CreateUserInput) entities.User {
	return entities.User{
		Name:     i.Name,
		Email:    i.Email,
		Password: i.Password,
	}
}
