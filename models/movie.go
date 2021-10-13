package models

import "time"

type Movie struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Genres    string    `json:"genres"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (movie *Movie) TableName() string {
	return "movie"
}

func (movie *Movie) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = movie.ID
	response["name"] = movie.Name
	response["genres"] = movie.Genres
	return response
}
