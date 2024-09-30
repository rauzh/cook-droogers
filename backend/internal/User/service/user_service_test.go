package service

import (
	"cookdroogers/internal/repo/mocks"
	userErrors "cookdroogers/internal/user/errors"
	"cookdroogers/models/data_builders"
	"database/sql"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/mock"
	"log/slog"
	"testing"
)

type _depFields struct {
	userRepo *mocks.UserRepo
	logger   *slog.Logger
}

type UserServiceSuite struct {
	suite.Suite
}

func _newMockUserDepFields(t provider.T) *_depFields {
	mockUserRepo := mocks.NewUserRepo(t)

	f := &_depFields{
		userRepo: mockUserRepo,
		logger:   slog.Default(),
	}

	return f
}

func (s *UserServiceSuite) TestUserService_CreateOK(t provider.T) {
	t.Title("Create: OK")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(nil, sql.ErrNoRows).Once()
		df.userRepo.EXPECT().Create(mock.Anything, user).Return(nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		err := userService.Create(user)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserServiceSuite) TestUserService_CreateValidationError(t provider.T) {
	t.Title("Create: Validation Error")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Validation error", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithEmail("invalidemail").Build()

		userService := NewUserService(df.userRepo, df.logger)

		err := userService.Create(user)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(userErrors.ErrInvalidEmail, err)
	})
}

func (s *UserServiceSuite) TestUserService_CreateAlreadyTaken(t provider.T) {
	t.Title("Create: Already Taken Error")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Already taken", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(user, nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		err := userService.Create(user)

		sCtx.Assert().Error(err)
		sCtx.Assert().Equal(userErrors.ErrAlreadyTaken, err)
	})
}

func (s *UserServiceSuite) TestUserService_LoginOK(t provider.T) {
	t.Title("Login: OK")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithEmail("uzi@gmail.com").WithPassword("password").Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(user, nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		loggedInUser, err := userService.Login(user.Email, "password")

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, loggedInUser)
	})
}

func (s *UserServiceSuite) TestUserService_LoginInvalidPassword(t provider.T) {
	t.Title("Login: Invalid Password")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Invalid password", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithEmail("uzi@gmail.com").WithPassword("password").Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(user, nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		_, err := userService.Login(user.Email, "wrongpassword")

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, "invalid password")
	})
}

func (s *UserServiceSuite) TestUserService_GetByEmailOK(t provider.T) {
	t.Title("GetByEmail: OK")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithEmail("uzi@gmail.com").Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(user, nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		result, err := userService.GetByEmail(user.Email)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, result)
	})
}

func (s *UserServiceSuite) TestUserService_GetByEmailFailure(t provider.T) {
	t.Title("GetByEmail: Failure")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithEmail("uzi@gmail.com").Build()

		df.userRepo.EXPECT().GetByEmail(mock.Anything, user.Email).Return(nil, sql.ErrConnDone).Once()

		userService := NewUserService(df.userRepo, df.logger)

		result, err := userService.GetByEmail(user.Email)

		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
		sCtx.Assert().Nil(result)
	})
}

func (s *UserServiceSuite) TestUserService_GetOK(t provider.T) {
	t.Title("Get: OK")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithUserID(7).Build()

		df.userRepo.EXPECT().Get(mock.Anything, uint64(7)).Return(user, nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		result, err := userService.Get(uint64(7))

		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(user, result)
	})
}

func (s *UserServiceSuite) TestUserService_GetFailure(t provider.T) {
	t.Title("Get: Failure")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		df.userRepo.EXPECT().Get(mock.Anything, uint64(7)).Return(nil, sql.ErrConnDone).Once()

		userService := NewUserService(df.userRepo, df.logger)

		result, err := userService.Get(uint64(7))

		sCtx.Assert().Error(err, sql.ErrConnDone)
		sCtx.Assert().Nil(result)
	})
}

func (s *UserServiceSuite) TestUserService_UpdateOK(t provider.T) {
	t.Title("Update: OK")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithUserID(7).WithName("new_name").Build()

		df.userRepo.EXPECT().Update(mock.Anything, user).Return(nil).Once()

		userService := NewUserService(df.userRepo, df.logger)

		err := userService.Update(user)

		sCtx.Assert().NoError(err)
	})
}

func (s *UserServiceSuite) TestUserService_UpdateFailure(t provider.T) {
	t.Title("Update: Failure")
	t.Tags("UserService")
	t.Parallel()

	t.WithNewStep("Failure", func(sCtx provider.StepCtx) {
		df := _newMockUserDepFields(t)

		user := data_builders.NewUserBuilder().WithUserID(7).WithName("new_name").Build()

		df.userRepo.EXPECT().Update(mock.Anything, user).Return(sql.ErrConnDone).Once()

		userService := NewUserService(df.userRepo, df.logger)

		err := userService.Update(user)

		sCtx.Assert().ErrorIs(err, sql.ErrConnDone)
	})
}

func TestUserServiceSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(UserServiceSuite))
}
