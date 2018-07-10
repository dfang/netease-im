package neteaseim

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// http://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1

func (im *NeteaseIM) AddFriend(accid, faccid string, add_type int, msg string) {
	AddFriendURL := "https://api.netease.im/nimserver/friend/add.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("faccid", faccid)
	params.Set("type", strconv.Itoa(add_type))
	params.Set("msg", msg)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(AddFriendURL, postData)
	fmt.Println(result)
}

func (im *NeteaseIM) UpdateFriend(accid, faccid, alias, ex, serverex string) {
	UpdateFriendURL := "https://api.netease.im/nimserver/friend/update.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("faccid", faccid)
	params.Set("alias", alias)
	params.Set("ex", ex)
	params.Set("serverex", serverex)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(UpdateFriendURL, postData)
	fmt.Println(result)
}

func (im *NeteaseIM) DeleteFriend(accid, faccid string) {
	DeleteFriendURL := "https://api.netease.im/nimserver/friend/delete.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("faccid", faccid)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(DeleteFriendURL, postData)
	fmt.Println(result)
}

func (im *NeteaseIM) GetFriend(accid string, updatetime int64, createtime int64) {
	GetFriendURL := "https://api.netease.im/nimserver/friend/get.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("updatetime", strconv.FormatInt(updatetime, 10))
	params.Set("createtime", strconv.FormatInt(createtime, 10))
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(GetFriendURL, postData)
	fmt.Println(result)
}

func (im *NeteaseIM) BlockFriend(accid, targetAcc string, relationType int, value int) {
	BlockFriendURL := "https://api.netease.im/nimserver/user/setSpecialRelation.action"
	params := url.Values{}
	params.Set("accid", accid)
	params.Set("targetAcc", targetAcc)
	params.Set("relationType", strconv.Itoa(relationType))
	params.Set("value", strconv.Itoa(value))
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(BlockFriendURL, postData)
	fmt.Println(result)
}

func (im *NeteaseIM) GetBlockList(accid string) {
	BlockFriendURL := "https://api.netease.im/nimserver/user/listBlackAndMuteList.action"
	params := url.Values{}
	params.Set("accid", accid)
	postData := strings.NewReader(params.Encode())

	result, _ := im.DoNeteaseImRequest(BlockFriendURL, postData)
	fmt.Println(result)
}
