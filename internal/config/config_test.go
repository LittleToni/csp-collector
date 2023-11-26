package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expectedPort := "1234"

	t.Run("Port is set by environment", func(t *testing.T) {
		os.Setenv("PORT", expectedPort)
		defer os.Unsetenv("PORT")

		cfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("LoadConfig() error %v, want error %v", err, false)
		}
		if cfg.Port != expectedPort {
			t.Errorf("LoadConfig() = %v, want %v", cfg.Port, expectedPort)
		}
	})

	t.Run("Port is set by fallback", func(t *testing.T) {
		os.Unsetenv("PORT")

		cfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("LoadConfig() error %v, want error %v", err, false)
		}
		if cfg.Port != "8080" {
			t.Errorf("LoadConfig() default port = %v, want %v", cfg.Port, "8080")
		}
	})
}
