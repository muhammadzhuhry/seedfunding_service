package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadzhuhry/bwastartup/campaign"
	"github.com/muhammadzhuhry/bwastartup/helper"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	// handler : mapping id di url ke struct input -> service, memanggil formatter
	// service : inputnya struct input -> menangkap id di url, manggil repository
	// repository : get campaign by id
}
