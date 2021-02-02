package service

import (
	"project/app/admin/models"
	"project/app/admin/models/bo"
	"project/app/admin/models/dto"
	"project/pkg/jwt"
	"project/utils"
)

type User struct {
}

// Login 返回json web token
func (u *User) Login(p *dto.UserLoginDto) (token string, err error) {
	user := new(models.SysUser)
	user.Username = p.Username

	user.Password = p.Password
	if err = user.Login(); err != nil {
		return "", err
	}
	return jwt.GenToken(user.ID, user.Username)
}

func (u *User) InsertUser(p *dto.InsertUserDto, userID int) (err error) {
	//初始化 user数据
	user := &models.SysUser{
		DeptId:   p.DeptId,
		Email:    p.Email,
		NickName: p.NickName,
		Phone:    utils.Int64ToString(p.Phone),
		Username: p.UserName,
		Enabled:  utils.StrBoolIntoByte(p.Enabled),
		Gender:   utils.StrBoolIntoByte(p.Gender),
		CreateBy: userID,
		UpdateBy: userID,
		IsAdmin:  []byte{0},
	}
	jobs := p.Jobs
	roles := p.Roles

	if err := user.InsertUser(jobs, roles); err != nil {
		return err
	}
	return nil
}

func (u *User) SelectUserInfoList(p *dto.SelectUserInfoArrayDto) (data *bo.UserInfoListBo, err error) {
	user := new(models.SysUser)
	data, err = user.SelectUserInfoList(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *User) DeleteUser(ids *[]int) error {
	user := new(models.SysUser)
	return user.DeleteUser(ids)
}

func (u *User) UpdateUser(p *dto.UpdateUserDto, optionId int) error {
	user := new(models.SysUser)
	return user.UpdateUser(p, optionId)
}

func (u *User) UpdateUserCenter(p *dto.UpdateUserCenterDto, optionId int) (err error) {
	user := new(models.SysUser)
	return user.UpdateUserCenter(p, optionId)
}

func (u *User) SelectUserInfo(p *models.RedisUserInfo) (data *bo.UserCenterInfoBo, err error) {
	user := new(models.SysUser)
	return user.SelectUserInfo(p)
}

func (u *User) UpdatePassWord(p *dto.UpdateUserPassDto, optionId int) (err error) {
	user := new(models.SysUser)
	return user.UpdatePassWord(p, optionId)
}

func (u *User) UpdateAvatar(path string, userId int) (err error) {
	user := new(models.SysUser)
	return user.UpdateAvatar(path, userId)
}
