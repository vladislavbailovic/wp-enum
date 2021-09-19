package http

import (
	"net/http"
	"testing"
)

func TestUaIsNotRandomByDefault(t *testing.T) {
	ua := UserAgent{}
	if ua.isRandom {
		t.Fatalf("should not be random by default")
	}
}

func TestDefaultUserAgentIsConstant(t *testing.T) {
	ua := UserAgent{false}
	if ua.Agent() != DEFAULT_USER_AGENT {
		t.Fatalf("expected default user agent")
	}
}

func TestRandomUserAgentDoesNotReturnDefault(t *testing.T) {
	ua := UserAgent{true}
	if ua.Agent() == DEFAULT_USER_AGENT {
		t.Fatalf("expected non-default user agent")
	}
}

func TestRandomUserAgentSetsUniqueUaHeader(t *testing.T) {
	var req *http.Request
	old := DEFAULT_USER_AGENT
	ua := UserAgent{true}

	for i := 0; i < 5; i++ {
		req, _ = http.NewRequest("GET", "test.com", nil)
		ua.SetHeader(req)
		if req.UserAgent() == old {
			t.Fatalf("expected unique UA, got %s", old)
		}
		old = req.UserAgent()
	}
}

func TestNewRandomUAIsRandom(t *testing.T) {
	ua := NewRandomUA()
	if !ua.isRandom {
		t.Fatalf("new random UA should be random")
	}
}
