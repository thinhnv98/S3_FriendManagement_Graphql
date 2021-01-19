package repositories

import (
	"S3_FriendManagement_Graphql/modelss/orm"
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strings"

	"S3_FriendManagement_Graphql/modelss/api-models"
)

type IUserRepo interface {
	CreateUser(*api_models.UserRepoInput) (int, error)
	GetAllUser() (orm.UseremailSlice, error)
	GetUserWithCondition(int, string) (orm.UseremailSlice, error)
	IsExistedUser(string) (bool, error)
	GetUserIDByEmail(string) (int, error)
	GetUserIDsByEmails(emails []string) ([]int, error)
	GetEmailListByIDs(userIDs []int) ([]string, error)
	CheckInvalidEmails([]string) ([]string, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (_self UserRepo) CreateUser(userRepoInput *api_models.UserRepoInput) (int, error) {
	user := orm.Useremail{
		Email: userRepoInput.Email,
	}
	if err := user.Insert(context.Background(), _self.Db, boil.Infer()); err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (_self UserRepo) GetAllUser() (orm.UseremailSlice, error) {
	users, err := orm.Useremails().All(context.Background(), _self.Db)
	return users, err
}

func (_self UserRepo) GetUserWithCondition(id int, email string) (orm.UseremailSlice, error) {
	if id == 0 && email == "" {
		users, err := orm.Useremails().All(context.Background(), _self.Db)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	if id == 0 {
		users, err := orm.Useremails(qm.Where("email=?", email)).All(context.Background(), _self.Db)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	if email == "" {
		users, err := orm.Useremails(qm.Where("id=?", id)).All(context.Background(), _self.Db)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	users, err := orm.Useremails(qm.Where("id=? AND email=?", id, email)).All(context.Background(), _self.Db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (_self UserRepo) GetUserIDByEmail(email string) (int, error) {
	query := `select id from useremails where email=$1`
	var userID int
	err := _self.Db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return userID, nil
}

func (_self UserRepo) IsExistedUser(email string) (bool, error) {
	query := `select exists (select true from useremails where email=$1)`
	var existed bool
	err := _self.Db.QueryRow(query, email).Scan(&existed)
	if err != nil {
		return false, err
	}
	if existed {
		return true, nil
	}
	return false, nil
}

func (_self UserRepo) GetEmailListByIDs(userIDs []int) ([]string, error) {
	if len(userIDs) == 0 {
		return []string{}, nil
	}

	IDList := make([]string, len(userIDs))
	for i, id := range userIDs {
		IDList[i] = fmt.Sprintf("%v", id)
	}
	query := fmt.Sprintf(`select email from useremails where id in (%v)`, strings.Join(IDList, ","))
	rows, err := _self.Db.Query(query)
	if err != nil {
		return nil, err
	}

	emailList := make([]string, 0)
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		emailList = append(emailList, email)
	}
	return emailList, nil
}

func (_self UserRepo) GetUserIDsByEmails(emails []string) ([]int, error) {
	if len(emails) == 0 {
		return []int{}, nil
	}

	emailList := make([]string, len(emails))
	for i, email := range emails {
		emailList[i] = fmt.Sprintf("%v", email)
	}
	query := fmt.Sprintf(`select ID from useremails where email in ('%v')`, strings.Join(emailList, "','"))
	rows, err := _self.Db.Query(query)
	if err != nil {
		return nil, err
	}

	IDList := make([]int, 0)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		IDList = append(IDList, id)
	}
	return IDList, nil
}

func (_self UserRepo) CheckInvalidEmails(emails []string) ([]string, error) {
	if len(emails) == 0 {
		return []string{}, nil
	}
	emailList := make([]string, len(emails))
	for i, email := range emails {
		emailList[i] = fmt.Sprintf("%v", email)
	}
	query := fmt.Sprintf(`select email
								from (
							 		values ('%v')
								)as e(email)
								where not exists(
									select 1
									from useremails ue
									where ue.email = e.email
								)`, strings.Join(emailList, "'),('"))
	rows, err := _self.Db.Query(query)
	if err != nil {
		return nil, err
	}

	Emails := make([]string, 0)
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, err
		}
		Emails = append(Emails, email)
	}
	return Emails, nil
}
