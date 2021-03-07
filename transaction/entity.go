package transaction

import (
	"github.com/muhammadzhuhry/bwastartup/campaign"
	"github.com/muhammadzhuhry/bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
