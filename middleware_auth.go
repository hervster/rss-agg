package main

import (
	"fmt"
	"net/http"

	"github.com/hervster/rss-agg/internal/database"
	"github.com/hervster/rss-agg/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey) // In go there's a context pkg in std library,
		// gives you a way to track something that's happening across multiple goroutines,
		// you can cancel context which would kill the http request
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %s", err))
			return
		}

		handler(w, r, user)
	}
}
