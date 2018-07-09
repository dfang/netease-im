package netease-im

import "testing"

func setupTestCase(t *testing.T) *NeteaseIM {
	t.Log("setup test case")

	appKey := "d45545b3eeb821970eab26931859871e"
	appSecret := "d31182026a36"
	im := Init(appKey, appSecret)

	return im
}

func Test_CreateAccid(t *testing.T) {
	im := setupTestCase(t)

	info := userInfo{
		accid: "helloworld",
		name:  "zhangsan",
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

	info := userInfo{
		accid:  "helloworld",
		name:   "王二麻子",
		sign:   "我是王二麻子",
		gender: 0,
	}

	im.UpdateInfo(info)
}

func Test_GetUserInfos(t *testing.T) {
	im := setupTestCase(t)

	accids := []string{"helloworld"}

	im.GetUserinfos(accids)
}
