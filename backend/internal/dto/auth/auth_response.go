package auth

import model "sociomile-apps/internal/models"

type AuthResponse struct {
	AccessToken string      `json:"access_token"`
	User        *model.User `json:"user"`
}
