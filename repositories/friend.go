package repositories

import (
	"S3_FriendManagement_Graphql/modelss/api-models"
	"S3_FriendManagement_Graphql/modelss/orm"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
)

type IFriendRepo interface {
	CreateFriend(*api_models.FriendsRepoInput) error
	GetFriendListByID(int) ([]int, error)
	GetBlockedListByID(int) ([]int, error)
	GetBlockingListByID(int) ([]int, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
	GetSubscriberList(int) ([]int, error)
	GetEmailsFriendOrSubscribedWithNoBlocked(int) ([]int, error)
}

type FriendRepo struct {
	Db *sql.DB
}

func (_self FriendRepo) CreateFriend(friendsRepoInput *api_models.FriendsRepoInput) error {
	friend := orm.Friend{
		FirstID:  int64(friendsRepoInput.FirstID),
		SecondID: int64(friendsRepoInput.SecondID),
	}
	if err := friend.Insert(context.Background(), _self.Db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (_self FriendRepo) GetFriendListByID(userID int) ([]int, error) {
	query := `select first_id, second_id from friends where first_id=$1 or second_id = $1`

	var friendListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var firstID, secondID int
		if err := rows.Scan(&firstID, &secondID); err != nil {
			return nil, err
		}
		if firstID == userID {
			friendListID = append(friendListID, secondID)
		}
		if secondID == userID {
			friendListID = append(friendListID, firstID)
		}
	}
	return friendListID, err
}

func (_self FriendRepo) GetBlockingListByID(userID int) ([]int, error) {
	query := `select target_id from blocks where requestor_id = $1`

	var blockedListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blockedUserID int
		if err := rows.Scan(&blockedUserID); err != nil {
			return nil, err
		}
		blockedListID = append(blockedListID, blockedUserID)
	}
	return blockedListID, err
}

func (_self FriendRepo) GetBlockedListByID(userID int) ([]int, error) {
	query := `select requestor_id from blocks where target_id = $1`

	var blockingListID = make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var blockingUserID int
		if err := rows.Scan(&blockingUserID); err != nil {
			return nil, err
		}
		blockingListID = append(blockingListID, blockingUserID)
	}
	return blockingListID, err
}

func (_self FriendRepo) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(select true from blocks WHERE (
    						    	requestor_id in ($1, $2) 
								    AND 
    						    	target_id in ($1, $2)
    						      ))`
	var isBlocked bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&isBlocked)
	if err != nil {
		return true, err
	}
	if isBlocked {
		return true, nil
	}
	return false, nil
}

func (_self FriendRepo) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	query := `select exists(
    						select true 
    						from friends 
    						where (
    						    	first_id in ($1, $2) 
								    AND 
    						    	second_id in ($1, $2)
    						      )
    						)`
	var existed bool
	err := _self.Db.QueryRow(query, firstUserID, secondUserID).Scan(&existed)
	if err != nil {
		return true, err
	}
	if existed {
		return true, nil
	}
	return false, nil
}

func (_self FriendRepo) GetSubscriberList(userID int) ([]int, error) {
	query := `select requestor_id from subscriptions where target_id=$1`
	subscribers := make([]int, 0)
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, id)
	}
	return subscribers, nil
}

func (_self FriendRepo) GetEmailsFriendOrSubscribedWithNoBlocked(userID int) ([]int, error) {
	query := `select distinct val.ID
			  from
				(
					select ue.ID
					from
						useremails ue
							join friends f
								 on (ue.id = f.first_id or ue.id = f.second_id)
					where ue.id <> $1
					  and (f.first_id = $1 or f.second_id = $1)
					union
					select ue.ID
					from
						subscriptions s
							join useremails ue
								 on s.target_id = ue.id
					where ue.id <> $1
				) as val
				where not exists(
				    select 1
				    from blocks b 
				    where b.requestor_id = val.id
				    and b.target_id = $1
				)`
	rows, err := _self.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	UserIDs := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		UserIDs = append(UserIDs, id)
	}
	return UserIDs, nil
}
