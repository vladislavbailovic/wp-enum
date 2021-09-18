package enum

import (
	"testing"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateJsonRoute(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	res, err := enumerateJsonRoute("http://multiwp.test/calendar/")(client)
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["bog"]
	if !exists {
		t.Fatalf("expected user bog to exist")
	}
}
