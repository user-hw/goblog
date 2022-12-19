package service

import (
	"errors"
	"goblog/dao"
	"goblog/models"
	"goblog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {

	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {

		return nil, errors.New("账号密码不正确")

	}
	uid := user.Uid
	// 生成token jwt技术进行生成 A.B.C 生成加密规则 加载数据,过期时间 进行加密算法
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token 未能生成")

	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
