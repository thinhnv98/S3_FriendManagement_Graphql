package repositories

import (
	"database/sql"

	"S3_FriendManagement_Graphql/modelss/api-models"
)

type IBlockingRepo interface {
	CreateBlocking(input *api_models.BlockingRepoInput) error
	IsExistedBlocking(requestorID int, targetID int) (bool, error)
}

type BlockingRepo struct {
	Db *sql.DB
}

func (_self BlockingRepo) CreateBlocking(blocking *api_models.BlockingRepoInput) error {
	query := `insert into blocks(requestor_id, target_id) VALUES ($1, $2)`
	_, err := _self.Db.Exec(query, blocking.Requestor, blocking.Target)
	return err
}

func (_self BlockingRepo) IsExistedBlocking(requestorID int, targetID int) (bool, error) {
	query := `select exists(select true from blocks WHERE requestor_id=$1 AND target_id=$2)`
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
