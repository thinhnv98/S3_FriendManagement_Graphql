package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"S3_FriendManagement_Graphql/graph/generated"
	"S3_FriendManagement_Graphql/graph/graphqlmodels"
	"S3_FriendManagement_Graphql/modelss/api-models"
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

func (r *mutationResolver) CreateFriend(ctx context.Context, input graphqlmodels.Friends) (*graphqlmodels.IsSuccess, error) {
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
	return &graphqlmodels.IsSuccess{
		Success: true,
	}, nil
}

func (r *mutationResolver) FriendList(ctx context.Context, input graphqlmodels.Email) (*graphqlmodels.FriendList, error) {
	//Decode request body
	friendRequest := api_models.FriendGetFriendListRequest{
		Email: input.Email,
	}

	//Validation
	if err := friendRequest.Validate(); err != nil {
		return nil, err
	}

	//Check existed email and get ID by email
	userID, err := r.GetFriendListValidation(friendRequest.Email)
	if err != nil {
		return nil, err
	}

	//Call services
	friendList, err := r.IFriendService.GetFriendListByID(userID)
	if err != nil {
		return nil, err
	}

	//Response
	return &graphqlmodels.FriendList{
		Success: true,
		Friends: friendList,
		Count:   len(friendList),
	}, nil
}

func (r *mutationResolver) CommonFriends(ctx context.Context, input graphqlmodels.Friends) (*graphqlmodels.FriendList, error) {
	//Decode request body
	friendRequest := api_models.FriendGetCommonFriendsRequest{
		Friends: input.Friends,
	}

	//Validation
	if err := friendRequest.Validate(); err != nil {
		return nil, err
	}

	//Check Existed email and get IDList
	userIDList, err := r.GetCommonFriendListValidation(friendRequest.Friends)
	if err != nil {
		return nil, err
	}

	//Call services
	friendList, err := r.IFriendService.GetCommonFriendListByID(userIDList)
	if err != nil {
		return nil, err
	}

	//Response
	return &graphqlmodels.FriendList{
		Success: true,
		Friends: friendList,
		Count:   len(friendList),
	}, nil
}

func (r *mutationResolver) Subscribe(ctx context.Context, input graphqlmodels.RequestTarget) (*graphqlmodels.IsSuccess, error) {
	//Decode request body
	subscriptionRequest := api_models.CreateSubscriptionRequest{
		Requestor: input.Requestor,
		Target:    input.Target,
	}

	//Validate request
	if err := subscriptionRequest.Validate(); err != nil {
		return nil, err
	}

	//Validate and get UserID by email
	userIDList, err := r.CreateSubscribeValidation(subscriptionRequest)
	if err != nil {
		return nil, err
	}
	//Create input services model
	modelServiceInput := &api_models.SubscriptionServiceInput{
		Requestor: userIDList[0],
		Target:    userIDList[1],
	}
	//Call services
	if err := r.ISubscriptionService.CreateSubscription(modelServiceInput); err != nil {
		return nil, err
	}

	// Response
	return &graphqlmodels.IsSuccess{
		Success: true,
	}, nil
}

func (r *mutationResolver) BlockUpdate(ctx context.Context, input graphqlmodels.RequestTarget) (*graphqlmodels.IsSuccess, error) {
	//Decode request body
	blockingRequest := api_models.BlockingRequest{
		Requestor: input.Requestor,
		Target:    input.Target,
	}

	// Validate request
	if err := blockingRequest.Validate(); err != nil {
		return nil, err
	}
	// Validate and get UserID by email
	userIDList, err := r.createBlockingValidation(blockingRequest)
	if err != nil {
		return nil, err
	}

	//Create block services input model
	blockingServiceInput := &api_models.BlockingServiceInput{
		Requestor: userIDList[0],
		Target:    userIDList[1],
	}

	//Call services
	if err := r.IBlockingService.CreateBlocking(blockingServiceInput); err != nil {
		return nil, err
	}

	//Response
	return &graphqlmodels.IsSuccess{
		Success: true,
	}, nil
}

func (r *mutationResolver) RetrieveEmailReceiveUpdate(ctx context.Context, input graphqlmodels.SendMail) (*graphqlmodels.ReceiveUpdateEmailList, error) {
	//decode request body
	emailReceiveUpdateRequest := api_models.EmailReceiveUpdateRequest{
		Sender: input.Sender,
		Text:   input.Text,
	}

	// Validate request body
	if err := emailReceiveUpdateRequest.Validate(); err != nil {
		return nil, err
	}

	// Check existed email and get userID
	senderID, err := r.GetEmailsReceiveUpdateValidation(emailReceiveUpdateRequest.Sender)
	if err != nil {
		return nil, err
	}

	//Call services
	recipientList, err := r.IFriendService.GetEmailsReceiveUpdate(senderID, emailReceiveUpdateRequest.Text)
	if err != nil {
		return nil, err
	}

	// Response
	return &graphqlmodels.ReceiveUpdateEmailList{
		Success:    true,
		Recipients: recipientList,
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
func (r *mutationResolver) GetEmailsReceiveUpdateValidation(email string) (int, error) {
	userID, err := r.IUserService.GetUserIDByEmail(email)
	if err != nil {
		return 0, err
	}
	if userID == 0 {
		return 0, errors.New("the sender does not exist")
	}
	return userID, nil
}
func (r *mutationResolver) createBlockingValidation(blockingRequest api_models.BlockingRequest) ([]int, error) {
	// Get user id of the requestor
	requestorUserID, err := r.IUserService.GetUserIDByEmail(blockingRequest.Requestor)

	if err != nil {
		return nil, err
	}
	if requestorUserID == 0 {
		return nil, errors.New("the requestor does not exist")
	}

	// Get user id of the target
	targetUserID, err := r.IUserService.GetUserIDByEmail(blockingRequest.Target)
	if err != nil {
		return nil, err
	}
	if targetUserID == 0 {
		return nil, errors.New("the target does not exist")
	}

	//Check blocked
	blocked, err := r.IBlockingService.IsExistedBlocking(requestorUserID, targetUserID)
	if err != nil {
		return nil, err
	}
	if blocked {
		return nil, errors.New("target's email have already been blocked by requestor's email")
	}
	return []int{requestorUserID, targetUserID}, nil
}
func (r *mutationResolver) CreateSubscribeValidation(subscriptionRequest api_models.CreateSubscriptionRequest) ([]int, error) {
	//Check requestor email
	requestorUSerID, err := r.IUserService.GetUserIDByEmail(subscriptionRequest.Requestor)
	if err != nil {
		return nil, err
	}
	if requestorUSerID == 0 {
		return nil, errors.New("requestor email does not exist")
	}

	//Check target email
	targetUserID, err := r.IUserService.GetUserIDByEmail(subscriptionRequest.Target)
	if err != nil {
		return nil, err
	}
	if targetUserID == 0 {
		return nil, errors.New("target email does not exist")
	}

	//Check subscription existed
	exist, err := r.ISubscriptionService.IsExistedSubscription(requestorUSerID, targetUserID)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("those email address have already subscribed the each other")
	}

	//Check blocked
	blocked, err := r.ISubscriptionService.IsBlockedByOtherEmail(requestorUSerID, targetUserID)
	if err != nil {
		return nil, err
	}
	if blocked {
		return nil, errors.New("those emails have already been blocked by the each other")
	}
	return []int{requestorUSerID, targetUserID}, nil
}
func (r *mutationResolver) GetCommonFriendListValidation(friends []string) ([]int, error) {
	//check first email
	firstUserID, err := r.IUserService.GetUserIDByEmail(friends[0])
	if err != nil {
		return nil, err
	}
	if firstUserID == 0 {
		return nil, errors.New("first email does not exist")
	}

	secondUserID, err := r.IUserService.GetUserIDByEmail(friends[1])
	if err != nil {
		return nil, err
	}
	if secondUserID == 0 {
		return nil, errors.New("second email does not exist")
	}
	return []int{firstUserID, secondUserID}, nil
}
func (r *mutationResolver) GetFriendListValidation(email string) (int, error) {
	//Check first email valid
	userID, err := r.IUserService.GetUserIDByEmail(email)

	if err != nil {
		return 0, err
	}
	if userID == 0 {
		return 0, errors.New("email does not exist")
	}

	return userID, nil
}
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
