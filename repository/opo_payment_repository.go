package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/config"
	"github.com/edwardsuwirya/wmbPos/dto"
	"io/ioutil"
	"net"
	"net/http"
)

type IOpoPaymentRepository interface {
	Payment(phoneNo string, total int) (string, error)
}

type OpoPaymentRepository struct {
	httpClient       *http.Client
	opoPaymentConfig config.OpoPaymentConfig
}

func NewOpoPaymentRepository(httpClient *http.Client, config config.OpoPaymentConfig) IOpoPaymentRepository {
	return &OpoPaymentRepository{
		httpClient:       httpClient,
		opoPaymentConfig: config,
	}
}

func (o *OpoPaymentRepository) Payment(phoneNo string, total int) (string, error) {
	postBody, _ := json.Marshal(dto.PaymentRequest{
		CustomerPhoneNo: phoneNo,
		Total:           total,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, o.opoPaymentConfig.ApiBaseUrl, responseBody)
	req.Header.Set("Opo-Client-Key", o.opoPaymentConfig.ClientSecretKey)
	resp, err := o.httpClient.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "", apperror.ClientTimeOut
		}
		return "", errors.New("Can not connect to OPO Payment API")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Failed OPO Payment API")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))
	var opoResp dto.OpoHttpResponse
	err = json.Unmarshal(body, &opoResp)
	if err != nil {
		return "", err
	}
	fmt.Println(opoResp.Data.ReceiptId)
	return opoResp.Data.ReceiptId, nil
}
