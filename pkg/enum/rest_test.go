package enum

import (
	"testing"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateAuthorId(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	res, err := enumerateAuthorId("http://multiwp.test/calendar/")(client)
	if err != nil {
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["bog"]
	if !exists {
		t.Fatalf("expected user bog to exist")
	}
}
