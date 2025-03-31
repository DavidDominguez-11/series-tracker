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

// para los demas endpoints

func IncrementEpisode(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]

    // Obtener el último episodio visto y el total de episodios
    var lastEp int
    var totalEp *int
    row := db.DB.QueryRow("SELECT last_episode_watched, total_episodes FROM series WHERE id = ?", id)
    if err := row.Scan(&lastEp, &totalEp); err != nil {
        http.Error(w, "Serie no encontrada", http.StatusNotFound)
        return
    }

    // Verificar si se puede incrementar
    if totalEp != nil && lastEp >= *totalEp {
        http.Error(w, "No se puede incrementar el episodio", http.StatusBadRequest)
        return
    }

    // Incrementar el último episodio visto
    _, err := db.DB.Exec("UPDATE series SET last_episode_watched = last_episode_watched + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Episodio incrementado"})
}

func UpvoteRanking(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    _, err := db.DB.Exec("UPDATE series SET ranking = COALESCE(ranking, 0) + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Upvote exitoso"})
}

func DownvoteRanking(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    _, err := db.DB.Exec("UPDATE series SET ranking = COALESCE(ranking, 0) - 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Downvote exitoso"})
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    
    var requestBody struct {
        Status string `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
        return
    }

    // Validar el estado
    validStatuses := map[string]bool{
        "Plan to Watch": true,
        "Watching":      true,
        "Dropped":       true,
        "Completed":     true,
    }
    if !validStatuses[requestBody.Status] {
        http.Error(w, "Estado inválido", http.StatusBadRequest)
        return
    }

    _, err := db.DB.Exec("UPDATE series SET status = ? WHERE id = ?", requestBody.Status, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Estado actualizado"})
}