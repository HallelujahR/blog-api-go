package userservice

import (
	securityQuestionDao "blog/model/dao/security_question"
	dao "blog/model/dao/user"
	usersecurityquestionservice "blog/model/service/user_security_question"
	"errors"
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type InputParams struct {
	User             dao.User
	SecurityQuestion []securityQuestionDao.UserSecurityQuestion
}

// 注册用户
func Register(g *gin.Context) (*dao.User, error) {
	var params InputParams
	err := g.BindJSON(&params)

	if err != nil {
		return nil, err
	}
	// 验证参数
	if err := CheckParams(g, params); err != nil {
		return nil, err
	}

	//密码加密
	password, err := HandlePassword(params.User.Password)
	if err != nil {
		return nil, err
	}
	params.User.Password = password
	//插入数据
	data, err := dao.CreateUser(g, params.User)
	if err != nil {
		return nil, err
	}
	// 处理密保问题
	err = usersecurityquestionservice.HandleSecurityQuestion(g, data.ID, params.SecurityQuestion)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CheckParams(g *gin.Context, params InputParams) error {
	userData := params.User
	// 验证手机号/邮箱是否已经存在
	record, err := dao.GetUserByCond(g, userData.Phonenumber, userData.Email)
	if err != nil {
		return err
	}
	if record.ID != 0 {
		return errors.New("手机号/邮箱已经存在")
	}

	//验证手机号是否合法
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	matched := reg.MatchString(userData.Phonenumber)
	if matched == false {
		return errors.New("手机号格式不正确")
	}

	//验证邮箱是否合法
	_, err = mail.ParseAddress(userData.Email)
	if err != nil {
		return errors.New("邮箱格式错误")
	}

	//验证密码是否合法
	err = CheckPassword(userData.Password)
	if err != nil {
		return err
	}
	//验证验证码是否正确

	return nil
}

// 密码强度必须为字⺟⼤⼩写+数字+符号，9位以上
func CheckPassword(password string) error {
	if len(password) < 6 || len(password) > 18 {
		return errors.New("密码长度必须在6-18位之间")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		return errors.New("密码中需要包含数字")
	}
	if b, err := regexp.MatchString(a_z, password); !b || err != nil {
		return errors.New("密码中需要包含小写英文字母")
	}
	if b, err := regexp.MatchString(A_Z, password); !b || err != nil {
		return errors.New("密码中需要包含大写英文字母")
	}
	if b, err := regexp.MatchString(symbol, password); !b || err != nil {
		return errors.New("密码中需要包含至少一个特殊符号")
	}
	return nil
}

// 密码加密
func HandlePassword(password string) (string, error) {
	// GenerateFromPassword 返回密码的bcrypt哈希值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
