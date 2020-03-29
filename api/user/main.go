package main

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/qianbaidu/go-micro-mall/api/user/handler"
	cfgUtil "github.com/qianbaidu/go-micro-mall/common/config/util"
	"github.com/qianbaidu/go-micro-mall/common/tracer"
	"github.com/qianbaidu/go-micro-mall/common/util/log"
	"github.com/qianbaidu/go-micro-mall/common/warapper/tracer/opentracing/gin2micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"

	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/web"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/opentracing/opentracing-go"
	"time"

	"github.com/qianbaidu/go-micro-mall/common/token"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/qianbaidu/go-micro-mall/api/user/docs"
)

const appName = "user-api"

var (
	appCfg = &cfgUtil.AppCfg{}
)

func init() {
	appCfg = cfgUtil.InitGetAppCfg(appName)
}

func main() {
	// token
	token := &token.Token{}

	// tracer
	gin2micro.SetSamplingFrequency(50)
	t, io, err := tracer.NewTracer(appCfg.Name, cfgUtil.GetJaegerAddress())
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := web.NewService(
		web.Name(appCfg.Name),
		web.Version("lastest"),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.MicroService(grpc.NewService()),
		web.Flags(cli.StringFlag{
			Name:   "consul_address",
			Usage:  "consul address for K/V",
			EnvVar: "CONSUL_ADDRESS",
			Value:  cfgUtil.GetConsulAddress(),
		}),
		web.Action(func(ctx *cli.Context) {
			token.InitConfig("jwt-key", "key")
		}),
		web.Address(appCfg.Addr()),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	hystrix.DefaultTimeout = 15000

	sClient := hystrixplugin.NewClientWrapper()(service.Options().Service.Client())
	sClient.Init(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
		client.Retries(3),
		client.DialTimeout(time.Second*15),
		client.RequestTimeout(time.Second*15),

		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Info(req.Method(), retryCount, " client retry")
			return true, nil
		}),

	)

	//pub := micro.NewPublisher("/test", sClient)

	apiService := handler.New(sClient, token)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r := router.Group("/user")
	r.Use(gin2micro.TracerWrapper)
	r.GET("/test", apiService.Anything)
	r.POST("/register", apiService.Create)
	r.POST("/login", apiService.Login)


	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
