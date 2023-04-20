package redis

import (
	"NGB/internal/config"
	"NGB/pkg/logrus"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

// TO DO
// 更改为事务

// 更新验证码
func (c *Client)UpdateCode(email string, code string) error {
	err := c.Set(email, code, 5 * time.Minute)
	if err != nil {
		return err
	}
	return nil
}

// 获取验证码
func (c *Client)GetCode(email string) (string, error) {
	code, err := c.Get(email)
	if err != nil {
		return "", err
	}
	return code, nil
}

// 获取最近一次发邮件时间
func (c *Client)GetSendMailTime(email string) (int, error) {
	key := fmt.Sprintf("send_time:%s", email)
	
}