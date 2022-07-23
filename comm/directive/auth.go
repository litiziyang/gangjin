package directive

import (
	"comm/auth"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
)

func Auth() func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		user := auth.User(ctx)
		if user == nil {
			return nil, fmt.Errorf("用户未登录")
		}
		return next(ctx)
	}
}
