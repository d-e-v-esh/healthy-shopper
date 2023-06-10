package healthyshopper

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
)

type StoreKey struct{}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, userInfo RegisterInput) (*UserResponse, error) {

	// Create a new user
	user, err := r.client.User.
		Create().
		SetUsername(userInfo.Username).
		SetEmailAddress(userInfo.EmailAddress).
		SetPassword(userInfo.Password).
		SetFirstName(userInfo.FirstName).
		SetLastName(userInfo.LastName).
		Save(ctx)

	redisStore := ctx.Value("redisStore").(*RedisStore)

	redisStore.SetCookie("Register Cookie")

	// store, ok := ctx.Value(StoreKey{}).(RedisStore)

	// store.SetCookie("Register Cookie")
	// if !ok {
	// 	// Handle case where there's no Redis store in the context.
	// 	log.Fatal("No Redis store in context")
	// }

	return &UserResponse{user, []*Error{}}, err
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, userInfo LoginInput) (*UserResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
