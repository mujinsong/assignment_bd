package service

import (
	"assignment_bd/consts"
	_ "assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/utils"
	"errors"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/DanPlayer/randomname"
	"gorm.io/gorm"
)

// Register 注册用户，验证用户输入的注册信息，并且随机加密盐，给用户加密
// todo 注册，因为和登录一样需要传入的都是用户名的密码，所以我这里传的model.Login(可改)
func Register(username, password string) (out *model.User, err error) {
	in := model.Login{
		Username: username,
		Password: password,
	}

	// 验证用户名合法性
	if utf8.RuneCountInString(username) > consts.MAX_USERNAME_LENGTH ||
		utf8.RuneCountInString(username) <= 0 {
		//c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "非法用户名"})
		return nil, errors.New("非法用户名")
	}

	// 验证密码合法性
	if ok, _ := regexp.MatchString(consts.MIN_PASSWORD_PATTERN, password); !ok {
		//c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "密码长度6-32，由字母大小写下划线组成"})
		return nil, errors.New("密码长度6-32，由字母大小写下划线组成")
	}

	//密码用户名不能为空
	if in.Password == "" || in.Username == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	//查询是否已有该用户名
	result := global.DB.Where("username = ?", in.Username).Take(&out)
	if result.RowsAffected != 0 {
		return nil, errors.New("已有该用户名，请登录或换一个用户名注册")
	} else {
		CreateNewUser(in.Username, in.Password)
	}

	return
}

/*注册成功后将用户信息插入数据库*/
func CreateNewUser(username, password string) {
	NewUser := model.User{
		Username: username,
		Password: password,
		Salt:     utils.RandStr(10),
		CreateAt: time.Now(),
		Name:     randomname.GenerateName(),
		UserFollowCount: model.UserFollowCount{
			FollowCount:   0,
			FollowerCount: 0,
		},
	}
	NewUser.Password = utils.EncryptPassword(NewUser.Password, NewUser.Salt)
	//插入users数据库
	global.DB.Table("users").Create(&NewUser)
}

// Login 执行登录 废弃
func Login(in *model.Login) (user *model.User, err error) {
	//检测有没有这个用户名
	err = global.DB.Where("username = ?", in.Username).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未注册")
	}
	//有这个用户名，检测密码是否对
	if utils.EncryptPassword(in.Password, user.Salt) != user.Password {
		return nil, errors.New("帐号或密码不对")
	}

	return
}

/*
	获取用户信息 通过用户id

uid  当前登录的用户id
userID 表示被查询的用户id
*/
func UserInfoGetByUserID(userID, uid uint64) (userinfo model.UserInfo) {

	user := model.User{}
	global.DB.Table("users").Where("id = ?", userID).Limit(1).Find(&user)
	userinfo = model.UserInfo{
		ID:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      IsFollow(user.ID, uid),
	}
	return
}
