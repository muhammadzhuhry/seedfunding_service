package transaction

import "github.com/muhammadzhuhry/bwastartup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
