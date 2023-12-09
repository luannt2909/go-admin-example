package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reminder-bot/pkg/user"
	"net/http"
	"strconv"
)

type Handler struct {
	userStorage user.Storage
}

func NewHandler(userStorage user.Storage) Handler {
	return Handler{userStorage: userStorage}
}

func (h Handler) findUsers(c *gin.Context) {
	var req GetListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	p := req.toGetListParams()
	list, count, err := h.userStorage.GetList(c, user.GetListParams{GetListParams: p})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("Content-Range", fmt.Sprintf("%d-%d/%d", p.Offset, p.Limit, count))
	c.JSON(http.StatusOK, transformUsersFromUsersDB(list))
}

func (h Handler) getOneUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	user, err := h.userStorage.GetOne(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transformUserFromUserDB(user))
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
	var req User
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
	var req User
	err := c.ShouldBind(&req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user := user.User{
		Email:    req.Email,
		IsActive: req.IsActive,
		Role:     req.Role,
	}
	user, err = h.userStorage.Create(c, user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, transformUserFromUserDB(user))
}
