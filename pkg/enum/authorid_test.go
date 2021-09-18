package enum

import (
	"fmt"
	"testing"
	"wp-enum/pkg/data"
	wp_http "wp-enum/pkg/http"
)

func TestEnumerateRestPassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	authors, err := enumerateAuthorId("http://multiwp.test/calendar/")(client, data.DefaultConstraints())
	if err != nil {
		t.Fatalf("expected error to not be nil")
	}
	if len(authors) > 0 {
		t.Fatalf("expected no results")
	}
}

func TestEnumerateRestSuccess(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, restSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateAuthorId(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Fatalf("expected error to be nil")
	}

	if res[0].Name != "admin" {
		t.Fatalf("expected user admin to exist")
	}
}

func TestEnumerateRestFailsWithLimit(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, restSuccess())
	defer serverCloser.Close()

	opts := data.DefaultConstraints()
	opts.End = 5
	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateAuthorId(fmt.Sprintf("http://%s/", address))(client, opts)
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	if len(res) > 0 {
		t.Fatalf("expected no users to be found within limit")
	}
}
