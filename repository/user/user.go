package user

import (
	"context"
	"fmt"
	"strings"

	models "github.com/hjcalderon10/bunny-backend/model"
	"github.com/hjcalderon10/bunny-backend/repository/storage"
	"github.com/kisielk/sqlstruct"
)

var repo *userRepository

type userRepository struct {
	db storage.StorageRepository
}

func New() *userRepository {
	if repo == nil {
		repo = &userRepository{storage.NewPostgres()}
	}
	return repo
}

func (repo userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users := []models.User{}
	rows, err := repo.db.Raw(get_all_users)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		user = models.User{}
		sqlstruct.Scan(&user, rows)

		users = append(users, user)
	}

	err = rows.Err()

	return users, err
}

func (repo userRepository) CreateUser(ctx context.Context, user models.User) (uint16, error) {
	var id uint16
	rows, err := repo.db.Raw(create_user, user.Name, user.NickName, user.ImgURL)
	if err == nil {
		if rows.Next() {
			rows.Scan(&id)
		}
	}
	return id, err
}

func (repo userRepository) ReadUser(ctx context.Context, userID models.UserID) (models.User, error) {
	user := models.User{}
	rows, err := repo.db.Raw(get_user, userID)

	if err == nil {
		defer rows.Close()

		if rows.Next() {
			sqlstruct.Scan(&user, rows)
		}

		err = rows.Err()
	}

	return user, err
}

func (repo userRepository) UpdateUser(ctx context.Context, user models.User) error {
	str, err := repo.buildUpdateQuery(user)
	if err == nil {
		_, err = repo.db.Exec(fmt.Sprintf(update_user, str), user.ID)
	}
	return err
}

func (repo userRepository) DeleteUser(ctx context.Context, userID models.UserID) error {
	_, err := repo.db.Exec(delete_user, userID)
	return err
}

func (repo userRepository) buildUpdateQuery(user models.User) (string, error) {
	var strBuff strings.Builder

	if user.Name != "" {
		fmt.Fprintf(&strBuff, "name='%s', ", user.Name)
	}
	if user.ImgURL != "" {
		fmt.Fprintf(&strBuff, "img_url='%s', ", user.ImgURL)
	}
	if user.NickName != "" {
		fmt.Fprintf(&strBuff, "nickname='%s', ", user.NickName)
	}

	if strBuff.Len() == 0 {
		return "", fmt.Errorf("No properties to update")
	}
	return strBuff.String()[:strBuff.Len()-2], nil
}
