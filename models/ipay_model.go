package models

type GetLoanContractRequest struct {
	Operation            int    `json:"Operation"`
	FinancialCode        string `json:"FinancialCode"`
	ContractNo           string `json:"ContractNo"`
	IdNo                 string `json:"IdNo"`
	AreaCode             string `json:"AreaCode"`
	ReferenceId          string `json:"ReferenceId"`
	BillId               string `json:"BillId"`
	CurrentPaymentAmount string `json:"CurrentPaymentAmount"`
}

type CheckBalanceRequest struct {
	PartnerCode string `json:"partnerCode"`
	Sign        string `json:"sign"`
}

type CheckBalanceResponse struct {
	ResCode        string `json:"resCode"`
	ResMessage     string `json:"resMessage"`
	CurrentBalance int64  `json:"currentBalance"`
	Sign           string `json:"sign"`
	SignData       string `json:"signData,omitempty"`
}

type DirectTopupRequest struct {
	TelcoCode   string `json:"telcoCode"`
	MobileNo    string `json:"mobileNo"`
	TopupAmount int    `json:"topupAmount"`
}

type DirectTopupInput struct {
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	TelcoCode      string `json:"telcoCode"`
	MobileNo       string `json:"mobileNo"`
	TopupAmount    int    `json:"topupAmount"`
	Sign           string `json:"sign"`
}

type DirectTopupResponse struct {
	ResCode        string `json:"resCode"`
	ResMessage     string `json:"resMessage"`
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	DiscountValue  int64  `json:"discountValue"`
	DebitValue     int64  `json:"debitValue"`
	MobileType     string `json:"mobileType,omitempty"`
	Sign           string `json:"sign"`
	SignData       string `json:"signData,omitempty"`
}

type BuyCardRequest struct {
	ProductCode string `json:"productCode"`
	Reciever    string `json:"reciever,omitempty"`
	Quantity    int    `json:"quantity"`
}

type BuyCardInput struct {
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	ProductCode    string `json:"productCode"`
	Reciever       string `json:"reciever,omitempty"`
	Quantity       int    `json:"quantity"`
	Sign           string `json:"sign"`
}

type CardInfo struct {
	Serial     string `json:"serial"`
	Pincode    string `json:"pincode"`
	Expiredate int64  `json:"expiredate"`
}

type BuyCardResponse struct {
	ResCode        string      `json:"resCode"`
	ResMessage     string      `json:"resMessage"`
	PartnerCode    string      `json:"partnerCode"`
	PartnerTransId string      `json:"partnerTransId"`
	TotalValue     int64       `json:"totalValue"`
	DiscountValue  int64       `json:"discountValue"`
	DebitValue     int64       `json:"debitValue"`
	CardList       []*CardInfo `json:"cardList"`
	Sign           string      `json:"sign"`
	SignData       string      `json:"signData,omitempty"`
}

type CheckTransactionRequest struct {
	PartnerTransId string `json:"partnerTransId"`
	TransType      string `json:"transType"`
}

type CheckTransactionInput struct {
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	TransType      string `json:"transType"`
	Sign           string `json:"sign"`
}

type CheckTransactionResponse struct {
	ResCode        string `json:"resCode"`
	ResMessage     string `json:"resMessage"`
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	TransStatus    int8   `json:"transStatus"`
	Sign           string `json:"sign"`
	SignData       string `json:"signData,omitempty"`
}

type RetrieveCardInfoRequest struct {
	PartnerTransId string `json:"partnerTransId"`
}

type RetrieveCardInfoInput struct {
	PartnerCode    string `json:"partnerCode"`
	PartnerTransId string `json:"partnerTransId"`
	Sign           string `json:"sign"`
}

type RetrieveCardInfoResponse struct {
	ResCode        string      `json:"resCode"`
	ResMessage     string      `json:"resMessage"`
	PartnerCode    string      `json:"partnerCode"`
	PartnerTransId string      `json:"partnerTransId"`
	TotalValue     int64       `json:"totalValue"`
	DiscountValue  int64       `json:"discountValue"`
	DebitValue     int64       `json:"debitValue"`
	CardList       []*CardInfo `json:"cardList"`
	Sign           string      `json:"sign"`
	SignData       string      `json:"signData,omitempty"`
}

type CheckProductInfoRequest struct {
	ProductCode string `json:"productCode"`
}

type CheckProductInfoInput struct {
	PartnerCode string `json:"partnerCode"`
	ProductCode string `json:"productCode"`
	Sign        string `json:"sign"`
}

type CheckProductInfoResponse struct {
	ResCode    string `json:"resCode"`
	ResMessage string `json:"resMessage"`
	Quantity   int    `json:"quantity"`
	Sign       string `json:"sign"`
	SignData   string `json:"signData,omitempty"`
}
