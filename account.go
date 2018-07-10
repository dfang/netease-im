package neteaseim

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// UserInfo 创建或更新云通信ID时需用到
type UserInfo struct {
	Accid  string
	Name   string
	Icon   string
	Sign   string
	Email  string
	Birth  string
	Mobile string
	Gender int
	Ex     string
	Props  string
	Token  string
}

// http://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID
// http://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户名片

// CreateAccid 创建网易云通信ID
func (im *NeteaseIM) CreateAccid(info UserInfo) string {
	CreateAccidURL := NETEASE_BASE_URL + "/user/create.action"

	params := url.Values{}
	if len(info.Accid) <= 0 {
		panic("accid 参数必须")
	} else {
		params.Set("accid", info.Accid)
	}

	if len(info.Token) > 0 {
		params.Set("token", info.Token)
	}
	if len(info.Name) > 0 {
		params.Set("name", info.Name)
	}

	if len(info.Props) > 0 {
		params.Set("props", info.Props)
	}

	if len(info.Icon) > 0 {
		params.Set("icon", info.Icon)
	}

	if len(info.Sign) > 0 {
		params.Set("sign", info.Sign)
	}

	if len(info.Email) > 0 {
		params.Set("email", info.Email)
	}

	if len(info.Birth) > 0 {
		params.Set("birth", info.Birth)
	}

	if len(info.Mobile) > 0 {
		params.Set("mobile", info.Mobile)
	}

	if info.Gender < 2 {
		params.Set("gender", strconv.Itoa(info.Gender))
	}

	if len(info.Ex) > 0 {
		params.Set("ex", info.Ex)
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
func (im *NeteaseIM) UpdateInfo(info UserInfo) string {
	UpdateInfoURL := "https://api.netease.im/nimserver/user/updateUinfo.action"
	params := url.Values{}
	params.Set("accid", info.Accid)
	params.Set("name", info.Name)
	params.Set("icon", info.Icon)
	params.Set("sign", info.Sign)
	params.Set("email", info.Email)
	params.Set("birth", info.Birth)
	params.Set("mobile", info.Mobile)
	params.Set("gender", strconv.Itoa(info.Gender))
	params.Set("ex", info.Ex)
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
