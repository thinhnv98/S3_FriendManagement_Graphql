package api_models

import (
	"errors"

	"S3_FriendManagement_Graphql/utils"
)

type CreateSubscriptionRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (_self CreateSubscriptionRequest) Validate() error {
	if _self.Requestor == "" {
		return errors.New("\"requestor\" is required")
	}
	if _self.Target == "" {
		return errors.New("\"target\" is required")
	}

	if _self.Target == _self.Requestor {
		return errors.New("two email addresses must be different")
	}

	isValidFirstEmail, firstErr := utils.IsValidEmail(_self.Requestor)
	if firstErr != nil {
		return errors.New("validate \"requestor\" format failed")
	}
	if !isValidFirstEmail {
		return errors.New("\"requestor\" is not valid. (ex: \"andy@abc.xyz\")")
	}

	isValidSecondEmail, secondErr := utils.IsValidEmail(_self.Target)
	if secondErr != nil {
		return errors.New("validate \"target\" format failed")
	}
	if !isValidSecondEmail {
		return errors.New("\"target\" is not valid. (ex: \"andy@abc.xyz\")")
	}

	return nil
}

//Service
type SubscriptionServiceInput struct {
	Requestor int `json:"requestor"`
	Target    int `json:"target"`
}

//Repository
type SubscriptionRepoInput struct {
	Requestor int `json:"requestor"`
	Target    int `json:"target"`
}
