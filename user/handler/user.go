package handler

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/qianbaidu/go-micro-mall/common/token"
	"github.com/qianbaidu/go-micro-mall/common/util/log"
	"github.com/qianbaidu/go-micro-mall/common/util/tool"
	db "github.com/qianbaidu/go-micro-mall/user/model"
	pb "github.com/qianbaidu/go-micro-mall/user/proto/user"
	"golang.org/x/crypto/pbkdf2"

	"time"
)

const issuer = "go.micro.srv.auth"

type UserService struct {
	token *token.Token
}

func New(token *token.Token) *UserService {
	return &UserService{token: token}
}

// Create 创建新User
func (ser *UserService) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	var err error
	log.Info(req.Phone, "   ", req.Password)
	if req.Phone == "" || req.Password == "" || len(req.Phone) > 11 {
		return errors.New("incomplete information")
	}

	// 重复检查
	if u, err := db.GetByTel(req.Phone); err == nil && u.Id > 0 {
		log.Errorf("User[%s]already exists.", req.Phone)
		return errors.New("User name already exists.")
	}

	req.CreatedUnix = time.Now().Unix()
	req.UpdatedUnix = time.Now().Unix()
	req.Salt, err = tool.RandomString(10)
	if err != nil {
		log.Error("generate salt RandomString error ")
	}

	req.Password = EncodePasswd(req.Password, req.Salt)
	err = db.CreateUser(req)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

// Delete 删除用户
func (ser *UserService) Delete(ctx context.Context, req *pb.User, resp *pb.Response) error {
	return db.DelUser(req)
}

// Get 获取用户信息
func (ser *UserService) Get(ctx context.Context, req *pb.User, resp *pb.User) (err error) {
	if req.Id > 0 {
		*resp, err = db.GetByID(req.Id)
		if err != nil {
			return err
		}
	} else if len(req.Phone) > 1 {
		*resp, err = db.GetByTel(req.Phone)
		if err != nil {
			return err
		}
	}

	resp.Password = ""
	return nil
}

// GetAll 获取所有用户信息
func (ser *UserService) GetAll(ctx context.Context, req *pb.Request, resp *pb.Users) (err error) {
	resp.Users, err = db.GetAllUsers()
	return err
}

// UpdateInfo 更新用户信息
func (ser *UserService) UpdateInfo(ctx context.Context, req *pb.User, resp *pb.Response) error {
	if req.Id < 1 || len(req.Phone) < 1 {
		return errors.New("Illegal user ID")
	}
	return db.UpdateUserInfo(req)
}

// Auth 验证用户账号密码
func (ser *UserService) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	time.Sleep(time.Second * 10)

	var user pb.User
	var err error
	if len(req.Phone) > 0 {
		user, err = db.GetByTel(req.Phone)
		if err != nil {
			log.Errorf("user[%s] not exists", req.Phone)
			return err
		}
	} else {
		log.Error("parames error， phone is empty")
		return errors.New("phone cannot be empty")
	}

	inputPasswd := EncodePasswd(req.Password, user.Salt)
	log.Infof("input[%s] salt[%s] = [%s] ,db secrt[%s]", req.Password, user.Salt, inputPasswd, user.Password)
	v := ValidatePassword(user.Password, inputPasswd)
	log.Info("ValidatePassword", v)
	if ValidatePassword(user.Password, inputPasswd) {
		var tokenStr string
		expireTime := time.Now().Add(time.Hour * 24).Unix() // 1天后过期
		tokenStr, err = ser.token.Encode(issuer, user.Phone, expireTime)
		if err != nil {
			log.Error("get token error ")
			return err
		}
		resp.Token = tokenStr
		return nil
	}
	log.Error("Incorrect account or password")
	return errors.New("Incorrect account or password")
}

func (ser *UserService) GenToken(user *pb.User) (token string, err error) {
	expireTime := time.Now().Add(time.Hour * 24).Unix()
	token, err = ser.token.Encode(issuer, user.Phone, expireTime)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Ping test
func (ser *UserService) Ping(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	log.Info("Received User.Ping request")
	return nil
}

// EncodePasswd encodes password to safe format.
func EncodePasswd(passwd string, salt string) string {
	newPasswd := pbkdf2.Key([]byte(passwd), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", newPasswd)
}

// ValidatePassword checks if given password matches the one belongs to the user.
func ValidatePassword(dbPasswd string, inputPasswd string) bool {
	return subtle.ConstantTimeCompare([]byte(dbPasswd), []byte(inputPasswd)) == 1
}
