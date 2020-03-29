package handler

import (
	"encoding/json"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qianbaidu/go-micro-mall/common/token"
	"github.com/qianbaidu/go-micro-mall/common/util/log"
	"github.com/qianbaidu/go-micro-mall/common/warapper/tracer/opentracing/gin2micro"
	"github.com/micro/go-micro/client"
	"net/http"

	helloS "github.com/qianbaidu/go-micro-mall/hello/proto/hello"
	userS "github.com/qianbaidu/go-micro-mall/user/proto/user"
)

const name = "go.micro.api.user"

type Resp struct {
	Error   string      `json:"error"`
	Success bool        `json:"success"`
	Msg     string      `json:"message"`
	Data    interface{} `json:"data"`
}

// UserAPIService 服务
type UserAPIService struct {
	jwt    *token.Token
	helloC helloS.ExampleService
	userC  userS.UserService
	//pub    micro.Publisher
}

// New UserAPIService
func New(client client.Client, token *token.Token) *UserAPIService {
	return &UserAPIService{
		jwt:    token,
		helloC: helloS.NewExampleService("go.micro.srv.hello", client),
		userC:  userS.NewUserService("go.micro.srv.user", client),
		//pub:    pub,
	}
}

// Anything 测试demo，调用hello服务和user两个服务
func (s *UserAPIService) Anything(c *gin.Context) {
	log.Info("Received Say.Anything API request")

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Error("get context err")
	}

	//serviceClient := helloS.NewExampleService("go.micro.srv.example", client.DefaultClient)
	//res, err := serviceClient.Call(ctx, &helloS.Request{Name: "xuxu"})
	res, err := s.helloC.Call(ctx, &helloS.Request{Name: "xuxu"})
	if err != nil {
		log.Error("call error : ", err)
		//	c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Info(res)

	//s.pub.Publish(context.TODO(), &helloS.Message{Say: "你好"})

	// userres, err := s.userC.Ping(ctx, &userS.Request{})
	// if err != nil {
	// 	log.Error(err)
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	// log.Info(userres)

	c.JSON(http.StatusOK, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

// Login godoc
// @Summary user register
// @Description user register
// @Router /user/register/ [post]
func (s *UserAPIService) Create(c *gin.Context) {

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Error("get context err")
	}
	var user userS.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusOK, resetError(err, "JWT decode failed."))
		return
	}

	resp, err := s.userC.Create(ctx, &user)
	if err != nil {
		log.Errorf("error : %s  , resp : %s ", err, resp)
		c.AbortWithStatusJSON(http.StatusOK, resetError(err, "register failed."))
		return
	}
	c.JSON(http.StatusCreated, resetRespData("register success", ""))
}

func resetRespData(msg string, data interface{}) Resp {
	return Resp{
		Msg:     msg,
		Error:   "",
		Success: true,
		Data:    data,
	}
}

func resetError(err error, msg string) Resp {
	e := parserRespError(err)
	return Resp{
		Msg:     fmt.Sprintf("%s %s", msg, e),
		Error:   e,
		Success: false,
	}
}

func parserRespError(err error) string {
	type RespError struct {
		Id     string `json:"id"`
		Code   int    `json:"code"`
		Detail string `json:"detail"`
		Status string `json:"status"`
	}
	r := &RespError{}
	if err := json.Unmarshal([]byte(fmt.Sprintf("%s", err)), r); err != nil {
		return ""
	}

	return r.Detail
}

// Login godoc
// @Summary user login
// @Description user login
// @Router /user/login/ [post]
func (s *UserAPIService) Login(c *gin.Context) {

	ctx, ok := gin2micro.ContextWithSpan(c)
	if ok == false {
		log.Error("get context err")
	}
	var user userS.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error("login BindJson error ", err)
		c.AbortWithStatusJSON(http.StatusOK, resetError(err, "JWT decode failed."))
		return
	}

	token, err := s.userC.Auth(ctx, &user)
	if err != nil {
		log.Error("login auth error ", err)
		c.AbortWithStatusJSON(http.StatusOK, resetError(err, "phone or password is wrong."))
		return
	}

	c.JSON(http.StatusCreated, resetRespData("login success", token))
}
