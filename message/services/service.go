package services

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"message/proto/message"
)

type Services struct {
	Log *logrus.Logger
}

func (s Services) SendMessage(ctx context.Context, request *message.MessageRequest, response *message.MessageResponse) error {
	panic(fmt.Errorf("not implemented"))
}
