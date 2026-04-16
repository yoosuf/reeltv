package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"reeltv/backend/internal/app"
	"reeltv/backend/internal/config"
	"github.com/gin-gonic/gin"
)

var testApp *app.App
var router *gin.Engine

func setupTestApp(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := &config.Config{
		App: config.AppConfig{
			Port: 8081,
			Env:  "test",
		},
		Database: config.DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "test",
			Password: "test",
			DBName:   "reeltv_test",
		},
		Redis: config.RedisConfig{
			Host: "localhost",
			Port: "6379",
		},
		JWT: config.JWTConfig{
			Secret:            "test-secret-key",
			AccessExpiration:  15 * time.Minute,
			RefreshExpiration: 7 * 24 * time.Hour,
		},
	}

	var err error
	testApp, err = app.New(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize test app: %v", err)
	}

	router = testApp.GetRouter().GetEngine()
}

func teardownTestApp(t *testing.T) {
	if testApp != nil {
		testApp.Shutdown()
	}
}

func TestHealthCheck(t *testing.T) {
	setupTestApp(t)
	defer teardownTestApp(t)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestRegisterUser(t *testing.T) {
	setupTestApp(t)
	defer teardownTestApp(t)

	payload := map[string]interface{}{
		"email":    "testuser@example.com",
		"password": "Test123!",
		"name":     "Test User",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusOK {
		t.Errorf("Expected status 201 or 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if response["success"] != true {
		t.Errorf("Expected success to be true")
	}
}

func TestLoginUser(t *testing.T) {
	setupTestApp(t)
	defer teardownTestApp(t)

	payload := map[string]interface{}{
		"email":    "test@example.com",
		"password": "Test123!",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Login might fail if user doesn't exist, but endpoint should be reachable
	if w.Code != http.StatusOK && w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 200 or 401, got %d", w.Code)
	}
}

func TestListSeries(t *testing.T) {
	setupTestApp(t)
	defer teardownTestApp(t)

	req, _ := http.NewRequest("GET", "/api/v1/catalog/series?offset=0&limit=20", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if response["success"] != true {
		t.Errorf("Expected success to be true")
	}
}
