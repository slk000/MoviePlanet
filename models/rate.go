package models

import "time"

type Rate struct {
	ID        int64 `gorm:"primary_key;auto_increment" json:"id"`
	User      User
	UserID    int64 `json:"user_id"`
	Movie     Movie
	MovieID   int64     `json:"movie_id"`
	Comment   string    `json:"comment"`
	Score     int64     `json:"score"`
	Status    int64     `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (rate *Rate) TableName() string {
	return "rate"
}

func (rate *Rate) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = rate.ID
	response["uid"] = rate.UserID
	response["mid"] = rate.MovieID
	response["commnet"] = rate.Comment
	response["score"] = rate.Score
	response["status"] = rate.Status
	return response
}
