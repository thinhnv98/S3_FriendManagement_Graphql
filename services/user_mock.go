package services

import (
	"S3_FriendManagement_Graphql/modelss/api-models"
	"S3_FriendManagement_Graphql/modelss/orm"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (_self mockUserRepo) CreateUser(userRepoInput *api_models.UserRepoInput) (int, error) {
	args := _self.Called(userRepoInput)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetAllUser() (orm.UseremailSlice, error) {
	args := _self.Called()
	r0 := args.Get(0).(orm.UseremailSlice)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetUserWithCondition(id int, email string) (orm.UseremailSlice, error) {
	args := _self.Called()
	r0 := args.Get(0).(orm.UseremailSlice)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetUserIDByEmail(email string) (int, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) IsExistedUser(email string) (bool, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetEmailListByIDs(userIDs []int) ([]string, error) {
	args := _self.Called(userIDs)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) GetUserIDsByEmails(emails []string) ([]int, error) {
	args := _self.Called(emails)
	r0 := args.Get(0).([]int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserRepo) CheckInvalidEmails(emails []string) ([]string, error) {
	args := _self.Called(emails)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
