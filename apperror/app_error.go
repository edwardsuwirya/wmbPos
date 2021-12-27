package apperror

import "errors"

var TableOccupiedError = errors.New("Table is occupied")
var ClientTimeOut = errors.New("Time Out")
