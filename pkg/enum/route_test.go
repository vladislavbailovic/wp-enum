package enum

import (
	"fmt"
	"testing"
	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateRoutePassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	_, err := enumerateJsonRoute("http://multiwp.test/calendar/")(client, data.DefaultConstraints())
	if err == nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}
}

func TestEnumerateJsonRouteSuccess(t *testing.T) {
	address := "127.0.0.1:6667"
	serverCloser := fakeJsonApiSuccessServer(address, jsonSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateJsonRoute(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["admin"]
	if !exists {
		t.Fatalf("expected user admin to exist")
	}
}
