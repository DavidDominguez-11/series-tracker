package models

type Series struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Status             string  `json:"status"`
	LastEpisodeWatched int     `json:"last_episode_watched"`
	TotalEpisodes      *int    `json:"total_episodes"`
	Ranking            *int    `json:"ranking"`
}
