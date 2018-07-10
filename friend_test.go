package neteaseim_test

import "testing"

func Test_AddFriend(t *testing.T) {
	im := setupTestCase(t)
	// appKey := "d45545b3eeb821970eab26931859871e"
	// appSecret := "d31182026a36"
	// im := Init(appKey, appSecret)

	im.AddFriend("helloworld", "helloworld2", 1, "")
}
