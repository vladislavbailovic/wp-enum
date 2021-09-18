package http

import "testing"

func TestNormalizeUrlAssumesProtocol(t *testing.T) {
	suite := map[string]string{
		"1312.net/something": "http://1312.net/something",
		"//161.com?test=11":  "http://161.com?test=11",
		"https://test.com":   "https://test.com",
		"http://nanana.test": "http://nanana.test",
	}
	for test, expected := range suite {
		actual := NormalizeUrl(test)
		if actual != expected {
			t.Fatalf("expected %s but got %s", expected, actual)
		}
	}
}

func TestNormalizeRootUrlAddsTrailingSlash(t *testing.T) {
	suite := map[string]string{
		"1312.net":            "http://1312.net/",
		"//161.com/":          "http://161.com/",
		"https://test.com///": "https://test.com/",
		"http://nanana.test":  "http://nanana.test/",
	}
	for test, expected := range suite {
		actual := NormalizeRootUrl(test)
		if actual != expected {
			t.Fatalf("expected %s but got %s", expected, actual)
		}
	}
}
