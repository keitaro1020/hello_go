package hello_go

import (
	"testing"
)

func TestHello(t *testing.T) {

	ret := Hello()

	if ret != "hello go world" {
		t.Error("not 'hello go world'")
	}
}
