package domain

import "time"

type HealthPayload struct {
	ClientID   string         `json:"client_id"`
	AppVersion string         `json:"app_version"`
	Platform   string         `json:"platform"`
	Status     string         `json:"status"`
	Metadata   map[string]any `json:"metadata,omitempty"`
}

type HealthEvent struct {
	ID         string        `bson:"_id,omitempty" json:"id"`
	ReceivedAt time.Time     `bson:"receivedAt" json:"receivedAt"`
	RemoteIP   string        `bson:"remoteIP" json:"remoteIP"`
	RequestID  string        `bson:"requestId" json:"requestId"`
	UserAgent  string        `bson:"userAgent" json:"userAgent"`
	Payload    HealthPayload `bson:"payload" json:"payload"`
}
