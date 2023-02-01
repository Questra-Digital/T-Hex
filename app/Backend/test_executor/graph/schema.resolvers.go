package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"test_executor/graph/model"
	"test_executor/service"
)

// StartTest is the resolver for the startTest field.
func (r *mutationResolver) StartTest(ctx context.Context, input model.StartTestInput) (*model.Test, error) {
	return service.StartTest(ctx,input);
}

// Tests is the resolver for the tests field.
func (r *queryResolver) Tests(ctx context.Context) ([]*model.Test, error) {
	return service.GetAllTest(ctx);
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context, id int) (*model.Test, error) {
	return service.GetTestByID(ctx,id);
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
