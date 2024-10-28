package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bensmile/rssaggregator/internal/database"
	"github.com/bensmile/rssaggregator/internal/models"
	"github.com/google/uuid"
)

func (h *Handler) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when parsing JSON: %v", err))
		return
	}

	feed, err := h.apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when creating user: %v", err))
		return
	}

	RespondWithJson(w, 201, models.DbFeed2Feed(feed))

}

func (h *Handler) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feedsDB, err := h.apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when fetching feeds: %v", err))
		return
	}

	RespondWithJson(w, 200, models.DbFeeds2Feeds(feedsDB))
}
