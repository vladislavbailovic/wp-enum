package enum

import (
	"fmt"
	"testing"
	"wp-user-enum/pkg/data"
	wp_http "wp-user-enum/pkg/http"
)

func TestEnumerateApiPassthrough(t *testing.T) {
	client := wp_http.NewHttpClient(wp_http.CLIENT_PASSTHROUGH)
	_, err := enumerateJsonApi("http://multiwp.test/calendar/")(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestEnumerateJsonApiSuccess(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, jsonSuccess())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	res, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	if res[0].Username != "admin" {
		t.Fatalf("expected user admin to exist")
	}
}

func TestEnumerateJsonApiFailure(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, jsonFailure())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	_, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestEnumerateJsonApiFailureWithDefaultUa(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, jsonFailureForDefaultUa())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	_, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}

	ua := wp_http.NewRandomUA()
	client.SetAgent(&ua)
	_, err = enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Fatalf("expected success")
	}
}

func TestEnumerateJsonApiFailuerWithoutWpCookies(t *testing.T) {
	address := getListenerAddress()
	serverCloser := fakeJsonApiSuccessServer(address, jsonFailureWithNoWpCookies())
	defer serverCloser.Close()

	client := wp_http.NewHttpClient(wp_http.CLIENT_WEB)
	_, err := enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err == nil {
		t.Fatalf("expected error")
	}

	wp_http.AddMockWPCookies(client)
	_, err = enumerateJsonApi(fmt.Sprintf("http://%s/", address))(client, data.DefaultConstraints())
	if err != nil {
		t.Fatalf("expected success")
	}
}
