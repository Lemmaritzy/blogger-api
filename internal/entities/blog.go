package entities

import "github.com/jackc/pgtype"

type Blog struct {
	Id          int64            `json:"id"`
	Creator     string           `json:"creator"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Content     string           `json:"content"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
}
