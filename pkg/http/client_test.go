package http

import (
	"testing"
)

func TestNewHttpClientReturnsPassthroughClientByDefault(t *testing.T) {
	var nc Client
	nc = NewHttpClient()

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client by default")
	}

	_, ok2 := nc.(PassthroughClient)
	if !ok2 {
		t.Fatalf("should be passthrough client by default")
	}
}

func TestNewHttpClientReturnsPassthroughClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewHttpClient(CLIENT_PASSTHROUGH)

	_, ok1 := nc.(WebClient)
	if ok1 {
		t.Fatalf("should not be web client when passthrough requested")
	}

	_, ok2 := nc.(PassthroughClient)
	if !ok2 {
		t.Fatalf("should be passthrough client when requested")
	}
}

func TestNewHttpClientReturnsWebClientWhenRequested(t *testing.T) {
	var nc Client
	nc = NewHttpClient(CLIENT_WEB)

	_, ok1 := nc.(PassthroughClient)
	if ok1 {
		t.Fatalf("should not be passthrough client when web requested")
	}

	_, ok2 := nc.(WebClient)
	if !ok2 {
		t.Fatalf("should be web client when requested")
	}
}

func TestPassthroughClientSend(t *testing.T) {
	nc := NewHttpClient(CLIENT_PASSTHROUGH)

	resp := nc.Send("whatever.com")
	if resp.StatusCode != -1 {
		t.Fatalf("expected negative status code")
	}
}

func TestWebClientSend(t *testing.T) {
	nc := NewHttpClient(CLIENT_WEB)

	resp := nc.Send("whatever.com")
	t.Log(resp)
}
