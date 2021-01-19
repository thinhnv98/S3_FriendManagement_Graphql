package services

import (
	"S3_FriendManagement_Graphql/modelss/api-models"
	"S3_FriendManagement_Graphql/repositories"
)

type IBlockingService interface {
	CreateBlocking(*api_models.BlockingServiceInput) error
	IsExistedBlocking(int, int) (bool, error)
}

type BlockingService struct {
	IBlockingRepo repositories.IBlockingRepo
}

func (_self BlockingService) CreateBlocking(blocking *api_models.BlockingServiceInput) error {
	//Create repo input graphqlmodels
	blockingRepoInputModel := &api_models.BlockingRepoInput{
		Requestor: blocking.Requestor,
		Target:    blocking.Target,
	}
	err := _self.IBlockingRepo.CreateBlocking(blockingRepoInputModel)
	return err
}

func (_self BlockingService) IsExistedBlocking(requestorID int, targetID int) (bool, error) {
	exist, err := _self.IBlockingRepo.IsExistedBlocking(requestorID, targetID)
	return exist, err
}
