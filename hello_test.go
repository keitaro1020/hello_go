package hello_go

import (
	"testing"
)

func TestHello(t *testing.T) {

	hello := Hello()

	if hello != "hello go world" {
		t.Error("not 'hello go world'")
	}
}
