package server

import (
	"encoding/json"
	"go-reminder-bot/pkg/enum"
	"go-reminder-bot/pkg/user"
	"go-reminder-bot/pkg/util"
	"time"
)

type GetListRequest struct {
	Filter map[string]interface{} `form:"filter"`
	Range  string                 `form:"range"`
	Sort   string                 `form:"sort"`
}

func (r GetListRequest) toGetListParams() util.GetListParams {
	p := util.GetListParams{
		Filter:   r.Filter,
		Limit:    10,
		Offset:   0,
		SortBy:   "id",
		SortType: "ASC",
	}
	if r.Range != "" {
		var queryRange []int
		_ = json.Unmarshal([]byte(r.Range), &queryRange)
		if len(queryRange) == 2 {
			p.Offset, p.Limit = queryRange[0], queryRange[1]
		}
	}
	if r.Sort != "" {
		var querySort []string
		_ = json.Unmarshal([]byte(r.Sort), &querySort)
		if len(querySort) == 2 {
			p.SortBy, p.SortType = querySort[0], querySort[1]
		}
	}
	return p
}

type Reminder struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	IsActive      bool   `json:"is_active"`
	Type          string `json:"type"`
	Schedule      string `json:"schedule"`
	ScheduleHuman string `json:"schedule_human"`
	NextTime      string `json:"next_time"`
	Message       string `json:"message"`
	Webhook       string `json:"webhook"`
	WebhookType   string `json:"webhook_type"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     int64  `json:"updated_at"`
}

type User struct {
	ID        int64         `json:"id"`
	Email     string        `json:"email"`
	IsActive  bool          `json:"is_active"`
	Role      enum.UserRole `json:"role"`
	RoleText  string        `json:"role_text"`
	CreatedAt time.Time     `json:"created_at"`
	Token     string        `json:"token"`
}

func transformUsersFromUsersDB(users []user.User) []User {
	result := make([]User, 0, len(users))
	for _, user := range users {
		result = append(result, transformUserFromUserDB(user))
	}
	return result
}

func transformUserFromUserDB(t user.User) User {
	return User{
		ID:        int64(t.Model.ID),
		Email:     t.Email,
		Role:      t.Role,
		RoleText:  t.Role.String(),
		IsActive:  t.IsActive,
		CreatedAt: t.CreatedAt,
	}
}
