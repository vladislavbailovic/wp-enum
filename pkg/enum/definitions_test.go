package enum

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
