package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"S3_FriendManagement_Graphql/graph/generated"
	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	api_models "S3_FriendManagement_Graphql/modelss/api-models"
	"context"
	"errors"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input graphqlmodels.NewUser) (*graphqlmodels.User, error) {
	userRequest := api_models.UserRequest{
		Email: input.Email,
	}

	//Validation
	if err := userRequest.Validate(); err != nil {
		return nil, err
	}

	if err := r.IsExistedUser(userRequest.Email); err != nil {
		return nil, err
	}

	//Convert to services input api-orm
	userServiceInp := &api_models.UserServiceInput{
		Email: userRequest.Email,
	}

	//Call services
	userID, err := r.IUserService.CreateUser(userServiceInp)
	if err != nil {
		return nil, err
	}

	return &graphqlmodels.User{
		ID:    userID,
		Email: input.Email,
	}, nil
}

func (r *mutationResolver) CreateFriend(ctx context.Context, input graphqlmodels.Friends) (*graphqlmodels.IsFriend, error) {
	// Decode request body
	friendRequest := api_models.FriendConnectionRequest{}
	for _, friend := range input.Friends {
		friendRequest.Friends = append(friendRequest.Friends, friend)
	}

	//Validation
	if err := friendRequest.Validate(); err != nil {
		return nil, err
	}

	// Validate before creating friend
	IDs, err := r.CreateFriendValidation(friendRequest)
	if err != nil {
		return nil, err
	}

	//Model UserIDs services input
	friendsInputModel := &api_models.FriendsServiceInput{
		FirstID:  IDs[0],
		SecondID: IDs[1],
	}

	//Call services to create friend connection
	if err := r.IFriendService.CreateFriend(friendsInputModel); err != nil {
		return nil, err
	}

	//Response
	return &graphqlmodels.IsFriend{
		Success: true,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, id *int, email *string) ([]*graphqlmodels.User, error) {
	//handle input
	var userIDVal int
	var emailVal string

	if id == nil && email == nil {
		userIDVal = 0
		emailVal = ""
	}
	if id == nil && email != nil {
		userIDVal = 0
		emailVal = *email
	}
	if email == nil && id != nil {
		userIDVal = *id
		emailVal = ""
	}

	//Call service
	if userIDVal == 0 && emailVal == "" {
		users, err := r.IUserService.GetAllUser()
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	users, err := r.IUserService.GetUserWithCondition(userIDVal, emailVal)
	if err != nil {
		return nil, err
	}
	return users, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateFriendValidation(friendConnectionRequest api_models.FriendConnectionRequest) ([]int, error) {
	//Check first email valid
	firstUserID, err := r.IUserService.GetUserIDByEmail(friendConnectionRequest.Friends[0])

	if err != nil {
		return nil, err
	}
	if firstUserID == 0 {
		return nil, errors.New("the first email does not exist")
	}

	//Check first email valid
	secondUserID, err := r.IUserService.GetUserIDByEmail(friendConnectionRequest.Friends[1])

	if err != nil {
		return nil, err
	}
	if secondUserID == 0 {
		return nil, errors.New("the second email does not exist")
	}

	// Check friend connection exists
	existed, err := r.IFriendService.IsExistedFriend(firstUserID, secondUserID)
	if err != nil {
		return nil, err
	}
	if existed {
		return nil, errors.New("friend connection existed")
	}

	//check blocking between 2 emails
	blocked, err := r.IFriendService.IsBlockedByOtherEmail(firstUserID, secondUserID)
	if err != nil {
		return nil, err
	}
	if blocked {
		return nil, errors.New("emails blocked each other")
	}

	return []int{firstUserID, secondUserID}, nil
}
func (r *mutationResolver) IsExistedUser(email string) error {
	//Call services
	existed, err := r.IUserService.IsExistedUser(email)
	if err != nil {
		return err
	}
	if existed {
		return errors.New("this email address existed")
	}
	return nil
}
