package delivery

import (
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/delivery/appresponse"
	"github.com/edwardsuwirya/wmbPos/dto"
	"github.com/edwardsuwirya/wmbPos/usecase"
	"github.com/gin-gonic/gin"
)

type OrderApi struct {
	usecase     usecase.IOrderUseCase
	publicRoute *gin.RouterGroup
}

func NewOrderApi(publicRoute *gin.RouterGroup, usecase usecase.IOrderUseCase) *OrderApi {
	orderApi := OrderApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	orderApi.initRouter()
	return &orderApi
}
func (api *OrderApi) initRouter() {
	userRoute := api.publicRoute.Group("/order")
	userRoute.POST("/open", api.openOrder)
	userRoute.POST("/close", api.closeOrder)
}
func (api *OrderApi) openOrder(c *gin.Context) {
	var orderRequest dto.CustomerOrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewBadRequestError(err, "Failed Order"))
		return
	}
	_, err := api.usecase.OpenOrder(orderRequest)
	if err != nil {
		if err == apperror.TableOccupiedError {
			appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("FAILED", "Table Is Occupied", nil))
			return
		}
		if err == apperror.ClientTimeOut {
			appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("FAILED", "Table Reservation Time out", nil))
			return
		}
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Order"))
		return
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Success Order", nil))
}
func (api *OrderApi) closeOrder(c *gin.Context) {
	var closeOrderRequest dto.CloseOrderRequest
	if err := c.ShouldBindJSON(&closeOrderRequest); err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewBadRequestError(err, "Failed Close Order"))
		return
	}
	_, err := api.usecase.CloseOrder(closeOrderRequest)
	if err != nil {
		if err == apperror.ClientTimeOut {
			appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("FAILED", "Table Reservation Time out", nil))
			return
		}
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Close Order"))
		return
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Success Close Order", nil))
}
