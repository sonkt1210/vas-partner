package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"strconv"

	"github.com/sonkt1210/tiki-vas/helpers"
	"github.com/sonkt1210/tiki-vas/logger"
	"github.com/sonkt1210/tiki-vas/models"
	"github.com/sonkt1210/tiki-vas/repositories"
	"github.com/sonkt1210/tiki-vas/services"
)

func CheckBalance(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========CheckBalance==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	partnerCode := os.Getenv("PARTNER_CODE")
	rqData := fmt.Sprintf(partnerCode)
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)

	rqParams := models.CheckBalanceRequest{
		PartnerCode: partnerCode,
		Sign:        stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.CheckBalance(rqParams)
	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		response.Data = rs
	}
	c.JSON(200, response)
}

func DirectTopup(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========DirectTopup==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	request := &models.DirectTopupRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		logger.Warnf("Invalid input. err: %+v", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	err = godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
	}

	partnerCode := os.Getenv("PARTNER_CODE")

	letterRunes := []rune("0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	partnerTransId := partnerCode + "_" + string(b)

	rqData := fmt.Sprintf(partnerCode + partnerTransId + request.TelcoCode + request.MobileNo + strconv.Itoa(request.TopupAmount))
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)

	rqParams := models.DirectTopupInput{
		PartnerCode:    partnerCode,
		PartnerTransId: partnerTransId,
		TelcoCode:      request.TelcoCode,
		MobileNo:       request.MobileNo,
		TopupAmount:    request.TopupAmount,
		Sign:           stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.DirectTopup(rqParams)

	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		response.Data = rs
	}
	c.JSON(200, response)
}

func BuyCard(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========BuyCard==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	request := &models.BuyCardRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		logger.Warnf("Invalid input. err: %+v", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	err = godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
	}

	partnerCode := os.Getenv("PARTNER_CODE")

	letterRunes := []rune("0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	partnerTransId := partnerCode + "_" + string(b)

	rqData := fmt.Sprintf(partnerCode + partnerTransId + request.ProductCode + strconv.Itoa(request.Quantity))
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)
	rqParams := models.BuyCardInput{
		PartnerCode:    partnerCode,
		PartnerTransId: partnerTransId,
		ProductCode:    request.ProductCode,
		Quantity:       request.Quantity,
		Sign:           stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.BuyCard(rqParams)
	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		for i := 0; i < len(rs.CardList); i++ {
			decrypt, err := helpers.RSADecrypt(rs.CardList[i].Pincode)
			if err != nil {
				response.Status = 400
				response.Message = "Err"
				response.Error = err.Error()
				break
			} else {
				rs.CardList[i].Pincode = decrypt
			}
		}
		response.Data = rs
	}
	c.JSON(200, response)
}

func CheckTransaction(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========CheckTransaction==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	request := &models.CheckTransactionRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		logger.Warnf("Invalid input. err: %+v", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	err = godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
	}

	partnerCode := os.Getenv("PARTNER_CODE")

	rqData := fmt.Sprintf(partnerCode + request.PartnerTransId + request.TransType)
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)

	rqParams := models.CheckTransactionInput{
		PartnerCode:    partnerCode,
		PartnerTransId: request.PartnerTransId,
		TransType:      request.TransType,
		Sign:           stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.CheckTransaction(rqParams)

	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		response.Data = rs
	}
	c.JSON(200, response)
}

func RetrieveCardInfo(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========RetrieveCardInfo==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	request := &models.RetrieveCardInfoRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		logger.Warnf("Invalid input. err: %+v", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	err = godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
	}

	partnerCode := os.Getenv("PARTNER_CODE")

	rqData := fmt.Sprintf(partnerCode + request.PartnerTransId)
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)

	rqParams := models.RetrieveCardInfoInput{
		PartnerCode:    partnerCode,
		PartnerTransId: request.PartnerTransId,
		Sign:           stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.RetrieveCardInfo(rqParams)

	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		for i := 0; i < len(rs.CardList); i++ {
			decrypt, err := helpers.RSADecrypt(rs.CardList[i].Pincode)
			if err != nil {
				response.Status = 400
				response.Message = "Err"
				response.Error = err.Error()
				break
			} else {
				rs.CardList[i].Pincode = decrypt
			}
		}
		response.Data = rs
	}
	c.JSON(200, response)
}

func CheckProductInfo(c *gin.Context) {
	cflogger := logger.NewAppLogService(nil)
	logger := cflogger.GetLogger("==========CheckProductInfo==========\n")

	rest := services.NewREST(logger)

	response := models.BaseResponse{
		Status:  200,
		Message: "Success",
		Data:    nil,
		Error:   "",
	}

	request := &models.CheckProductInfoRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		logger.Warnf("Invalid input. err: %+v", err)
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	}

	err = godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Load env fall Err: %s", err)
	}

	partnerCode := os.Getenv("PARTNER_CODE")

	rqData := fmt.Sprintf(partnerCode + request.ProductCode)
	sign, err := helpers.RSAEncrypt([]byte(rqData))
	stringSign := b64.StdEncoding.EncodeToString(sign)

	rqParams := models.CheckProductInfoInput{
		PartnerCode: partnerCode,
		ProductCode: request.ProductCode,
		Sign:        stringSign,
	}

	ipayRepo := repositories.NewIpayRepository(logger, rest)

	rs, err := ipayRepo.CheckProductInfo(rqParams)

	if err != nil {
		response.Status = 400
		response.Message = "Err"
		response.Error = err.Error()
	} else {
		err = helpers.RSAVerify(rs.SignData, rs.Sign)
		if err != nil {
			response.Status = 400
			response.Message = "Err"
			response.Error = err.Error()
		}
		response.Data = rs
	}
	c.JSON(200, response)
}
