package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/pkg/token"
	"go-admin/pkg/user"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	userStorage user.Storage
	tokenizer   token.Tokenizer
}

func NewHandler(tokenizer token.Tokenizer, userStorage user.Storage) Handler {
	return Handler{userStorage: userStorage, tokenizer: tokenizer}
}
func (h Handler) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()
	var req struct {
		Username string `json:"username" binding:"required"`
	}
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	u, err := h.userStorage.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			u, err = h.registerUser(c, req.Username)
			if err != nil {
				_ = c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if !u.IsActive {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "account is inactive",
		})
		return
	}
	claim := token.Claim{
		UserID:   u.ID,
		IssuedAt: time.Now().Unix(),
	}
	tokenStr, err := h.tokenizer.Generate(claim)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	uRsp := transformUserFromUserDB(u)
	c.JSON(http.StatusOK, gin.H{
		"user":  uRsp,
		"token": tokenStr,
	})
}

func (h Handler) registerUser(ctx context.Context, username string) (u user.User, err error) {
	newUser := user.NewUser(username)
	u, err = h.userStorage.Create(ctx, newUser)
	if err != nil {
		return
	}
	return
}

func (h Handler) findUsers(c *gin.Context) {
	var req GetListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	p := req.toGetListParams()
	list, count, err := h.userStorage.GetList(c, p)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u, _ := ExtractUserFromCtx(c)
	c.Header("Content-Range", fmt.Sprintf("%d-%d/%d", p.Offset, p.Limit, count))
	c.JSON(http.StatusOK, transformUsersFromUsersDB(list, u.ID))
}

func (h Handler) getOneUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	user, err := h.userStorage.GetOne(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u, _ := ExtractUserFromCtx(c)
	c.JSON(http.StatusOK, transformUserFromUserDB(user, u.ID))
}

func (h Handler) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	user, err := h.userStorage.GetOne(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userStorage.Delete(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transformUserFromUserDB(user))
}

func (h Handler) updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var req UpdateUserRequest
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := h.userStorage.GetOne(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Username = req.Username
	user.Role = req.Role
	user.IsActive = req.IsActive
	err = h.userStorage.Update(c, user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user, err = h.userStorage.GetOne(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transformUserFromUserDB(user))
}

func (h Handler) createUser(c *gin.Context) {
	var req CreateUserRequest
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u, _ := ExtractUserFromCtx(c)
	user := user.User{
		Username:  req.Username,
		IsActive:  req.IsActive,
		Role:      req.Role,
		CreatedBy: u.Username,
	}
	user, err = h.userStorage.Create(c, user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transformUserFromUserDB(user))
}
