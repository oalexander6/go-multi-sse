package headers_test

import (
	"testing"

	"github.com/oalexander-dev/go-multi-sse/headers"
)

func TestGetHeaders(t *testing.T) {
	results := headers.GetSSEHeaders()

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
