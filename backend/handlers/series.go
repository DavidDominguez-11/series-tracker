package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"series-tracker-backend/models"
)

func GetSeries(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filters models.SeriesFilters
		if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
			respondWithError(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var series []models.Series
		query := db.Model(&models.Series{})

		if filters.Title != "" {
			query = query.Where("title LIKE ?", "%"+filters.Title+"%")
		}

		if filters.Status != "" {
			query = query.Where("status = ?", filters.Status)
		}

		if filters.Sort == "asc" {
			query = query.Order("ranking ASC")
		} else {
			query = query.Order("ranking DESC")
		}

		if err := query.Find(&series).Error; err != nil {
			respondWithError(w, "Failed to fetch series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Series retrieved successfully",
			Data:    series,
		})
	}
}

func GetSeriesByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		var series models.Series
		if err := db.First(&series, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				respondWithError(w, "Series not found", http.StatusNotFound)
				return
			}
			respondWithError(w, "Failed to fetch series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Series retrieved successfully",
			Series:  &series,
		})
	}
}

func CreateSeries(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var series models.Series
		if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
			respondWithError(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if series.Title == "" || series.Status == "" {
			respondWithError(w, "Title and status are required", http.StatusBadRequest)
			return
		}

		if err := db.Create(&series).Error; err != nil {
			respondWithError(w, "Failed to create series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Series created successfully",
			Series:  &series,
		}, http.StatusCreated)
	}
}

// Implementaciones similares para UpdateSeries, DeleteSeries, UpdateStatus, etc.

func UpdateSeries(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		var series models.Series
		if err := db.First(&series, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				respondWithError(w, "Series not found", http.StatusNotFound)
				return
			}
			respondWithError(w, "Failed to fetch series", http.StatusInternalServerError)
			return
		}

		var updateData models.Series
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			respondWithError(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		series.Title = updateData.Title
		series.Status = updateData.Status
		series.LastEpisodeWatched = updateData.LastEpisodeWatched
		series.TotalEpisodes = updateData.TotalEpisodes
		series.Ranking = updateData.Ranking

		if err := db.Save(&series).Error; err != nil {
			respondWithError(w, "Failed to update series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Series updated successfully",
			Series:  &series,
		})
	}
}

func DeleteSeries(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		if err := db.Delete(&models.Series{}, id).Error; err != nil {
			respondWithError(w, "Failed to delete series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Series deleted successfully",
		}, http.StatusNoContent)
	}
}

func UpdateStatus(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		var req struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if req.Status == "" {
			respondWithError(w, "Status is required", http.StatusBadRequest)
			return
		}

		result := db.Model(&models.Series{}).Where("id = ?", id).Update("status", req.Status)
		if result.Error != nil {
			respondWithError(w, "Failed to update status", http.StatusInternalServerError)
			return
		}

		if result.RowsAffected == 0 {
			respondWithError(w, "Series not found", http.StatusNotFound)
			return
		}

		var series models.Series
		if err := db.First(&series, id).Error; err != nil {
			respondWithError(w, "Failed to fetch updated series", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Status updated successfully",
			Series:  &series,
		})
	}
}

func IncrementEpisode(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		var series models.Series
		if err := db.First(&series, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				respondWithError(w, "Series not found", http.StatusNotFound)
				return
			}
			respondWithError(w, "Failed to fetch series", http.StatusInternalServerError)
			return
		}

		// Solo incrementar si no hay total o si no hemos llegado al final
		if series.TotalEpisodes == nil || series.LastEpisodeWatched < *series.TotalEpisodes {
			series.LastEpisodeWatched++
			if err := db.Save(&series).Error; err != nil {
				respondWithError(w, "Failed to increment episode", http.StatusInternalServerError)
				return
			}
		} else {
			respondWithError(w, "No episodes to increment", http.StatusBadRequest)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Episode incremented successfully",
			Series:  &series,
		})
	}
}

func Upvote(db *gorm.DB) http.HandlerFunc {
	return handleVote(db, 1)
}

func Downvote(db *gorm.DB) http.HandlerFunc {
	return handleVote(db, -1)
}

func handleVote(db *gorm.DB, delta int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			respondWithError(w, "Invalid series ID", http.StatusBadRequest)
			return
		}

		var series models.Series
		if err := db.First(&series, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				respondWithError(w, "Series not found", http.StatusNotFound)
				return
			}
			respondWithError(w, "Failed to fetch series", http.StatusInternalServerError)
			return
		}

		if series.Ranking == nil {
			initialRank := 0
			series.Ranking = &initialRank
		}

		newRank := *series.Ranking + delta
		series.Ranking = &newRank

		if err := db.Save(&series).Error; err != nil {
			respondWithError(w, "Failed to update ranking", http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, models.ApiResponse{
			Success: true,
			Message: "Ranking updated successfully",
			Series:  &series,
		})
	}
}