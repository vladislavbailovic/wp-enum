package enum

import (
	"fmt"
	"testing"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateApiPassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	_, err := enumerateJsonApi("http://multiwp.test/calendar/")(client)
	if err == nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}
}

func TestEnumerateJsonApiSuccess(t *testing.T) {
	address := ":6666"
	serverCloser := fakeJsonApiSuccessServer(address, jsonSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client)
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["admin"]
	if !exists {
		t.Fatalf("expected user admin to exist")
	}
}
