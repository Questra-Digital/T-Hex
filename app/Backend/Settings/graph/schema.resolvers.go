package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"Settings/graph/model"
	s "Settings/service"
	"context"
	"fmt"
)

// SetSettings is the resolver for the setSettings field.
func (r *mutationResolver) SetSettings(ctx context.Context, input model.GetSettingInput) (*model.Settings, error) {
	fmt.Printf("SetSettings")
	return s.SetSettings(ctx, input)
}

// SetProjectInfo is the resolver for the setProjectInfo field.
func (r *mutationResolver) SetProjectInfo(ctx context.Context, input model.ProjectInfoInput) (*model.Projectinfos, error) {
	fmt.Printf("SetProjectInfo")
	return s.SetProjectInfo(ctx, input)
}

// Settings is the resolver for the settings field.
func (r *queryResolver) Settings(ctx context.Context) ([]*model.Settings, error) {
	return s.GetAllSettings(ctx)
}

// Setting is the resolver for the setting field.
func (r *queryResolver) Setting(ctx context.Context, id int) (*model.Settings, error) {
	return s.GetSettingByID(ctx, id)
}

// ProjectInfos is the resolver for the projectInfos field.
func (r *queryResolver) ProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	return s.GetAllProjectInfos(ctx)
}

// ProjectInfo is the resolver for the projectInfo field.
func (r *queryResolver) ProjectInfo(ctx context.Context, id int) (*model.Projectinfos, error) {
	return s.GetProjectInfoByID(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
