package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bensmile/rssaggregator/internal/config"
	"github.com/bensmile/rssaggregator/internal/database"
	"github.com/bensmile/rssaggregator/internal/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	apiCfg *config.ApiConfig
}

func NewHandler(apiCfg *config.ApiConfig) *Handler {
	return &Handler{
		apiCfg,
	}
}

func (h *Handler) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when parsing JSON: %v", err))
		return
	}

	feedFollow, err := h.apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when creating feed follow: %v", err))
		return
	}

	RespondWithJson(w, 201, models.DbFeedFollow2FeedFollow(feedFollow))

}

func (h *Handler) HandlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedsDB, err := h.apiCfg.DB.GetFeedsFollows(r.Context(), user.ID)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when fetching feeds: %v", err))
		return
	}

	RespondWithJson(w, 200, models.DbFeedFollows2FeedFollows(feedsDB))
}

func (h *Handler) HandlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIdStr := chi.URLParam(r, "feed_follow_id")

	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("could not parse feed follow uuid: %v", err))
		return
	}

	if err := h.apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowId,
	}); err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when unfollowing feed: %v", err))
		return
	}

	RespondWithJson(w, 204, struct{}{})
}
