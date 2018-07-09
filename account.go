package netease-im

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type userInfo struct {
	accid  string
	name   string
	icon   string
	sign   string
	email  string
	birth  string
	mobile string
	gender int
	ex     string
	props  string
	token  string
}

// http://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID
// http://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户名片

// CreateAccid 创建网易云通信ID
func (im *NeteaseIM) CreateAccid(info userInfo) string {
	CreateAccidURL := NETEASE_BASE_URL + "/user/create.action"

	params := url.Values{}
	if len(info.accid) <= 0 {
		panic("accid 参数必须")
	} else {
		params.Set("accid", info.accid)
	}

	if len(info.token) > 0 {
		params.Set("token", info.token)
	}
	if len(info.name) > 0 {
		params.Set("name", info.name)
	}

	if len(info.props) > 0 {
		params.Set("props", info.props)
	}

	if len(info.icon) > 0 {
		params.Set("icon", info.icon)
	}

	if len(info.sign) > 0 {
		params.Set("sign", info.sign)
	}

	if len(info.email) > 0 {
		params.Set("email", info.email)
	}

	if len(info.birth) > 0 {
		params.Set("birth", info.birth)
	}

	if len(info.mobile) > 0 {
		params.Set("mobile", info.mobile)
	}

	if info.gender < 2 {
		params.Set("gender", strconv.Itoa(info.gender))
	}

	if len(info.ex) > 0 {
		params.Set("ex", info.ex)
	}
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(CreateAccidURL, postData)
	return result
}

// UpdateAccid 更新网易云通信ID
func (im *NeteaseIM) UpdateAccid(accid string, name string) string {

	UpdateAccidURL := "https://api.netease.im/nimserver/user/update.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("name", name)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(UpdateAccidURL, postData)
	return result
}

// RefreshToken 更新并获取新token
func (im *NeteaseIM) RefreshToken(accid string, name string) string {
	RefreshTokenURL := "https://api.netease.im/nimserver/user/refreshToken.action"

	params := url.Values{}
	params.Set("accid", accid)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(RefreshTokenURL, postData)
	return result
}

// BlockAccid 封禁网易云通信ID
func (im *NeteaseIM) BlockAccid(accid string) string {
	BlockAccidURL := "https://api.netease.im/nimserver/user/block.action"

	params := url.Values{}
	params.Set("accid", accid)

	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(BlockAccidURL, postData)
	return result
}

// UnBlockAccid 解禁网易云通信ID
func (im *NeteaseIM) UnBlockAccid(accid string) string {
	UnblockAccidURL := "https://api.netease.im/nimserver/user/unblock.action"

	params := url.Values{}
	params.Set("accid", accid)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(UnblockAccidURL, postData)
	return result
}

//UpdateInfo 更新用户名片
func (im *NeteaseIM) UpdateInfo(info userInfo) string {
	UpdateInfoURL := "https://api.netease.im/nimserver/user/updateUinfo.action"
	params := url.Values{}
	params.Set("accid", info.accid)
	params.Set("name", info.name)
	params.Set("icon", info.icon)
	params.Set("sign", info.sign)
	params.Set("email", info.email)
	params.Set("birth", info.birth)
	params.Set("mobile", info.mobile)
	params.Set("birth", info.birth)
	params.Set("gender", strconv.Itoa(info.gender))
	params.Set("ex", info.ex)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(UpdateInfoURL, postData)
	return result
}

func (im *NeteaseIM) GetUserinfos(accids []string) string {
	params := url.Values{}
	jsonarr, _ := json.Marshal(accids)

	params.Set("accids", string(jsonarr))
	postData := strings.NewReader(params.Encode())

	GetUserInfosURL := "https://api.netease.im/nimserver/user/getUinfos.action"
	result, _ := im.DoNeteaseImRequest(GetUserInfosURL, postData)
	return result
}
