package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type PaypalUtil struct {
}

type PaypalLink struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type PaypalResponse struct {
	Id               string                 `json:"id"`
	Name             string                 `json:"name"`
	Links            []PaypalLink           `json:"links"`
	Status           string                 `json:"status"`
	Error            string                 `json:"error"`
	ErrorDescription string                 `json:"error_description"`
	Scope            string                 `json:"scope"`
	AccessToken      string                 `json:"access_token"`
	TokenType        string                 `json:"token_type"`
	AppId            string                 `json:"app_id"`
	ExpiresIn        string                 `json:"expires_in"`
	Nonce            string                 `json:"nonce"`
	EventType        string                 `json:"event_type"`
	Resource         PaypalResponseResource `json:"resource"`
	PurchaseUnits    []PurchaseUnit         `json:"purchase_units"`
}

type PurchaseUnit struct {
	Payments map[string][]Capture `json:"payments"`
}

type Capture struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type PaypalResponseResource struct {
	Id                string                                  `json:"id"`
	Status            string                                  `json:"status"`
	SupplementaryData PaypalResponseResourceSupplementaryData `json:"supplementary_data"`
}

type PaypalResponseResourceSupplementaryData struct {
	RelatedIds map[string]string `json:"related_ids"`
}

func (this *PaypalUtil) getClientIdAndSecret(isDebugMode bool) (clientId string, secret string) {
	if isDebugMode {
		clientId = "ASmPnm0drajgbDGVeJzn89eFKmesq2b1S4JpG1Ve6wZvBICBsGGPZX8A2uER2U4nRC-O4WI5EQe7F3Nk"
		secret = "EK4aGxhznhKtP1ktSaJPQzpgYYbDMgUzmmrOf-CCDSidC_qDsx4PonGz55sjzbPxZV1waKYietASRuav"
	} else {
		clientId = "AZhmya0ouc71uyYDKndAtsxMHJEk2K-Sr7txlvz1fyd_gtr_Vd53iSuMEO_FIssUzd0YC_I6kGQI7asS"
		secret = "EFgztaNSTtk4jEkY6XQvmiV3FXxYWjGZ2oMpUcakpg4jEK2e23z9gQfvlGPg1nwk5TT6zVOLzl_Ozh0V"
	}
	return
}

func (this *PaypalUtil) getDomain(isDebugMode bool) string {
	if isDebugMode {
		return "https://api-m.sandbox.paypal.com" //	Sandbox.
	} else {
		return "https://api-m.paypal.com" //Live.
	}
}

func (this *PaypalUtil) getAccessToken(isDebugMode bool) (accessToken string, err error) {
	networkUtil := NetworkUtil{}
	var ioReader io.Reader
	responseBytes := []byte{}
	url := "/v1/oauth2/token"
	contentType := "application/x-www-form-urlencoded"
	clientId, clientSecret := this.getClientIdAndSecret(isDebugMode)
	input := []byte(clientId + ":" + clientSecret)
	encodedStr := base64.StdEncoding.EncodeToString(input) // 序列化
	headers := map[string]string{
		"Content-Type":  contentType,
		"Authorization": "Basic " + encodedStr,
	}

	ioReader, err = networkUtil.BuildHttpBody(contentType, map[string]interface{}{"grant_type": "client_credentials"})
	_, responseBytes, err, _ = networkUtil.HttpRequest("POST", this.getDomain(isDebugMode)+url, headers, ioReader)
	paypalResponse := PaypalResponse{}
	json.Unmarshal(responseBytes, &paypalResponse)
	accessToken = paypalResponse.AccessToken
	return
}

func (this *PaypalUtil) CreateOrder(isDebugMode bool, currencyCode string, amount string) (platformOutTradeNo string, transactionUrl string, err error, status string, name string) {
	networkUtil := NetworkUtil{}
	url := "/v2/checkout/orders"
	contentType := "application/json;charset=utf-8"
	var ioReader io.Reader
	responseBytes := []byte{}
	accessToken := ""
	accessToken, err = this.getAccessToken(isDebugMode)
	createOrderResponse := PaypalResponse{}

	headers := map[string]string{
		"Content-Type":  contentType,
		"Authorization": "Bearer " + accessToken,
	}

	purchaseUnits := []map[string]map[string]string{
		map[string]map[string]string{
			"amount": map[string]string{
				"currency_code": currencyCode,
				"value":         amount,
			},
		},
	}

	body := map[string]interface{}{"intent": "CAPTURE", "purchase_units": purchaseUnits}

	ioReader, err = networkUtil.BuildHttpBody(contentType, body)
	_, responseBytes, err, _ = networkUtil.HttpRequest("POST", this.getDomain(isDebugMode)+url, headers, ioReader)
	json.Unmarshal(responseBytes, &createOrderResponse)
	status = createOrderResponse.Resource.Status
	platformOutTradeNo = createOrderResponse.Id
	name = createOrderResponse.Name

	for _, link := range createOrderResponse.Links {
		if link.Rel == "approve" {
			transactionUrl = link.Href
			break
		}
	}

	return
}

//授权
func (this *PaypalUtil) Capture(isDebugMode bool, platformOutTradeNo string) (httpCode int, status string, captureId string, err error) {
	networkUtil := NetworkUtil{}
	url := "/v2/checkout/orders/" + platformOutTradeNo + "/capture"
	contentType := "application/json;charset=utf-8"
	var ioReader io.Reader
	responseBytes := []byte{}

	accessToken := ""
	accessToken, err = this.getAccessToken(isDebugMode)

	headers := map[string]string{
		"Content-Type":  contentType,
		"Authorization": "Bearer " + accessToken,
	}

	ioReader, err = networkUtil.BuildHttpBody(contentType, map[string]interface{}{})
	_, responseBytes, err, httpCode = networkUtil.HttpRequest("POST", this.getDomain(isDebugMode)+url, headers, ioReader)

	payPalResponse := &PaypalResponse{}
	json.Unmarshal(responseBytes, payPalResponse)

	fmt.Println("capture_debug:", string(responseBytes))
	fmt.Println(status, captureId)

	if len(payPalResponse.PurchaseUnits) > 0 {
		if len(payPalResponse.PurchaseUnits[0].Payments["captures"]) > 0 {
			status = payPalResponse.PurchaseUnits[0].Payments["captures"][0].Status
			captureId = payPalResponse.PurchaseUnits[0].Payments["captures"][0].Id
		}
	}

	return
}

func (this *PaypalUtil) Authorize(isDebugMode bool, platformOutTradeNo string) (err error) {
	networkUtil := NetworkUtil{}
	url := "/v2/checkout/orders/" + platformOutTradeNo + "/authorize"
	contentType := "application/json;charset=utf-8"
	var ioReader io.Reader
	accessToken := ""
	accessToken, err = this.getAccessToken(isDebugMode)

	headers := map[string]string{
		"Content-Type":  contentType,
		"Authorization": "Bearer " + accessToken,
	}

	ioReader, err = networkUtil.BuildHttpBody(contentType, map[string]interface{}{})
	_, _, err, _ = networkUtil.HttpRequest("POST", this.getDomain(isDebugMode)+url, headers, ioReader)
	return
}

func (this *PaypalUtil) QueryOrder(isDebugMode bool, platformOutTradeNo string) (status string, captureId string, err error) {
	networkUtil := NetworkUtil{}
	url := "/v2/checkout/orders/" + platformOutTradeNo
	contentType := "application/json;charset=utf-8"
	responseBytes := []byte{}
	accessToken := ""
	accessToken, err = this.getAccessToken(isDebugMode)
	paypalResponse := PaypalResponse{}
	headers := map[string]string{
		"Content-Type":  contentType,
		"Authorization": "Bearer " + accessToken,
	}

	_, responseBytes, err, _ = networkUtil.HttpRequest("GET", this.getDomain(isDebugMode)+url, headers, nil)
	json.Unmarshal(responseBytes, &paypalResponse)

	status = paypalResponse.Status

	// if status == "APPROVED" {
	// 	this.Capture(isDebugMode, platformOutTradeNo)
	// }

	if status == "COMPLETED" {
		captureId = paypalResponse.PurchaseUnits[0].Payments["captures"][0].Id
	} else {
		fmt.Println("query_response:", string(responseBytes))
	}

	return
}
