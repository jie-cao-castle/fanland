package manager

import (
	"fanland/common"
	dao "fanland/db"
	"fanland/db/converter"
	"fanland/model"
)

type UserManager struct {
	userDB *dao.UserDB
}

func (manager *UserManager) InitManager(options *common.ServerOptions) {
	manager.userDB = &dao.UserDB{}
	manager.userDB.InitDB(options.DbName)
}

func (manager *UserManager) AddUser(user *model.User) error {
	manager.userDB.Open()
	defer manager.userDB.Close()
	userDO := converter.ConvertToUserDO(user)
	if err := manager.userDB.Insert(userDO); err != nil {
		return err
	}

	return nil
}

func (manager *UserManager) UpdateUser(user *model.User) error {
	manager.userDB.Open()
	defer manager.userDB.Close()
	userDO := converter.ConvertToUserDO(user)
	if err := manager.userDB.Update(userDO); err != nil {
		return err
	}

	return nil
}

func (manager *UserManager) GetUser(userId uint64) (*model.User, error) {
	manager.userDB.Open()
	defer manager.userDB.Close()
	userDO, err := manager.userDB.GetById(userId)
	if err != nil {
		return nil, err
	}
	user := converter.ConvertToUser(userDO)
	return user, nil
}

func (manager *UserManager) GetUserByName(userName string) (*model.User, error) {
	manager.userDB.Open()
	defer manager.userDB.Close()
	userDO, err := manager.userDB.GetByName(userName)
	if err != nil {
		return nil, err
	}
	user := converter.ConvertToUser(userDO)
	return user, nil
}
