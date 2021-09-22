package http

import (
	"net/http"
	"testing"
)

func TestPassthroughClientAddsCookies(t *testing.T) {
	ck := http.Cookie{Name: "test", Value: "what"}
	client := NewHttpClient(CLIENT_PASSTHROUGH)
	if len(client.GetCookies()) != 0 {
		t.Fatalf("should have no cookies by default")
	}

	client.AddCookie(&ck)
	if len(client.GetCookies()) != 0 {
		t.Fatalf("should still have no cookies")
	}

	if client.HasCookies() {
		t.Fatalf("bool method should match state: 0")
	}
}

func TestWebClientAddsCookies(t *testing.T) {
	ck := http.Cookie{Name: "test", Value: "what"}
	client := NewHttpClient(CLIENT_WEB)
	if len(client.GetCookies()) != 0 {
		t.Fatalf("should have no cookies by default")
	}
	if client.HasCookies() {
		t.Fatalf("bool method should match state: 0")
	}

	client.AddCookie(&ck)
	if len(client.GetCookies()) != 1 {
		t.Fatalf("should have one cookie added")
	}
	if !client.HasCookies() {
		t.Fatalf("bool method should match state: 1")
	}
}

func TestAddMockWpCookies(t *testing.T) {
	client := NewHttpClient(CLIENT_WEB)
	AddMockWPCookies(client)

	if len(client.GetCookies()) != 3 {
		t.Fatalf("should have fake wp cookies added")
	}
	t.Log(client.GetCookies())
}
