package handler

import (
	"api_gateway/genproto/transportService"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBusSchedule(ctx *gin.Context) {
	number := ctx.Param("number")
	num, err := strconv.Atoi(number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := &transportService.Transport{Number: int32(num)}

	resp, err := h.Transport.GetBusSchedule(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackBusLocation(ctx *gin.Context) {
	number := ctx.Param("number")
	num, err := strconv.Atoi(number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := &transportService.Transport{Number: int32(num)}

	resp, err := h.Transport.TrackBusLocation(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportTrafficJam(ctx *gin.Context) {
	number := ctx.Param("number")
	num, err := strconv.Atoi(number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req := &transportService.Transport{Number: int32(num)}

	resp, err := h.Transport.ReportTrafficJam(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
