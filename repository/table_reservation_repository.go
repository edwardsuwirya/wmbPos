package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/delivery/appresponse"
	"github.com/edwardsuwirya/wmbPos/dto"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var baseURL = "http://localhost:8888/api/table"

type ITableOrderReservationRepository interface {
	ReserveOne(table dto.TableRequest) error
	Close(billNo string) error
}

type TableOrderReservationRepository struct {
	httpClient *http.Client
}

func NewTableOrderReservation() ITableOrderReservationRepository {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	return &TableOrderReservationRepository{
		httpClient: netClient,
	}
}
func (t *TableOrderReservationRepository) Close(billNo string) error {
	req, err := http.NewRequest(http.MethodPut, baseURL+"/checkout?billNo="+billNo, nil)
	if err != nil {
		return err
	}
	resp, err := t.httpClient.Do(req)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return apperror.ClientTimeOut
		}
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	return errors.New("Failed to check out table")
}
func (t *TableOrderReservationRepository) ReserveOne(table dto.TableRequest) error {
	postBody, _ := json.Marshal(table)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := t.httpClient.Post(baseURL+"/checkin", "application/json", responseBody)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return apperror.ClientTimeOut
		}
		return errors.New("Can not connect to Table Reservation API")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var responseMessage appresponse.ResponseMessage
	err = json.Unmarshal(body, &responseMessage)
	if err != nil {
		return err
	}
	if responseMessage.Status == "FAILED" {
		return apperror.TableOccupiedError
	}
	return nil
}
