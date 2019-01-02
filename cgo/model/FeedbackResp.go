package model

import "grass/cgo/entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}
