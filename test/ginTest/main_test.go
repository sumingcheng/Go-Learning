package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"test/ginTest/mocks"
	"test/ginTest/service"
	"testing"
)

func setupRouter(mockUserService *mocks.MockUserService) *gin.Engine {
	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := mockUserService.GetUser(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})
	return r
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	user := &service.User{ID: "1", Name: "Alice"}
	mockUserService.EXPECT().GetUser("1").Return(user, nil)

	r := setupRouter(mockUserService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
}
