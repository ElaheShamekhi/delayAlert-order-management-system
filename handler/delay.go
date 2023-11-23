package handler

import (
	"delayAlert-order-management-system/server"
	"delayAlert-order-management-system/service/delay"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupDelaysRoutes(s *server.Server, h DelayHandler) {
	s.Engine.POST("/delays", h.CreateOrderDelay)
	s.Engine.PUT("/agents/:agentId/delays", h.AssignDelayToAgent)
	s.Engine.GET("/vendors/delays")

}

type DelayHandler struct {
	delay *delay.Service
}

func NewDelayHandler(delay *delay.Service) DelayHandler {
	return DelayHandler{
		delay: delay,
	}
}

// CreateOrderDelay godoc
// @Summary      Create a Delay For Order
// @Description  Create a Delay For Order.
// @Tags         Delay
// @Accept       json
// @Produce      json
// @Param        body			body		delay.CreateOrderDelay		true	"create delay request"
// @Success      200			{object}	//TODO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /delays	[POST]
func (h DelayHandler) CreateOrderDelay(ctx *gin.Context) {
	var request delay.CreateOrderDelay
	if err := ctx.ShouldBind(&request); err != nil {
		handleError(ctx, err)
		return
	}
	//TODO
}

// AssignDelayToAgent
// @Summary      Assign Delay To Agent
// @Description  Assign Delay to Agent fo checking
// @Tags         Delay
// @Produce      json
// @Param        agentId     path       string  true  "agentId id"
// @Success      200  {object} //TODO
// @Router       /agents/{agentId}/delays [PUT]
func (h DelayHandler) AssignDelayToAgent(ctx *gin.Context) {
	_, err := strconv.Atoi(ctx.Param("agentId"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

}

// VendorsDelayWeeklyReport
// @Summary      Vendors Delay Weekly Report
// @Description  Weekly report of total delay of each vendor
// @Tags         Delay
// @Produce      json
// @Success      200  {object} //TODO
// @Router       /vendors/delays [GET]
func (h DelayHandler) VendorsDelayWeeklyReport(ctx *gin.Context) {

}
