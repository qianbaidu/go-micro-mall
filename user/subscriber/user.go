package subscriber

import (
	"context"
	"github.com/qianbaidu/go-micro-mall/common/util/log"

	user "github.com/qianbaidu/go-micro-mall/user/proto/user"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
