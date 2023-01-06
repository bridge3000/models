package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"../Adoter_Asset"
	"github.com/golang/protobuf/proto"
)

const (
	CONTENT_TYPE_X_WWW_FORM_URLENCODED = 1
	CONTENT_TYPE_JSON                  = 2
)

type NetworkUtil struct {
}

var ResponseErrors = []string{
	"成功",
	"没有权限",
	"参数错误",
	"没有账号数据",
	"未找到玩家",
	"玩家在线",
	"玩家离线",
	"物品未能找到",
	"未能找到服务器",
	"未能找到相关数据，由于策划配置造成",
	"邮件ID非法",
	"奖励ID非法",
	"物品ID非法",
	"不能删除",
	"排行榜未找到",
	"充值项未找到",
	"活动未找到",
	"名字非法",
}

//发送POST请求 用的时候重新整理代码
//@param url
//@param sendType发送方式 1form 2json
//@param data map[string]interface{} 因为要兼容发送JSON格式，所以用interface，但是纯文本发送的必须map里必须是string，断言了
//@param sign
//@return body
//@return err
func (this *NetworkUtil) HttpPost(url string, sendType int, dataMap map[string]interface{}) (body []byte, err error) {
	contentType := ""
	var postData io.Reader

	if sendType == 1 {
		contentType = "application/x-www-form-urlencoded"

		slice1 := []string{}
		for k, v := range dataMap {
			value, ok := v.(string)
			if ok {
				slice1 = append(slice1, k+"="+value)
			}
		}

		postData = strings.NewReader(strings.Join(slice1, "&"))
	} else if sendType == 2 {
		contentType = "application/json;charset=utf-8"

		b, err := json.Marshal(dataMap)
		if err != nil {
			return nil, err
		}

		postData = bytes.NewBuffer(b)
	}

	resp, err := http.Post(url, contentType, postData)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//将map数据转化为  可发送的 httpbody
func (this *NetworkUtil) BuildHttpBody(contentType string, dataMap map[string]interface{}) (io.Reader, error) {
	var requestBody io.Reader

	if contentType == "application/x-www-form-urlencoded" {
		slice1 := []string{}
		for k, v := range dataMap {
			value, ok := v.(string)
			if ok {
				slice1 = append(slice1, k+"="+value)
			}
		}

		requestBody = strings.NewReader(strings.Join(slice1, "&"))
	} else if contentType == "application/json;charset=utf-8" {
		b, err := json.Marshal(dataMap)
		if err != nil {
			return nil, err
		}

		requestBody = bytes.NewBuffer(b)
	}

	return requestBody, nil
}

func (this *NetworkUtil) HttpRequest(method string, url string, headers map[string]string, requestBody io.Reader) (responseHeader http.Header, body []byte, err error, returnCode int) {
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, nil, err, 0
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err, 0
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err, 0
	}

	return resp.Header, body, nil, resp.StatusCode
}

// 给游戏服务器发消息
// @param curZone大区对象
// @param 协议类型
// @param 数据包
// @return responseStuff 字节数组，需要反序列化成具体的对象, err 错误
func (this *NetworkUtil) SendMessageToGameServer(ipAddress string, linkPort int, typeT Adoter_Asset.GameMessageDefine, stuff proto.Message) (responseBytes []byte, err error) {
	arrayUtil := ArrayUtil{}
	stuffMarshalData, _ := proto.Marshal(stuff)

	meta := &Adoter_Asset.Meta{}
	meta.TypeT = uint32(typeT) //协议号
	meta.Stuff = stuffMarshalData
	marshalMeta, err := proto.Marshal(meta)

	if err != nil {

	} else {
		var buf = make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, uint32(len(marshalMeta)))

		var buf2 = make([]byte, 4) //协议类型
		binary.LittleEndian.PutUint32(buf2, uint32(meta.TypeT))
		byteData := arrayUtil.BytesCombine(buf, buf2, marshalMeta)

		//		connTimeout := 5 * time.Second
		connTimeout := 20 * time.Second
		conn, err2 := net.DialTimeout("tcp", ipAddress+":"+strconv.Itoa(linkPort-1), connTimeout)

		if err2 != nil {
			err = err2
		} else {
			_, err = conn.Write(byteData)
			if err != nil {

			} else {
				//接收服务端反馈
				var buffer bytes.Buffer

				//				err = conn.SetReadDeadline(time.Now().Add(time.Second))
				err = conn.SetReadDeadline(time.Now().Add(connTimeout))
				if err == nil {
					buf := make([]byte, 8)
					_, err2 := conn.Read(buf)
					if err2 != nil {
						err = err2
					} else {
						length := int(binary.LittleEndian.Uint32(buf[0:4])) //数据正文的长度
						io.CopyN(&buffer, conn, int64(length))

						responseMeta := &Adoter_Asset.Meta{}
						bufferBytes := buffer.Bytes()
						err = proto.Unmarshal(bufferBytes[:], responseMeta)
						responseBytes = responseMeta.Stuff[:]
					}
				}
			}

			conn.Close()
		}
	}

	return responseBytes, err
}

func (this *NetworkUtil) Get(reqUrl string, reqParams map[string]string, headers map[string]string) string {
	var httpClient = &http.Client{}
	urlParams := url.Values{}
	Url, _ := url.Parse(reqUrl)

	for key, value := range reqParams {
		urlParams.Set(key, value)
	}

	//参数种有中文参数，进行urlencode
	Url.RawQuery = urlParams.Encode()
	//完整的urL
	urlPath := Url.String()

	httpRequest, _ := http.NewRequest("GET", urlPath, nil)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}

	//发请求
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	return string(response)

}
