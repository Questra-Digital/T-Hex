package graph


import (
	"Settings/graph/model"
	s "Settings/service"
	"context"
	"fmt"
)


func (r *mutationResolver) SetSettings(ctx context.Context, input model.GetSettingInput) (*model.Settings, error) {
	fmt.Printf("SetSettings")
	return s.SetSettings(ctx, input)
}


func (r *mutationResolver) SetProjectInfo(ctx context.Context, input model.ProjectInfoInput) (*model.Projectinfos, error) {
	fmt.Printf("SetProjectInfo")
	return s.SetProjectInfo(ctx, input)
}


func (r *queryResolver) Settings(ctx context.Context) ([]*model.Settings, error) {
	return s.GetAllSettings(ctx)
}


func (r *queryResolver) Setting(ctx context.Context, id int) (*model.Settings, error) {
	return s.GetSettingByID(ctx, id)
}


func (r *queryResolver) ProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	return s.GetAllProjectInfos(ctx)
}


func (r *queryResolver) ProjectInfo(ctx context.Context, id int) (*model.Projectinfos, error) {
	return s.GetProjectInfoByID(ctx, id)
}


func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }


func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
