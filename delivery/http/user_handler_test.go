package http_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	httpUser "go-store/delivery/http"
	"net/http/httptest"
	"testing"
)

func TestForTest(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	httpUser.Ping(ctx)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

//func TestLogin(t *testing.T) {
//	mockUser := domain.User{
//		Username:  "TestUser",
//		Password:  "password",
//		CreatedAt: time.Now(),
//	}
//
//	tempMockUser := mockUser
//
//	mockUsecase := new(mocks.UserUsecase)
//
//	j, err := json.Marshal(tempMockUser)
//	assert.NoError(t, err)
//
//	mockUsecase.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)
//
//	w := httptest.NewRecorder()
//	ctx, router := gin.CreateTestContext(w)
//
//	req, err := http.NewRequest("POST", "/login", strings.NewReader(string(j)))
//	assert.NoError(t, err)
//	router.ServeHTTP(w, req)
//
//
//
//	handler := userHttp.UserHandler{
//		UserUsecase: mockUsecase,
//	}
//
//	handler.Login(ctx)
//
//
//	assert.Equal(t, 200, w.Code)
//	mockUsecase.AssertExpectations(t)
//}
