package converter

import (
	"fanland/db/dao"
	"fanland/model"
)

func ConvertToUserDO(user *model.User) *dao.UserDO {
	return nil
}
func ConvertToUser(userDO *dao.UserDO) *model.User {
	user := &model.User{
		Id:         userDO.Id,
		UserName:   userDO.UserName,
		UserDesc:   userDO.UserDesc,
		AvatarUrl:  userDO.AvatarUrl,
		UserHash:   userDO.UserHash,
		CreateTime: userDO.CreateTime,
		UpdateTime: userDO.UpdateTime,
	}
	return user
}
