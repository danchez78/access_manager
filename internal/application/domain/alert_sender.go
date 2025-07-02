package domain

import (
	"fmt"

	"access_manager/internal/common/alert_sender"
)

type AlertSender struct {
	as *alert_sender.AlertSender
}

func NewAlertSender(as *alert_sender.AlertSender) *AlertSender {
	return &AlertSender{as: as}
}

func (s *AlertSender) SendAlert(userID, previuosIPAddress, newIPAddress string) error {
	message := fmt.Sprintf(
		"User with ID: `%s` refreshed tokens with new IP address: `%s`, previous IP address: `%s`",
		userID, previuosIPAddress, newIPAddress,
	)

	return s.as.SendAlert(message)
}
