package enum

import (
	"fmt"
	"testing"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateRestPassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	_, err := enumerateAuthorId("http://multiwp.test/calendar/")(client)
	if err == nil {
		t.Fatalf("expected error to be nil")
	}
}

func TestEnumerateRestSuccess(t *testing.T) {
	address := "127.0.0.1:6666"
	serverCloser := fakeJsonApiSuccessServer(address, restSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateAuthorId(fmt.Sprintf("http://%s/", address))(client)
	if err != nil {
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["admin"]
	if !exists {
		t.Fatalf("expected user admin to exist")
	}
}
