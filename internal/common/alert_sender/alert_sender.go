package alert_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const contentType = "application/json"

type AlertSender struct {
	webhookURL string
	httpClient *http.Client
}

func NewAlertSender(cfg Config) *AlertSender {
	return &AlertSender{
		webhookURL: cfg.WebhookURL,
		httpClient: &http.Client{},
	}
}
func (s *AlertSender) SendAlert(message string) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	resp, err := s.httpClient.Post(s.webhookURL, contentType, bytes.NewBuffer(messageJSON))
	if err != nil {
		return fmt.Errorf("failed to send alert")
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got unexpected response from webhook, %d: %v", resp.StatusCode, err)
	}

	return nil
}
