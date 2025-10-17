package service

type MeetingCreateRequest struct {
	Name     string `json:"name,omitempty"`
	CreateAt int64  `json:"create_at"`
	DeleteAt int64  `json:"delete_at"`
}
