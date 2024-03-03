package authx

import (
	"context"
	"errors"

	"github.com/bezhai/multi-bot-task/biz/dal/mysql"
	"github.com/bezhai/multi-bot-task/biz/utils/sec"
)

func CheckUser(ctx context.Context, username, password string) error {
	UserDao := mysql.GenDao.User
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
