package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"discussion/graph/generated"
	"discussion/graph/model"
	"fmt"
)

func (r *entityResolver) FindDiscussionByID(ctx context.Context, id uint64) (*model.Discussion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindPostByID(ctx context.Context, id uint64) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindUserByID(ctx context.Context, id uint64) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
