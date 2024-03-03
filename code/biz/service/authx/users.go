package authx

import (
	"context"
	"errors"

	"github.com/bezhai/multi-bot-task/biz/dal/mysql"
	"github.com/bezhai/multi-bot-task/biz/mysql/model"
	"github.com/bezhai/multi-bot-task/biz/utils/sec"
)

var UserDao = mysql.GenDao.User

func CheckUserPassword(ctx context.Context, username, password string) error {
	user, err := UserDao.WithContext(ctx).
		Where(UserDao.Username.Eq(username)).
		First()
	if err != nil {
		return err
	}
	if !sec.CheckPasswordHash(password, user.PasswordHash) {
		return errors.New("密码错误")
	}
	return nil
}

func CheckUser(ctx context.Context, username string) error {
	_, err := UserDao.WithContext(ctx).
		Where(UserDao.Username.Eq(username)).
		First()
	if err != nil {
		return err
	}

	return nil
}

func Register(ctx context.Context, username, password string) error {

	count, err := UserDao.WithContext(ctx).Where(UserDao.Username.Eq(username)).Count()
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户名已存在")
	}

	passwordHash, err := sec.HashPassword(password)
	if err != nil {
		return errors.New("密码加密失败，请更换密码")
	}

	err = UserDao.WithContext(ctx).Create(&model.User{
		Username:     username,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return err
	}

	return nil
}
