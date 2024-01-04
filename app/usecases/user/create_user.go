package usecases

import (
	"github.com/pedro3g/go-clean-arch/domain/entities"
	"github.com/pedro3g/go-clean-arch/domain/errors"
	"github.com/pedro3g/go-clean-arch/domain/mappers"
	"github.com/pedro3g/go-clean-arch/shared/inputs"
)

func CreateUser(input inputs.CreateUserInput) errors.Either {
	user := mappers.CreateUserMapper(input)

	return errors.Right[entities.User]{Value: user}
}
