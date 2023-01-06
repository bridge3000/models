package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"
)

//构造签名类
type SignUtil struct {
}

func (this *SignUtil) Md5Hash(plainText string) string {
	var h = md5.New()
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}

//大王支付交互的签名，可以整合到下面的函数里
func (this *SignUtil) MakeDawangSign(gameId int32, serverId int32, accountId int64, roleId string, ttp int64, secret string) string {
	keySlice := []string{"game_id", "server_id", "account_id", "role_id", "ttp", "secret"}
	sign_str := ""
	comma := ""
	params := map[string]string{
		"game_id":    strconv.Itoa(int(gameId)),
		"server_id":  strconv.Itoa(int(serverId)),
		"account_id": strconv.FormatInt(int64(accountId), 10),
		"role_id":    roleId,
		"ttp":        strconv.FormatInt(ttp, 10),
		"secret":     secret,
	}

	for _, mapKey := range keySlice {
		sign_str += comma + mapKey + "=" + params[mapKey]
		comma = "&"
	}

	var buf = bytes.Buffer{}
	buf.WriteString(sign_str)

	var h = md5.New()
	h.Write(buf.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}

func (this *SignUtil) MakeSign(params map[string]interface{}, appKey string) string {
	sign_str := ""
	comma := ""

	//按key排序
	keySlice := []string{}
	for k, _ := range params {
		keySlice = append(keySlice, k)
	}
	sort.Strings(keySlice)

	for _, mapKey := range keySlice {
		/* 去掉为空的和不参与签名的字段 */
		if mapKey != "sign" && mapKey != "KEY" {
			value, ok := params[mapKey].(string)
			if ok {
				sign_str += comma + mapKey + "=" + value
				comma = "&"
			}
		}
	}

	var buf = bytes.Buffer{}
	buf.WriteString(sign_str + appKey)

	return this.Md5Hash(this.Md5Hash(buf.String()))
}
