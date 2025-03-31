package models

import (
	"gorm.io/gorm"
)

type Series struct {
	gorm.Model
	Id uint `gorm:"primaryKey" json:"id"`
	Title             string `gorm:"size:100;not null;unique" json:"title"`
	Status            string `gorm:"type:enum('Plan to Watch','Watching','Dropped','Completed');not null" json:"status"`
	LastEpisodeWatched int    `gorm:"default:0" json:"last_episode_watched"`
	TotalEpisodes     *int   `gorm:"default:NULL" json:"total_episodes"`
	Ranking           *int   `gorm:"default:NULL" json:"ranking"`
}

// SeriesFilters contiene los parámetros de filtrado
type SeriesFilters struct {
	Title  string `form:"title"`
	Status string `form:"status"`
	Sort   string `form:"sort"` // "asc" o "desc"
}

type ApiResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Series  *Series `json:"series,omitempty"`
	Data    any     `json:"data,omitempty"`
}