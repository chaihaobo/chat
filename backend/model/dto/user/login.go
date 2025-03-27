package user

type (
	LoginRequest struct {
		Code string `json:"code" validate:"required"`
	}

	LoginByPasswordRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		ID           uint64 `json:"id"`
		Avatar       string `json:"avatar"`
		UserName     string `json:"username"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
