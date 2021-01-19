package services

import (
	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	"S3_FriendManagement_Graphql/modelss/api-models"
	"S3_FriendManagement_Graphql/repositories"
)

type IUserService interface {
	CreateUser(*api_models.UserServiceInput) (int, error)
	GetAllUser() ([]*graphqlmodels.User, error)
	GetUserWithCondition(int, string) ([]*graphqlmodels.User, error)
	IsExistedUser(string) (bool, error)
	GetUserIDByEmail(string) (int, error)
	CheckInvalidEmails([]string) ([]string, error)
}

type UserService struct {
	IUserRepo repositories.IUserRepo
}

func (_self UserService) CreateUser(userServiceInput *api_models.UserServiceInput) (int, error) {
	//Convert to repo input
	userRepoInput := &api_models.UserRepoInput{
		Email: userServiceInput.Email,
	}

	userID, err := _self.IUserRepo.CreateUser(userRepoInput)
	return userID, err
}

func (_self UserService) GetAllUser() ([]*graphqlmodels.User, error) {
	var users []*graphqlmodels.User
	//Call repo
	userSlice, err := _self.IUserRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	for _, u := range userSlice {
		user := &graphqlmodels.User{
			ID:    int(u.ID),
			Email: u.Email,
		}
		users = append(users, user)
	}
	return users, nil
}

func (_self UserService) GetUserWithCondition(id int, email string) ([]*graphqlmodels.User, error) {
	var users []*graphqlmodels.User
	//Call repo
	userSlice, err := _self.IUserRepo.GetUserWithCondition(id, email)
	if err != nil {
		return nil, err
	}
	for _, u := range userSlice {
		user := &graphqlmodels.User{
			ID:    int(u.ID),
			Email: u.Email,
		}
		users = append(users, user)
	}
	return users, nil
}

func (_self UserService) GetUserIDByEmail(email string) (int, error) {
	result, err := _self.IUserRepo.GetUserIDByEmail(email)
	return result, err
}

func (_self UserService) IsExistedUser(email string) (bool, error) {
	//call repo
	existed, err := _self.IUserRepo.IsExistedUser(email)
	return existed, err
}

func (_self UserService) CheckInvalidEmails(emails []string) ([]string, error) {
	results, err := _self.IUserRepo.CheckInvalidEmails(emails)
	return results, err
}
