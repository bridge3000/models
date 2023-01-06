package models

import ()

type LoginLog struct {
	DateTime     string `json:"datetime"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	ClientIP     string `json:"client_ip"`
	Proto        string `json:"proto"`
	StatusCode   string `json:"status_code"`
	Latency      string `json:"latency"`
	UserAgent    string `json:"user_agent"`
	ErrorMessage string `json:"error_message"`
}
