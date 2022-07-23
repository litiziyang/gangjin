package dataloader

import (
	"gorm.io/gorm"
	"post/graph/model"
	"time"
)

func GetPostLoader(db *gorm.DB) *PostLoader {
	return &PostLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []uint64) ([]*model.Post, []error) {
			var posts []*model.Post
			if err := db.Find(&posts, ids).Error; err != nil {
				return nil, nil
			}
			postById := map[uint64]*model.Post{}
			for _, post := range posts {
				postById[post.ID] = post
			}
			result := make([]*model.Post, len(ids))
			for i, id := range ids {
				result[i] = postById[id]
				i++
			}
			return result, nil
		}}
}
