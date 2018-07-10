package neteaseim_test

import (
	"testing"

	. "github.com/dfang/neteaseim"
)

func setupTestCase(t *testing.T) *NeteaseIM {
	t.Log("setup test case")

	appKey := "d45545b3eeb821970eab26931859871e"
	appSecret := "d31182026a36"
	im := Init(appKey, appSecret)

	return im
}

func Test_CreateAccid(t *testing.T) {
	im := setupTestCase(t)

	info := UserInfo{
		Accid: "helloworld",
		Name:  "zhangsan",
	}

	t.Log("Creating accid")
	result := im.CreateAccid(info)
	t.Log(result)
}

func Test_UpdateAccid(t *testing.T) {
	im := setupTestCase(t)

	accid := "helloworld"
	name := "李四"

	im.UpdateAccid(accid, name)
}

func Test_RefreshToken(t *testing.T) {
	im := setupTestCase(t)

	accid := "helloworld"
	name := "李四"

	im.RefreshToken(accid, name)
}

func Test_BlockAccid(t *testing.T) {
	im := setupTestCase(t)

	accid := "helloworld"
	im.BlockAccid(accid)
}

func Test_UnBlockAccid(t *testing.T) {
	im := setupTestCase(t)

	accid := "helloworld"
	im.UnBlockAccid(accid)
}

func Test_UpdateInfo(t *testing.T) {
	im := setupTestCase(t)

	info := UserInfo{
		Accid:  "helloworld",
		Name:   "王二麻子",
		Sign:   "我是王二麻子",
		Gender: 0,
	}

	im.UpdateInfo(info)
}

func Test_GetUserInfos(t *testing.T) {
	im := setupTestCase(t)

	accids := []string{"helloworld"}

	im.GetUserinfos(accids)
}
