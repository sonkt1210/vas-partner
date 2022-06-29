package models

type LoginRequest struct {
	Username     string `json:"username"`
	MerchantPass string `json:"merchantPass"`
	//BuyItems        []BuyItem `json:"buyItems"`
	//RequestID       string    `json:"requestID"`
	//KeyBirthdayTime string    `json:"keyBirthdayTime"`
	//TargetAccount   string    `json:"targetAccount"`
	//ProviderCode    string    `json:"providerCode"`
	//TopupAmount     int       `json:"topupAmount"`
	//AccountType     string    `json:"accountType"`
}

type BuyItem struct {
	ProductID int    `json:"operation"`
	Quantity  string `json:"quantity"`
}

type RequestHandleReturn struct {
	//XMLName      xml.Name  `xml:"GeoIP"`
	ErrorCode       int       `xml:"errorCode"`
	ErrorMessage    string    `xml:"errorMessage"`
	Products        []Product `xml:"products"`
	MerchantBalance float64   `xml:"merchantBalance"`
	RequestId       string    `xml:"requestId"`
	SysTransId      int       `xml:"sysTransId"`
	Signature       string    `xml:"signature"`
	Token           string    `xml:"token"`
}

type Product struct {
	CategoryName        string    `xml:"categoryName"`
	ProductId           string    `xml:"productId"`
	ProductValue        string    `xml:"productValue"`
	ServiceProviderName string    `xml:"serviceProviderName"`
	Softpins            []Softpin `xml:"softpins"`
}

type Softpin struct {
	ExpiryDate     string `xml:"expiryDate"`
	SoftpinId      string `xml:"softpinId"`
	SoftpinPinCode string `xml:"softpinPinCode"`
	SoftpinSerial  string `xml:"softpinSerial"`
}

// RequestHandleResponse will hold the Soap response
type RequestHandleResponse struct {
	RequestHandleReturn string `xml:"requestHandleReturn"`
}
