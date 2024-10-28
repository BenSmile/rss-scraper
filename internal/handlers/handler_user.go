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

func (h *Handler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when parsing JSON: %v", err))
		return
	}

	user, err := h.apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("error when creating user: %v", err))
		return
	}

	RespondWithJson(w, 201, models.DbUser2User(user))

}

func (h *Handler) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	RespondWithJson(w, 200, models.DbUser2User(user))

}

func (h *Handler) HandlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {

	postsDb, err := h.apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		RespondWithError(w, 400, "could not fetch posts")
		return
	}
	RespondWithJson(w, 200, models.DbPosts2Posts(postsDb))

}
