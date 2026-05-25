package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Empty Header Case", func(t *testing.T) {
		h := http.Header{}
		_, err := GetAPIKey(h)
		if err == nil {
			t.Error("expected error to exist, but got nil")
		}
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Errorf("expected error to be of type ErrNoAuthHeaderIncluded, but got: %v", err)
		}
	})

	t.Run("Incorrect Header Case", func(t *testing.T) {
		h := http.Header{"Authorization": []string{"wrongKeyName", "key"}}
		_, err := GetAPIKey(h)
		if err == nil {
			t.Error("expected error but received nil")
		}
		if !strings.Contains(err.Error(), "malformed") {
			t.Errorf("expected error to contain 'malformed' but got: %v", err)
		}
	})

	t.Run("Correct Header Case", func(t *testing.T) {
		h := http.Header{"Authorization": []string{"ApiKey TheKey"}}
		key, err := GetAPIKey(h)
		if err != nil {
			t.Errorf("expected success but got: %v", err)
		}
		if key != "TheKey" {
			t.Errorf("expected to get key: 'TheKey', but got: %s", key)
		}
	})
}
