package auth_test

import (
	"testing"

	"github.io/taserbeat/line-family-bot/modules/auth"
)

var client auth.Authenticate

func TestIsFamily(t *testing.T) {

	userId1 := "abc"
	if !client.IsFamily(userId1) {
		t.Errorf("%v is expected family, but actual is not family.", userId1)
	}

	userId2 := "aaa"
	if client.IsFamily(userId2) {
		t.Errorf("%v is expected not family, but actual is family.", userId2)
	}
}

func init() {
	client = auth.New("abc,def,ghi")
}
