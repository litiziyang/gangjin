package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"comm/auth"
	"comm/logger"
	"context"
	"errors"
	"fmt"
	"security/proto/token"
	"security/proto/verify"
	"user/graph/generated"
	"user/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) SendCode(ctx context.Context, phone *string) (string, error) {
	sendVerify, err := r.VerifyService.SendVerify(ctx, &verify.VerifyRequest{Phone: *phone})
	fmt.Println("shibia")
	if err != nil {
		if sendVerify != nil && sendVerify.Error != "" {
			graphql.AddError(ctx, errors.New("sendVerify.Error"))
		}
		logger.Error(r.Log, "短信服务调用失败", err, sendVerify)
		return "验证码发送失败", err
	}
	return "验证码发送成功", nil
}

func (r *mutationResolver) Login(ctx context.Context, phone string, password *string, code *string) (*model.LoginResponse, error) {
	user := &model.User{
		Phone: phone,
	}
	if code != nil {
		//验证码登录 ，判断验证码是否正确
		res, err := r.VerifyService.CheckVerify(ctx, &verify.VerifyRequest{Phone: phone, Code: *code})
		fmt.Println(res)
		if err != nil || !res.Check {
			graphql.AddError(ctx, err)
			logger.Error(r.Log, "验证码错误", err, nil)
			return nil, err
		}
		if err := r.Db.Where(user).FirstOrCreate(user).Error; err != nil {
			graphql.AddError(ctx, err)
			logger.Error(r.Log, "已经创建过了", err, nil)
			return nil, err
		}
	} else if password != nil {
		//密码登录判断是否正确
		if err := r.Db.Where(user).FirstOrCreate(user).Error; err != nil {
			graphql.AddError(ctx, errors.New("密码错误"))
			logger.Error(r.Log, "未注册，请先注册", err, user)
			return nil, err
		}
		if user.Password == " " {
			graphql.AddError(ctx, errors.New("密码未设置"))
			logger.Error(r.Log, "密码还没创建过了", errors.New("密码还没创建"), user)
			return nil, nil
		}
		err := bcrypt.CompareHashAndPassword([]byte((*user).Password), []byte(*password))
		if err != nil {
			graphql.AddError(ctx, err)
			logger.Error(r.Log, "密码错误", err, user)
			return nil, err
		}
	} else {
		graphql.AddError(ctx, errors.New("用户名或者密码错误"))
		return nil, nil
	}

	security, err := r.SecurityService.GetSecurityToken(ctx,
		&token.TokenRequest{User: &token.TokenUser{
			Id: user.ID,
		}})
	if err != nil {
		graphql.AddError(ctx, err)
		logger.Error(r.Log, "服务出错了", err, user)
		return nil, err
	}
	return &model.LoginResponse{
		User:  user,
		Token: security.Token,
	}, nil
}

func (r *mutationResolver) SetPassword(ctx context.Context, oldPassword string, newPassword string) (string, error) {
	u := auth.User(ctx)
	user := &model.User{}
	user.ID = u.Id
	if err := r.Db.First(user).Error; err != nil {
		graphql.AddError(ctx, errors.New("查询失败"))
		logger.Error(r.Log, "失败", err, user)
		return "密码修改失败", err
	}
	if len(user.Password) != 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
		if err != nil {
			graphql.AddError(ctx, errors.New("密码错误"))
			logger.Error(r.Log, "密码错误", err, user)
			return "密码错误", err
		}
		errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newPassword))
		if errs == nil {
			graphql.AddError(ctx, errors.New("新密码不能与旧密码一致"))
			logger.Error(r.Log, "新密码不能与旧密码一致", errs, user)
			return "新密码不能与旧密码一致", errs
		}
	}
	password, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		graphql.AddError(ctx, errors.New("修改失败"))
		logger.Error(r.Log, "密码加密失败", err, user)
		return "密码加密失败", err
	}
	pass := string(password)
	user.Password = pass
	r.Db.Save(user)
	return "修改密码成功", nil
}

func (r *queryResolver) User(ctx context.Context, phone string) (*model.User, error) {
	User := &model.User{}
	User.Phone = phone
	if err := r.Db.Find(User).Error; err != nil {
	}
	return nil, nil
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Password(ctx context.Context, password string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Phone(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) Password(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
