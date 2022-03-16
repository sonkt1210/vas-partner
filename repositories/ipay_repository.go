package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sonkt1210/tiki-vas/logger"
	"github.com/sonkt1210/tiki-vas/models"
	"github.com/sonkt1210/tiki-vas/services"
	"github.com/sonkt1210/tiki-vas/utils"
	"os"
)

type IpayRepository interface {
	CheckBalance(args models.CheckBalanceRequest) (models.CheckBalanceResponse, error)
	DirectTopup(args models.DirectTopupInput) (models.DirectTopupResponse, error)
	BuyCard(args models.BuyCardInput) (models.BuyCardResponse, error)
	CheckTransaction(args models.CheckTransactionInput) (models.CheckTransactionResponse, error)
	RetrieveCardInfo(args models.RetrieveCardInfoInput) (models.RetrieveCardInfoResponse, error)
	CheckProductInfo(args models.CheckProductInfoInput) (models.CheckProductInfoResponse, error)
}

type ipayRepository struct {
	logger  logger.Logger
	restApi services.REST
}

func NewIpayRepository(logger logger.Logger, restApi services.REST) IpayRepository {
	return &ipayRepository{
		logger:  logger,
		restApi: restApi,
	}
}

func (r *ipayRepository) CheckBalance(args models.CheckBalanceRequest) (models.CheckBalanceResponse, error) {
	var response []byte
	var err error
	var rs models.CheckBalanceResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "checkBalance")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}

func (r *ipayRepository) DirectTopup(args models.DirectTopupInput) (models.DirectTopupResponse, error) {
	var response []byte
	var err error
	var rs models.DirectTopupResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "directTopup")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}

func (r *ipayRepository) BuyCard(args models.BuyCardInput) (models.BuyCardResponse, error) {
	var response []byte
	var err error
	var rs models.BuyCardResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "buyCard")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}

func (r *ipayRepository) CheckTransaction(args models.CheckTransactionInput) (models.CheckTransactionResponse, error) {
	var response []byte
	var err error
	var rs models.CheckTransactionResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "checkTransaction")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}

func (r *ipayRepository) RetrieveCardInfo(args models.RetrieveCardInfoInput) (models.RetrieveCardInfoResponse, error) {
	var response []byte
	var err error
	var rs models.RetrieveCardInfoResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "retrieveCardInfo")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}

func (r *ipayRepository) CheckProductInfo(args models.CheckProductInfoInput) (models.CheckProductInfoResponse, error) {
	var response []byte
	var err error
	var rs models.CheckProductInfoResponse

	err = godotenv.Load(".env")
	if err != nil {
		r.logger.Errorf("Load env error Err: %s", err)
		return rs, err
	}

	apiEndpoint := fmt.Sprintf(os.Getenv("IPAY_URL") + "checkProductInfo")
	header := map[string]string{"Content-Type": "application/json"}

	params, _ := utils.Marshal(args)
	response, err = r.restApi.Post(apiEndpoint, params, header)
	if err != nil {
		r.logger.Errorf("Failed request to History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}

	err = utils.Unmarshal(response, &rs)
	if err != nil {
		r.logger.Errorf("Failed Unmarshal data History, \nrequest = %s, \nresponse = %s, \nerror = %s", apiEndpoint, string(response), err.Error())
		return rs, err
	}
	r.logger.Infof("data response: %s", string(response))
	return rs, nil
}
