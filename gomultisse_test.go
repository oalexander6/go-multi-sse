package multisse_test

import (
	"testing"

	"github.com/oalexander-dev/go-multi-sse"
)

func TestNewManager(t *testing.T) {
	mgr := multisse.New()

	if len(mgr.Streams) != 0 {
		t.Error("Unexpected streams after initializing manager")
	}
}

func TestGetHeaders(t *testing.T) {
	results := multisse.GetSSEHeaders()

	for _, hdr := range results {
		var expected string

		switch hdr.Key {
		case "Content-Type":
			expected = "text/event-stream"
		case "Cache-Control":
			expected = "no-cache"
		case "Connection":
			expected = "keep-alive"
		case "Transfer-Encoding":
			expected = "chunked"
		default:
			t.Error("Unexpected header present")
		}

		if hdr.Val != expected {
			t.Errorf("Unexpected value for header %s: %s", hdr.Key, hdr.Val)
		}
	}
}
