package user

import (
	"demo01/pkg/avatar"
	"demo01/pkg/gormx"
	"demo01/pkg/jwtx"
	"demo01/pkg/oauth/gitee"
	"demo01/pkg/uploader"
	"demo01/util"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var Service = newService()

func newService() *service {
	return &service{}
}

type service struct {
}

func (a *service) Login(params LoginParamsDto) (*LoginRespDto, error) {
	user, err := Repo.GetByUsername(params.Username)
	if err != nil {
		fmt.Println("Login error", err)
		return nil, errors.New("用户不存在")
	}
	if params.Username != user.Username.String {
		fmt.Println("Login Username error", err)
		return nil, errors.New("用户或密码错误")
	}
	if err := util.ValidatePassword(user.Password, params.Password); err != nil {
		fmt.Println("Login Password error", err)
		return nil, errors.New("用户或密码错误")
	}
	subject := fmt.Sprintf("%s-%s", user.ID, params.Username)
	tokenInfo, err := jwtx.GenerateToken(subject)
	if err != nil {
		return nil, err
	}
	return &LoginRespDto{
		User:      user,
		TokenInfo: tokenInfo,
	}, nil
}

func (a *service) LoginOAuth(code, state string) (*LoginRespDto, error) {
	userInfo, err := gitee.GetUserInfoByCode(code, state)
	if err != nil {
		return nil, err
	}
	fmt.Println("loginOauth service", userInfo)

	account, _ := Repo.GetByUsername(userInfo.Login)
	if account != nil {
		subject := fmt.Sprintf("%s-%s", account.ID, account.Username.String)
		token, _ := jwtx.GenerateToken(subject)
		return &LoginRespDto{
			User:      account,
			TokenInfo: token,
		}, nil
	}
	account = &User{
		ID:       strconv.Itoa(userInfo.Id),
		Avatar:   userInfo.AvatarUrl,
		Username: util.SqlNullString(userInfo.Login),
		Email:    util.SqlNullString(userInfo.Email),
		Website:  userInfo.Blog,
	}
	subject := fmt.Sprintf("%s-%s", userInfo.Id, userInfo.AvatarUrl)
	token, _ := jwtx.GenerateToken(subject)
	err = Repo.Create(account)
	if err != nil {
		return nil, err
	}
	return &LoginRespDto{
		User:      account,
		TokenInfo: token,
	}, nil

}

func (a *service) Get(id string) (*User, error) {
	return Repo.Get(id)
}

func (a *service) isExistsOfUsername(username string) bool {
	user, err := Repo.GetByUsername(username)
	if err == nil && user != nil {
		fmt.Println("isExistsOfUsername", err)
		return true
	}
	return false
}

func (a *service) isExistsOfUserEmail(email string) bool {
	user := Repo.GetByEmail(email)
	if user != nil {
		return true
	}
	return false
}

// 生成头像
// 优先级如下：1. 如果第三方登录带有头像；2. 生成随机默认头像
func (a *service) genAvatar(userId string, avatarUrl string) (string, error) {
	if len(avatarUrl) > 0 {
		return uploader.CopyImage(avatarUrl)
	}
	avatarBytes, err := avatar.Generate(userId)
	if err != nil {
		return "", err
	}
	return uploader.PutImage(avatarBytes)

}

func (a *service) Create(ctx *gin.Context, user CreateParamsDto) (*User, error) {

	username := strings.TrimSpace(user.Username)
	email := strings.TrimSpace(user.Email)

	// 验证用户名
	if err := util.IsValidateUsername(username); err != nil {
		return nil, err
	}
	if a.isExistsOfUsername(username) == true {
		return nil, errors.New("该用户名已存在")
	}

	// 验证密码
	err := util.IsValidatePassword(user.Password, user.RePassword)
	if err != nil {
		return nil, err
	}

	// 验证邮箱
	if err := util.IsValidateEmail(email); err != nil {
		return nil, err
	}
	if a.isExistsOfUserEmail(email) {
		return nil, errors.New("邮箱: " + email + " 已被占用")
	}

	//creatorName, _ := ctx.Get("username")
	//creatorID, _ := ctx.Get("userID")
	u := &User{
		ID:          util.GenUUID(""),
		Username:    util.SqlNullString(username),
		Password:    util.EncodePassword(user.Password),
		Email:       util.SqlNullString(email),
		Phone:       util.SqlNullString(user.Phone),
		Website:     user.Website,
		Level:       user.Level,
		Description: user.Description,
		//CreatorName: creatorName.(string),
		//CreatorID:   creatorID.(int64),
	}

	err = gormx.Tx(func(tx *gorm.DB) error {
		if err := Repo.Create(u); err != nil {
			return err
		}
		avatarUrl, err := a.genAvatar(u.ID, "")
		if err != nil {
			return err
		}

		if err := Repo.UpdateColumn(u.ID, "avatar", avatarUrl); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil

}
