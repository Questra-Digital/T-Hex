package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"Settings/graph/model"
	"context"
	"fmt"
)

// SetSettings is the resolver for the setSettings field.
func (r *mutationResolver) SetSettings(ctx context.Context, input model.GetSettingtInput) (*model.Settings, error) {
	panic(fmt.Errorf("not implemented: SetSettings - setSettings"))
}

// Settings is the resolver for the settings field.
func (r *queryResolver) Settings(ctx context.Context) ([]*model.Settings, error) {
	panic(fmt.Errorf("not implemented: Settings - settings"))
}

// Setting is the resolver for the setting field.
func (r *queryResolver) Setting(ctx context.Context, id int) (*model.Settings, error) {
	panic(fmt.Errorf("not implemented: Setting - setting"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
