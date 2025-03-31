package handlers

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"series-tracker/db"
	"series-tracker/models"
	"github.com/gorilla/mux"
)

func GetSeries(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, status, last_episode_watched, total_episodes, ranking FROM series")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var seriesList []models.Series
	for rows.Next() {
		var s models.Series
		if err := rows.Scan(&s.ID, &s.Title, &s.Status, &s.LastEpisodeWatched, &s.TotalEpisodes, &s.Ranking); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		seriesList = append(seriesList, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seriesList)
}

func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := db.DB.QueryRow("SELECT id, title, status, last_episode_watched, total_episodes, ranking FROM series WHERE id = ?", id)

	var s models.Series
	if err := row.Scan(&s.ID, &s.Title, &s.Status, &s.LastEpisodeWatched, &s.TotalEpisodes, &s.Ranking); err != nil {
		http.Error(w, "Serie no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func CreateSeries(w http.ResponseWriter, r *http.Request) {
	var s models.Series
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	// Insertar serie en la base de datos
	result, err := db.DB.Exec("INSERT INTO series (title, status, last_episode_watched, total_episodes, ranking) VALUES (?, ?, ?, ?, ?)",
		s.Title, s.Status, s.LastEpisodeWatched, s.TotalEpisodes, s.Ranking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func UpdateSeries(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.Series
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE series SET title = ?, status = ?, last_episode_watched = ?, total_episodes = ?, ranking = ? WHERE id = ?",
		s.Title, s.Status, s.LastEpisodeWatched, s.TotalEpisodes, s.Ranking, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func DeleteSeries(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.DB.Exec("DELETE FROM series WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
