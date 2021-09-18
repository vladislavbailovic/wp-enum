package enum

import (
	"fmt"
	"testing"
	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateApiPassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	_, err := enumerateJsonApi("http://multiwp.test/calendar/")(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestEnumerateJsonApiSuccess(t *testing.T) {
	address := "127.0.0.1:6666"
	serverCloser := fakeJsonApiSuccessServer(address, jsonSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	if res[0].Name != "admin" {
		t.Fatalf("expected user admin to exist")
	}
}

func TestEnumerateJsonApiFailure(t *testing.T) {
	address := "127.0.0.1:6669"
	serverCloser := fakeJsonApiSuccessServer(address, jsonFailure())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	_, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}
}
