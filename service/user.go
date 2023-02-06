package service

import (
	"assignment_bd/consts"
	_ "assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/utils"
	"errors"
	"github.com/DanPlayer/randomname"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"
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
	global.DB.Create(&NewUser)
}

// Login 执行登录
func Login(in *model.Login) (user *model.User, err error) {
	//检测有没有这个用户名
	println("username:", in.Username)
	err = global.DB.Where("username = ?", in.Username).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未注册")
	}
	//有这个用户名，检测密码是否对
	if utils.EncryptPassword(in.Password, user.Salt) != user.Password {
		return nil, errors.New("帐号或密码不对")
	}
	//todo 后续根据情况处理

	return
}

func UserInfoGetByUserID(userID string) (userinfo *model.UserInfo, err error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, errors.New("获取用户信息失败")
	}
	println("id:", id)
	// 检查 userID 是否存在；若存在，获取用户信息

	user := model.User{}
	err = global.DB.Table("users").Select("username").Where("id = ?", id).Limit(1).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userinfo, gorm.ErrRecordNotFound
	}
	userinfo = &model.UserInfo{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		// TODO 判断是否关注该用户
		IsFollow: false,
	}
	println("userinfo:", userinfo.ID, userinfo.Name, userinfo.FollowCount, userinfo.FollowerCount, userinfo.IsFollow)
	return userinfo, nil
}
