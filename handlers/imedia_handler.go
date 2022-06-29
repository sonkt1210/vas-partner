package handlers

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/sonkt1210/tiki-vas/services/gosoap"
	"log"
	"net/http"
	"time"
)

// RequestHandle a simple request
type RequestHandle struct {
	XMLName xml.Name `xml:"requestHandle"`
	XmlNS   string   `xml:"xmlns,attr"`

	RequestData string `xml:"requestData" json:"requestData"`
}

//// RequestData a simple request
//type RequestData struct {
//	XMLName xml.Name `xml:"requestData"`
//	Type    string   `xml:"type,attr"`
//}

// LoginRequest a simple request
type LoginRequest struct {
	Operation    int         `xml:"operation"`
	Username     string      `xml:"username"`
	MerchantPass string      `xml:"merchantPass"`
	Signature    interface{} `xml:"signature"`
}

type Params map[string]interface{}

// FooResponse a simple response
type FooResponse struct {
	Bar string
}

type RequestHandleData struct {
	Operation       int                      `json:"operation"`
	Username        string                   `json:"username"`
	RequestID       string                   `json:"requestID"`
	MerchantPass    string                   `json:"merchantPass"`
	BuyItems        []map[string]interface{} `json:"buyItems"`
	KeyBirthdayTime string                   `json:"keyBirthdayTime"`
	TargetAccount   string                   `json:"targetAccount"`
	ProviderCode    string                   `json:"providerCode"`
	TopupAmount     int                      `json:"topupAmount"`
	AccountType     string                   `json:"accountType"`
	Signature       string                   `json:"signature"`
	Token           string                   `json:"token"`
}

type RequestHandleResponse struct {
	RequestHandleReturn string `xml:"requestHandleReturn"`
}

type RequestHandleReturnData struct {
	ErrorCode       int                      `json:"errorCode"`
	ErrorMessage    string                   `json:"errorMessage"`
	Products        []map[string]interface{} `json:"products"`
	RequestID       string                   `json:"requestID"`
	MerchantBalance int                      `json:"merchantBalance"`
	SysTransId      int                      `json:"sysTransId"`
	Signature       string                   `json:"signature"`
	Token           string                   `json:"token"`
}

// QueryBalanceResponse ...
type QueryBalanceResponse struct {
	QueryBalanceReturnData QueryBalanceReturnData `xml:"queryBalanceReturn"`
}

// QueryBalanceReturnData ...
type QueryBalanceReturnData struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	DataValue    int    `json:"dataValue"`
	TransId      int    `json:"transID"`
}

// QueryBalanceInput ...
type QueryBalanceInput struct {
	QueryBalance QueryBalance `xml:"queryBalance"`
}

// QueryBalance ...
type QueryBalance struct {
	Username string `xml:"username"`
	Token    string `xml:"token"`
}

var (
	r RequestHandleResponse
)

func Login(c *gin.Context) {
	//rqData := fmt.Sprintf("IMEDIA_DEV20" + "|" + "62103363")
	//sign, err := helpers.IMediaRSAEncrypt([]byte(rqData))
	//if err != nil {
	//	log.Fatalf("IMedia encrypt error: %s", err)
	//}
	//stringSign := b64.StdEncoding.EncodeToString(sign)
	//fmt.Println(stringSign)
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	gosoap.SetCustomEnvelope("soapenv", map[string]string{
		"xmlns:soapenv": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:tem":     "http://tempuri.org/",
	})

	soap, err := gosoap.SoapClient("http://103.68.241.77:8080/ItopupService1.4_IMD/services/TopupInterface?wsdl", httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	//requestHanderData := &RequestHandleData{}
	//requestHanderData.Operation = 1400
	//requestHanderData.Username = "IMEDIA_DEV20"
	//requestHanderData.MerchantPass = "62103363"
	//requestHanderData.Signature = stringSign
	//requestDataJson, _ := json.Marshal(requestHanderData)
	//params := gosoap.Params{
	//	"requestData": string(requestDataJson),
	//}
	//

	params := gosoap.Params{
		"username": "IMEDIA_DEV20",
		"token":    "edee4624f1781c9f3982fe9124a584b88649f72e7115cbf1",
	}

	//params := QueryBalanceInput{
	//	QueryBalance: QueryBalance{
	//		Username: "IMEDIA_DEV20",
	//		Token:    "edee4624f1781c9f3982fe9124a584b88649f72e7115cbf1",
	//	},
	//}

	//req := gosoap.NewRequest("queryBalance", params)
	//gosoap.NewRequestByStruct()
	//stc, _ := gosoap.NewRequestByStruct(req)

	res, err := soap.Call("queryBalance", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}
	result := QueryBalanceResponse{}

	res.Unmarshal(&result)
	// RequestHandleReturn will be a string. We need to parse it to json
	//err = json.Unmarshal([]byte(r.QueryBalanceReturnData), &result)
	//if err != nil {
	//	log.Fatalf("json.Unmarshal error: %s", err)
	//}
	//
	//signData := fmt.Sprintf("%s|%s|%s|%s", strconv.Itoa(result.ErrorCode), "\"\"", strconv.Itoa(result.SysTransId), result.Token)
	//fmt.Println(signData)
	//// verify io response
	//err = helpers.IMediaRSAVerify(signData, result.Signature)
	////if err != nil {
	////	log.Fatalf("Verify data from IMedia error: %v", err)
	////}

	c.JSON(200, result)
}
