package randomuseragent

import "testing"

func TestGetRandomUserAgent(t *testing.T) {
	if got := GetRandomUserAgent(); got == "" {
		t.Errorf("GetRandomUserAgent() is empty string!")
	}
}
