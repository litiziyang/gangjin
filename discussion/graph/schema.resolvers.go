package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"comm/auth"
	"comm/logger"
	"context"
	"discussion/graph/generated"
	"discussion/graph/model"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func (r *discussionResolver) Hot(ctx context.Context, obj *model.Discussion) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *discussionResolver) Factions(ctx context.Context, obj *model.Discussion) (model.DiscussionType, error) {
	panic(fmt.Errorf("not implemented"))
}

//加入阵营
func (r *mutationResolver) ChoiceFactions(ctx context.Context, postID uint64) (string, error) {
	post := &model.Post{
		ID: postID,
	}
	if err := r.Db.Where(&model.Post{Enable: true}).Find(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有这篇文章"))
		logger.Error(r.Log, "没有这篇文章", err, nil)
		return "", nil
	}
	u := auth.User(ctx)
	dis := &model.Discussion{PostId: postID, UserId: u.Id}
	if err := r.Db.FirstOrCreate(dis).Error; err != nil {
		graphql.AddError(ctx, err)
		logger.Error(r.Log, "出错", err, dis)
		return "", err
	}

	return "欢迎你的加入", nil
}

//写评论
func (r *mutationResolver) Write(ctx context.Context, postID uint64, discussionID *uint64) (string, error) {
	post := &model.Post{
		ID: postID,
	}
	if err := r.Db.Where(&model.Post{Enable: true}).Find(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有这篇文章"))
		logger.Error(r.Log, "没有这篇文章", err, nil)
		return "", nil
	}
	u := auth.User(ctx)
	dis := &model.Discussion{DiscussionId: *discussionID, PostId: postID, UserId: u.Id}
	if err := r.Db.Create(dis).Error; err != nil {
		graphql.AddError(ctx, errors.New("创建失败"))
		logger.Error(r.Log, "创建失败", err, nil)
		return "创建失败", err
	}
	return "创建成功", nil

}

//删除自己的评论
func (r *mutationResolver) Delete(ctx context.Context, discussionID uint64) (string, error) {
	u := auth.User(ctx)
	dis := &model.Discussion{}
	dis.ID = discussionID
	dis.UserId = u.Id
	if err := r.Db.Delete(dis).Error; err != nil {
		graphql.AddError(ctx, errors.New("删除失败"))
		logger.Error(r.Log, "删除失败", err, dis)
		return "删除失败", err
	}
	return "删除成功", nil

}

//查看文章的评论
func (r *queryResolver) LookPostDiscussion(ctx context.Context, postID uint64) ([]*model.Discussion, error) {
	post := &model.Post{
		ID: postID,
	}
	var dis []*model.Discussion
	if err := r.Db.Where(&model.Post{Enable: true}).Find(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有这篇文章"))
		logger.Error(r.Log, "没有这篇文章", err, nil)
		return nil, nil
	}
	if err := r.Db.Where(&model.Discussion{PostId: postID}).Find(dis).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有评论"))
		logger.Error(r.Log, "没有评论", err, nil)
		return nil, err
	}
	return dis, nil

}

//查看自己的评论
func (r *queryResolver) CheckMyDiscussion(ctx context.Context) ([]*model.Discussion, error) {
	u := auth.User(ctx)
	var dis []*model.Discussion
	if err := r.Db.Where(&model.Discussion{UserId: u.Id}).Find(dis).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有评论"))
		logger.Error(r.Log, "没有评论", err, nil)
		return nil, err
	}
	return dis, nil
}

// Discussion returns generated.DiscussionResolver implementation.
func (r *Resolver) Discussion() generated.DiscussionResolver { return &discussionResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type discussionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
