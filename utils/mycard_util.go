package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	//	"fmt"
	"net/url"
	"strings"
)

type MycardUtil struct {
	IsDebugMode bool
	Key         string
}

//mycard响应回调
type MycardResponse struct {
	ReturnCode     string
	ReturnMsg      string
	AuthCode       string
	TradeSeq       string
	InGameSaveType string
	TransactionUrl string
	MyCardTradeNo  string
	FacTradeSeqs   []string //3.6.4使用，本来是FacTradeSeq，但跟下面那行冲突，所以字符串替换成复数形式
	FacTradeSeq    string   //3.2.4
	PaymentType    string
	PayResult      string
}

func (this *MycardUtil) makeSign(plainText string) string {
	preHashValue := plainText + this.Key
	hash := strings.ToLower(url.QueryEscape(preHashValue))
	sum := sha256.Sum256([]byte(hash))
	return hex.EncodeToString(sum[:])
}

//网页输入卡号 密码直接充值
func (this *MycardUtil) IngamePay(authCode string, cardId string, cardPw string) (isSuccess bool, errMsg, returnCode string, responseBytes []byte) {
	var err error
	networkUtil := NetworkUtil{}
	mycardResponse := MycardResponse{}
	verifyUrl := "https://b2b.mycard520.com.tw/MyBillingPay/v1.2/IngamePay" //正式

	if this.IsDebugMode {
		verifyUrl = "https://testb2b.mycard520.com.tw/MyBillingPay/v1.2/IngamePay" //測試
	}

	//下单参数，返回授权码
	params := map[string]interface{}{
		"AuthCode": authCode,
		"CardID":   cardId,
		"CardPW":   cardPw,
		"Hash":     this.makeSign(authCode + cardId + cardPw),
	}

	responseBytes, err = networkUtil.HttpPost(verifyUrl, 1, params)
	if err != nil {
		errMsg = err.Error()
	} else {
		json.Unmarshal(responseBytes, &mycardResponse)
		returnCode = mycardResponse.ReturnCode
		if mycardResponse.ReturnCode == "1" && mycardResponse.PayResult == "3" {
			isSuccess = true
		}
	}

	return
}

func (this *MycardUtil) TradeQuery(authCode string) (isSuccess bool, errMsg string, responseBytes []byte, myCardTradeNo string, paymentType string) {
	var err error
	networkUtil := NetworkUtil{}
	mycardResponse := MycardResponse{}
	verifyUrl := "https://b2b.mycard520.com.tw/MyBillingPay/v1.2/TradeQuery" //正式
	if this.IsDebugMode {
		verifyUrl = "http://testb2b.mycard520.com.tw/MyBillingPay/v1.2/TradeQuery" //測試
	}

	_, responseBytes, err, _ = networkUtil.HttpRequest("GET", verifyUrl+"?AuthCode="+authCode, map[string]string{}, nil)
	if err == nil {
		json.Unmarshal(responseBytes, &mycardResponse)

		if mycardResponse.ReturnCode == "1" && mycardResponse.PayResult == "3" {
			isSuccess = true
			myCardTradeNo = mycardResponse.MyCardTradeNo
			paymentType = mycardResponse.PaymentType
		}
	} else {
		errMsg = err.Error()
	}

	return
}

//请款
func (this *MycardUtil) PaymentConfirm(authCode string) (isSuccess bool, errMsg string, responseBytes []byte) {
	var err error
	networkUtil := NetworkUtil{}
	payConfirmResponse := MycardResponse{}

	paymentConfirmUrl := "https://b2b.mycard520.com.tw/MyBillingPay/v1.2/PaymentConfirm"

	if this.IsDebugMode {
		paymentConfirmUrl = "http://testb2b.mycard520.com.tw/MyBillingPay/v1.2/PaymentConfirm"
	}

	_, responseBytes, err, _ = networkUtil.HttpRequest("GET", paymentConfirmUrl+"?AuthCode="+authCode, map[string]string{}, nil)
	if err == nil {
		json.Unmarshal(responseBytes, &payConfirmResponse)
		if payConfirmResponse.ReturnCode == "1" { //请款成功
			isSuccess = true

		}

	} else {
		errMsg = err.Error()
	}

	return
}
