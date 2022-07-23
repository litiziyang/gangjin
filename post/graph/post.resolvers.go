package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"comm/auth"
	"comm/logger"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"post/graph/generated"
	"post/graph/model"
	"security/proto/hot"
)

func (r *mutationResolver) Write(ctx context.Context, title string, argument string, describe string) (string, error) {
	u := auth.User(ctx)
	post := &model.Post{}
	post.UserId = u.Id
	post.Title = title
	post.Argument = argument
	post.Describe = describe
	post.Enable = true
	if err := r.Db.Create(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("创建失败"))
		logger.Error(r.Log, "创建失败", err, post)
		return "创建失败", err
	}

	//创建es索引
	postJson, err := json.Marshal(post)
	if err != nil {
		graphql.AddError(ctx, errors.New("创建json失败"))
		logger.Error(r.Log, "创建json失败", err, post)
		return "创建失败", err
	}
	postjs := string(postJson)
	_, err = r.Es.Index().Index("post").BodyJson(postjs).Do(ctx)
	if err != nil {
		graphql.AddError(ctx, errors.New("创建索引失败"))
		logger.Error(r.Log, "创建索引失败", err, post)
		return "创建索引失败", err
	}
	_, err = r.HotServise.GetSecurityHot(ctx, &hot.HotRequest{
		Id:        post.ID,
		Time:      post.CreatedAt.String(),
		Number:    100,
		ModelName: "post",
	})
	if err != nil {
		return "", err
	}
	return "创建成功", nil

}

func (r *mutationResolver) Delete(ctx context.Context, postID uint64) (string, error) {
	u := auth.User(ctx)
	post := &model.Post{}
	post.ID = postID
	if err := r.Db.Where(&model.Post{Enable: true}).Find(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("并没有文章"))
		logger.Error(r.Log, "删除失败", err, post)
		return " ", err
	}
	if post.UserId != u.Id {
		graphql.AddError(ctx, errors.New("你没权利删除"))
		logger.Error(r.Log, "删除失败", nil, post)
		return " ", errors.New("傻逼")
	}
	post.Enable = false
	r.Db.Save(post)
	return "成功", nil
}

func (r *mutationResolver) SeeMyPost(ctx context.Context) ([]*model.Post, error) {
	u := auth.User(ctx)
	var posts []*model.Post
	if err := r.Db.Where(&model.Post{UserId: u.Id, Enable: true}).
		Limit(5).Find(&posts).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有记录"))
		logger.Error(r.Log, "没有", err, nil)
		return nil, err
	}
	return posts, nil
}

func (r *mutationResolver) Post(ctx context.Context, id uint64) (*model.Post, error) {
	post := &model.Post{}
	post.ID = id
	if err := r.Db.Where(&model.Post{Enable: true}).Find(post).Error; err != nil {
		graphql.AddError(ctx, errors.New("没有记录"))
		logger.Error(r.Log, "没有", err, nil)
		return nil, err
	}
	return post, nil
}

func (r *postResolver) Heat(ctx context.Context, obj *model.Post) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
