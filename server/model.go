package server

import (
	"encoding/json"
	"go-admin/pkg/enum"
	"go-admin/pkg/user"
	"time"
)

type GetListRequest struct {
	Filter map[string]interface{} `form:"filter"`
	Range  string                 `form:"range"`
	Sort   string                 `form:"sort"`
}

type CreateUserRequest struct {
	Username string        `json:"username"`
	IsActive bool          `json:"is_active"`
	Role     enum.UserRole `json:"role"`
}

type UpdateUserRequest struct {
	Username string        `json:"username"`
	IsActive bool          `json:"is_active"`
	Role     enum.UserRole `json:"role"`
}

var acceptFilters = map[string]struct{}{
	"username": {},
	"role":     {},
}

func (r GetListRequest) toGetListParams() user.GetListParams {
	acceptFiltersParam := make(map[string]interface{})
	for k, v := range r.Filter {
		if _, ok := acceptFilters[k]; ok {
			acceptFiltersParam[k] = v
		}
	}
	p := user.GetListParams{
		Filter:   acceptFiltersParam,
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
	ID          uint          `json:"id"`
	Username    string        `json:"username"`
	IsActive    bool          `json:"is_active"`
	Role        enum.UserRole `json:"role"`
	CreatedBy   string        `json:"created_by"`
	CreatedAt   time.Time     `json:"created_at"`
	CurrentUser bool          `json:"current_user"`
}

func transformUsersFromUsersDB(users []user.User, currentUserID uint) []User {
	result := make([]User, 0, len(users))
	for _, user := range users {
		result = append(result, transformUserFromUserDB(user, currentUserID))
	}
	return result
}

func transformUserFromUserDB(t user.User, currentUserID ...uint) User {
	createdBy := t.CreatedBy
	if createdBy == "" {
		createdBy = "System"
	}
	currentUser := false
	if len(currentUserID) != 0 {
		currentUser = currentUserID[0] == t.ID
	}
	return User{
		ID:          t.ID,
		Username:    t.Username,
		Role:        t.Role,
		IsActive:    t.IsActive,
		CreatedAt:   t.CreatedAt,
		CreatedBy:   createdBy,
		CurrentUser: currentUser,
	}
}
