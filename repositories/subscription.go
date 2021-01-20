package repositories

import (
	"database/sql"

	"S3_FriendManagement_Graphql/modelss/api-models"
)

type ISubscriptionRepo interface {
	CreateSubscription(*api_models.SubscriptionRepoInput) error
	IsExistedSubscription(int, int) (bool, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
}

type SubscriptionRepo struct {
	Db *sql.DB
}

func (_self SubscriptionRepo) CreateSubscription(model *api_models.SubscriptionRepoInput) error {
	query := `insert into subscriptions(requestor_id, target_id) VALUES ($1, $2)`
	_, err := _self.Db.Exec(query, model.Requestor, model.Target)
	return err
}

func (_self SubscriptionRepo) IsExistedSubscription(requestorID int, targetID int) (bool, error) {
	query := `select exists(select true from subscriptions where requestor_id=$1 AND target_id=$2)`
	var exist bool
	err := _self.Db.QueryRow(query, requestorID, targetID).Scan(&exist)
	if err != nil {
		return true, err
	}
	if exist {
		return true, nil
	}
	return false, nil
}

func (_self SubscriptionRepo) IsBlockedByOtherEmail(requestorID int, targetID int) (bool, error) {
	query := `select exists(select true from blocks where (requestor_id=$1 and target_id=$2) or (requestor_id=$2 and target_id=$1))`
	var isBlock bool
	err := _self.Db.QueryRow(query, requestorID, targetID).Scan(&isBlock)
	if err != nil {
		return true, err
	}
	if isBlock {
		return true, nil
	}
	return false, nil
}
