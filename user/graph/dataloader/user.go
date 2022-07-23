package dataloader

import (
	"gorm.io/gorm"
	"time"
	"user/graph/model"
)

func GetUserLoader(db *gorm.DB) *UserLoader {
	return &UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []uint64) ([]*model.User, []error) {
			var users []*model.User
			if err := db.Find(&users, ids).Error; err != nil {
				return nil, nil
			}
			userById := map[uint64]*model.User{}
			for _, user := range users {
				userById[user.ID] = user
			}
			result := make([]*model.User, len(ids))
			for i, id := range ids {
				result[i] = userById[id]
				i++
			}
			return result, nil
		},
	}
}
