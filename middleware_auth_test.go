package main

/*
import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

// In your test file
type mockDB struct {
	// store test user data
	user database.User
}

func (m *mockDB) GetUser(ctx context.Context, apiKey string) (database.User, error) {
	// return our predetermined user for successful case
	if apiKey == "valid-test-key" {
		return m.user, nil
	}
	return database.User{}, errors.New("user not found")
}

func TestMiddlewareAuth(t *testing.T) {
	// Setup mock config
	mockConfig := &apiConfig{
		DB: &mockDB{
			user: database.User{
				ID:        "test-user-id",
				CreatedAt: time.Now().UTC().Format(time.RFC3339),
				UpdatedAt: time.Now().UTC().Format(time.RFC3339),
				Name:      "test-user",
				ApiKey:    "valid-test-key",
			},
		},
	}

	// Create test request with header
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "ApiKey valid-test-key")

	mockConfig.middlewareAuth(func(w http.ResponseWriter, r *http.Request, user database.User) {
		if user.ID != "test-user-id" {
			t.Errorf("unexpected user ID: %s", user.ID)
		}
	})
}
;
*/
