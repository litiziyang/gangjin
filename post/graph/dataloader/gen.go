package dataloader

import "context"

//go:generate go run github.com/vektah/dataloaden PostLoader uint64 *post/graph/model.Post

const loadersKey = "dataloaders"

type Loaders struct {
	PostById PostLoader
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
