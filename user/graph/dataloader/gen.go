package dataloader

import "context"

//go:generate go run github.com/vektah/dataloaden UserLoader uint64 *user/graph/model.User

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
