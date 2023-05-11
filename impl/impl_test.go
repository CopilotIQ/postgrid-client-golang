package impl

import (
	"github.com/jgroeneveld/trial/assert"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

const ApiKeyEnvKey = "POST_GRID_API_KEY"

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file: %s", err)
	}
}

func TestNew(t *testing.T) {
	t.Run("verify live_ initializes with live=true", func(t *testing.T) {
		// Test case 1: apiKey starts with "live_"
		apiKey1 := "live_randomstring"
		pg := New(apiKey1)
		assert.True(t, pg.live)
	})
	t.Run("verify test_ initializes with live=false", func(t *testing.T) {
		// Test case 2: apiKey starts with "test_"
		apiKey2 := "test_randomstring"
		pg := New(apiKey2)
		assert.False(t, pg.live)
	})
	t.Run("verify live=false for random API key", func(t *testing.T) {
		// Test case 3: apiKey starts with a random string
		apiKey3 := "randomstring"
		pg := New(apiKey3)
		assert.False(t, pg.live)
	})
	t.Run("verify local API key is a test key", func(t *testing.T) {
		// Test case 3: apiKey starts with a random string
		apiKey3 := os.Getenv(ApiKeyEnvKey)
		assert.NotEqual(t, "", apiKey3)
		pg := New(apiKey3)
		assert.False(t, pg.live)
	})
}
