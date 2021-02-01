package graphql

import (
	"context"
	"easyfood/pkg/graphql/gqlgen"
	"easyfood/pkg/graphql/models"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{}

func NewMutationResolver() gqlgen.MutationResolver {
	return new(mutationResolver)
}

func (m mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	return models.NewUser(0), nil
}

func (m mutationResolver) Auth(ctx context.Context, input models.AuthInput) (string, error) {
	// buscar no banco email e senha, filtrar por email
	err := bcrypt.CompareHashAndPassword([]byte(input.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Email), 0)
	return string(hash), err
}