package neteaseim_test

import (
	"testing"

	. "github.com/dfang/neteaseim"
)

func Test_Init(t *testing.T) {
	appKey := "d45545b3eeb821970eab26931859871e"
	appSecret := "d31182026a36"
	client := Init(appKey, appSecret)

	if client.AppKey != appKey || client.AppSecret != appSecret {
		t.Error("Initialization error")
	}
}
