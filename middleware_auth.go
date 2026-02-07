package main


import (
	"net/http"
	"fmt"
	"github.com/KaiyueWei/rssagg/internal/database"
	"github.com/KaiyueWei/rssagg/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Counldn't get the user: %v", err))
		}
		handler(w, r, user)
	}
}
