package enum

import "testing"

func TestEnumerateJsonApi(t *testing.T) {
	res, err := enumerateJsonApi("http://multiwp.test/calendar/")
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["bog"]
	if !exists {
		t.Fatalf("expected user bog to exist")
	}
}
