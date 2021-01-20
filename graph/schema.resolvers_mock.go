package graph

import (
	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	"S3_FriendManagement_Graphql/modelss/api-models"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (_self mockUserService) CreateUser(model *api_models.UserServiceInput) (int, error) {
	args := _self.Called(model)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) GetAllUser() ([]*graphqlmodels.User, error) {
	args := _self.Called()
	r0 := args.Get(0).([]*graphqlmodels.User)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) GetUserWithCondition(int, string) ([]*graphqlmodels.User, error) {
	args := _self.Called()
	r0 := args.Get(0).([]*graphqlmodels.User)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) IsExistedUser(email string) (bool, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) GetUserIDByEmail(email string) (int, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) CheckInvalidEmails(emails []string) ([]string, error) {
	args := _self.Called(emails)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

type mockFriendService struct {
	mock.Mock
}

func (_self mockFriendService) CreateFriend(model *api_models.FriendsServiceInput) error {
	args := _self.Called(model)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockFriendService) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	args := _self.Called(firstUserID, secondUserID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetFriendListByID(userID int) ([]string, error) {
	args := _self.Called(userID)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetCommonFriendListByID(userIDList []int) ([]string, error) {
	args := _self.Called(userIDList)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockFriendService) GetEmailsReceiveUpdate(userID int, text string) ([]string, error) {
	args := _self.Called(userID, text)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

type mockSubscriptionService struct {
	mock.Mock
}

func (_self mockSubscriptionService) CreateSubscription(subscriptionServiceInput *api_models.SubscriptionServiceInput) error {
	args := _self.Called(subscriptionServiceInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockSubscriptionService) IsExistedSubscription(requestorid int, targetid int) (bool, error) {
	args := _self.Called(requestorid, targetid)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockSubscriptionService) IsBlockedByOtherEmail(requestorid int, targetid int) (bool, error) {
	args := _self.Called(requestorid, targetid)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

type mockBlockingService struct {
	mock.Mock
}

func (_self mockBlockingService) CreateBlocking(input *api_models.BlockingServiceInput) error {
	args := _self.Called(input)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockBlockingService) IsExistedBlocking(requestorID int, targetID int) (bool, error) {
	args := _self.Called(requestorID, targetID)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
