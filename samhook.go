package samhook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// AddAttachment 添加一個attachment
func (m *Message) AddAttachment(attachment Attachment) *Message {
	m.Attachments = append(m.Attachments, attachment)
	return m
}

// AddAttachments 添加多個attachment
func (m *Message) AddAttachments(attachments []Attachment) *Message {
	m.Attachments = append(m.Attachments, attachments...)
	return m
}

// Send 發送message
func Send(url string, msg Message) error {
	payloadBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

// SendReader 發送message
func SendReader(url string, r io.Reader) error {

	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
