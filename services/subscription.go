package services

import (
	"S3_FriendManagement_Graphql/modelss/api-models"
	"S3_FriendManagement_Graphql/repositories"
)

type ISubscriptionService interface {
	CreateSubscription(*api_models.SubscriptionServiceInput) error
	IsExistedSubscription(int, int) (bool, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
}

type SubscriptionService struct {
	ISubscriptionRepo repositories.ISubscriptionRepo
}

func (_self SubscriptionService) CreateSubscription(subscriptionServiceInput *api_models.SubscriptionServiceInput) error {
	//Create repo input graphqlmodels
	repoInput := &api_models.SubscriptionRepoInput{
		Requestor: subscriptionServiceInput.Requestor,
		Target:    subscriptionServiceInput.Target,
	}
	err := _self.ISubscriptionRepo.CreateSubscription(repoInput)
	return err
}

func (_self SubscriptionService) IsExistedSubscription(requestorID int, targetID int) (bool, error) {
	exist, err := _self.ISubscriptionRepo.IsExistedSubscription(requestorID, targetID)
	return exist, err
}

func (_self SubscriptionService) IsBlockedByOtherEmail(requestorID int, targetID int) (bool, error) {
	blocked, err := _self.ISubscriptionRepo.IsBlockedByOtherEmail(requestorID, targetID)
	return blocked, err
}
