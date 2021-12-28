package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/config"
	"github.com/edwardsuwirya/wmbPos/delivery/appresponse"
	"github.com/edwardsuwirya/wmbPos/dto"
	"io/ioutil"
	"net"
	"net/http"
)

type ITableOrderReservationRepository interface {
	CallTableCheckIn(table dto.TableRequest) error
	CallTableCheckOut(billNo string) error
}

type TableOrderReservationRepository struct {
	httpClient            *http.Client
	tableManagementConfig config.TableManagementConfig
}

func NewTableOrderReservation(httpClient *http.Client, config config.TableManagementConfig) ITableOrderReservationRepository {
	return &TableOrderReservationRepository{
		httpClient:            httpClient,
		tableManagementConfig: config,
	}
}
func (t *TableOrderReservationRepository) CallTableCheckOut(billNo string) error {
	req, err := http.NewRequest(http.MethodPut, t.tableManagementConfig.ApiBaseUrl+"/checkout?billNo="+billNo, nil)
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
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("Failed to check out table")
}
func (t *TableOrderReservationRepository) CallTableCheckIn(table dto.TableRequest) error {
	postBody, _ := json.Marshal(table)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := t.httpClient.Post(t.tableManagementConfig.ApiBaseUrl+"/checkin", "application/json", responseBody)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return apperror.ClientTimeOut
		}
		return errors.New("Can not connect to Table Reservation API")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed Table Reservation API")
	}

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
