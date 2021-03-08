package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadzhuhry/bwastartup/helper"
	"github.com/muhammadzhuhry/bwastartup/transaction"
	"github.com/muhammadzhuhry/bwastartup/user"
	"net/http"
)

// parameter di uri
// tangkap parameter dan mapping ke input struct
// panggil service, passing input struct
// service nerima campaignId, lalu memanggil functin ke repo
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// mengambil key currentUser yang mana currentUser ini di set melalau auth middleware dan membinding ke struct User
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	// mengambil key currentUser yang mana currentUser ini di set melalau auth middleware dan membinding ke struct User
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.transactionService.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// input dari user
// handler tangkap input lalu mapping ke input struct dan manggil service
// manggil service buat transaksi, manggil sistem midtrans
// manggil repository create new transaction data
