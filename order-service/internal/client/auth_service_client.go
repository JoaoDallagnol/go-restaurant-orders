package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/JoaoDallagnol/go-restaurant-orders/order-service/internal/config"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type AuthClient interface {
	GetUserById(userID uint) (UserResponse, error)
}

type authClient struct {
	baseURL    string
	endpoint   string
	httpClient *http.Client
}

func NewAuthClient(cfg *config.Config) AuthClient {
	return &authClient{
		baseURL:  cfg.AuthService.BaseURL,
		endpoint: cfg.AuthService.Endpoint.GetUserById,
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.AuthService.ConnectTimeout) * time.Second,
		},
	}
}

func (m *authClient) GetUserById(userId uint) (UserResponse, error) {
	url := fmt.Sprintf("%s%s", m.baseURL, strings.Replace(m.endpoint, ":id", fmt.Sprintf("%d", userId), 1))

	resp, err := m.httpClient.Get(url)
	if err != nil {
		return UserResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UserResponse{}, fmt.Errorf("failed to fetch user %d, status: %d", userId, resp.StatusCode)
	}

	var user UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return UserResponse{}, err
	}

	return user, nil
}
