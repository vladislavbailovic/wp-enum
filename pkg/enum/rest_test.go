package enum

import "testing"

func TestEnumerateAuthorId(t *testing.T) {
	res, err := enumerateAuthorId("http://multiwp.test/calendar/")
	if err != nil {
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["bog"]
	if !exists {
		t.Fatalf("expected user bog to exist")
	}
}
