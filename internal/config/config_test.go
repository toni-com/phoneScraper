package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("ValidJSON", func(t *testing.T) {
		const validConfig = `
[
  {
    "name": "Smartphone A",
    "url": "https://example.com/a",
    "selector": ".price-a",
    "threshold": 999.00
  },
  {
    "name": "Smartphone B",
    "url": "https://example.com/b",
    "selector": ".price-b",
    "threshold": 500.50
  }
]
`
		// Create a temporary file to simulate the config.json
		tmpfile, err := os.CreateTemp("", "valid-config-*.json")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpfile.Name()) // Clean up the file afterward

		if _, err := tmpfile.Write([]byte(validConfig)); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatalf("Failed to close temp file: %v", err)
		}

		items, err := LoadConfig(tmpfile.Name())

		if err != nil {
			t.Errorf("LoadConfig failed unexpectedly with error: %v", err)
			return
		}

		if len(items) != 2 {
			t.Fatalf("Expected 2 items, got %d", len(items))
		}

		if items[0].Name != "Smartphone A" || items[0].Threshold != 999.00 {
			t.Errorf("Item 1 parsing failed. Got: %+v", items[0])
		}
		if items[1].Name != "Smartphone B" || items[1].Threshold != 500.50 {
			t.Errorf("Item 2 parsing failed. Got: %+v", items[1])
		}
	})

	t.Run("InvalidSyntax", func(t *testing.T) {
		const invalidConfig = `[{"name": "Bad Item", "threshold": 1000.00,}` // Trailing comma makes it invalid JSON

		tmpfile, err := os.CreateTemp("", "invalid-config-*.json")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write([]byte(invalidConfig)); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		if err := tmpfile.Close(); err != nil {
			t.Fatalf("Failed to close temp file: %v", err)
		}

		_, err = LoadConfig(tmpfile.Name())

		// We expect an error here! Robust code should return one.
		if err == nil {
			t.Errorf("LoadConfig succeeded unexpectedly for invalid JSON syntax. Expected error.")
		}
	})

	// --- Case 3: Missing File ---
	t.Run("MissingFile", func(t *testing.T) {
		// Attempt to load a file we know doesn't exist
		_, err := LoadConfig("nonexistent-file.json")

		// We expect an error here as well (file I/O error).
		if err == nil {
			t.Errorf("LoadConfig succeeded unexpectedly for a missing file. Expected error.")
		}
	})
}
