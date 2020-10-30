package user

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
	settings "github.com/hjcalderon10/bunny-backend/setting"
	errors "github.com/hjcalderon10/bunny-backend/util/error"
	"github.com/hjcalderon10/bunny-backend/util/logger"
)

func New(repo IRepo) service {
	return service{
		repo: repo,
	}
}

func (srv service) GetAllUsers(ctx context.Context) ([]models.User, error) {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	users, err := srv.repo.GetAllUsers(ctx)

	if err != nil {
		log.Errorf("[GetAllUsers:%s]", err)
		err = errors.InternalServerError
	}

	return users, err
}

func (srv service) CreateUser(ctx context.Context, user models.User) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	err := srv.repo.CreateUser(ctx, user)

	if err != nil {
		log.Errorf("[CreateUser:%s]", err)
		err = errors.InternalServerError
	}

	return err
}
func (srv service) ReadUser(ctx context.Context, user *models.User) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	userDB, err := srv.repo.ReadUser(ctx, user.ID)
	if err != nil {
		log.Errorf("[ReadUser:%s]", err)
		err = errors.InternalServerError

	} else {
		user.Name = userDB.Name
		user.ImgURL = userDB.ImgURL
		user.IsActive = userDB.IsActive
		user.NickName = userDB.NickName
		user.CreatedAt = userDB.CreatedAt
	}
	return err
}
func (srv service) UpdateUser(ctx context.Context, user models.User) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	err := srv.repo.UpdateUser(ctx, user)

	if err != nil {
		log.Errorf("[UpdateUser:%s]", err)
		err = errors.InternalServerError
	}

	return err
}
func (srv service) DeleteUser(ctx context.Context, user models.User) error {
	log := ctx.Value(settings.Commons.LogKey).(logger.Logger)
	err := srv.repo.DeleteUser(ctx, user.ID)

	if err != nil {
		log.Errorf("[DeleteUser:%s]", err)
		err = errors.InternalServerError
	}

	return err
}
