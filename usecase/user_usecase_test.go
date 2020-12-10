package usecase_test

// func TestUsecaseCreateUser(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	mockUser := domain.User{
// 		Username:  "TestUser",
// 		Password:  "12345678",
// 		CreatedAt: time.Now(),
// 	}

// 	mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(int64, error).Once()

// 	fmt.Println(mockUserRepo)
// 	u := usecase.NewInteractor(mockUserRepo)

// 	err := u.CreateUser(&mockUser)

// 	assert.NoError(t, err)
// 	assert.Equal(t, mockUser.Username, "TestUser")
// 	mockUser.AssertExpectations(t)
// }
