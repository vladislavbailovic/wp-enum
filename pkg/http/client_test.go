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

func TestPtClientUaIsNotRandom(t *testing.T) {
	pc := NewHttpClient(CLIENT_PASSTHROUGH)
	if pc.GetAgent().isRandom {
		t.Fatalf("passthrough client UA should not be random")
	}

	ua := UserAgent{true}
	pc.SetAgent(&ua)
	if pc.GetAgent().isRandom {
		t.Fatalf("passthrough client UA should not be random even if set")
	}
}

func TestWebClientUaRandomn(t *testing.T) {
	client := NewHttpClient(CLIENT_WEB)
	if client.GetAgent().isRandom {
		t.Fatalf("web client UA should not be random by default")
	}

	ua := UserAgent{true}
	client.SetAgent(&ua)
	if !client.GetAgent().isRandom {
		t.Fatalf("web client UA should be random when requested")
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
