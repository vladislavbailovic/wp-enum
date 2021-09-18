package main

import "testing"

func TestEnumerateReturnsErrorWithInvalidEnumType(t *testing.T) {
	for i := 10; i < 15; i++ {
		res, err := Enumerate(EnumerationType(i), "test.com")
		if res != nil {
			t.Fatalf("expected nil result for invalid enumeration type: %d", i)
		}
		if err == nil {
			t.Fatalf("expected error for invalid enumeration type")
		}
	}
}

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

func TestEnumerateJsonRoute(t *testing.T) {
	res, err := enumerateJsonRoute("http://multiwp.test/calendar/")
	if err != nil {
		t.Log(err)
		t.Fatalf("expected error to be nil")
	}

	_, exists := res["bog"]
	if !exists {
		t.Fatalf("expected user bog to exist")
	}
}
