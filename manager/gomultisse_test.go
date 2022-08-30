package gomultisse_test

import (
	"testing"

	"github.com/oalexander-dev/go-multi-sse/manager"
)

func TestNewManager(t *testing.T) {
	mgr := gomultisse.New()

	if len(mgr.Streams) != 0 {
		t.Error("Unexpected streams after initializing manager")
	}
}
