package handlers

import (
	"fmt"
	"net/http"

	"github.com/bensmile/rssaggregator/internal/auth"
	"github.com/bensmile/rssaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (h *Handler) MiddlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := h.apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
