package redis

import (
	"NGB/internal/config"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// TO DO
// 更改为事务

// 更新验证码
func (c *Client) UpdateCode(email string, code string) error {
	err := c.Set(email, code, time.Duration(config.C.User.Code.Expire)*time.Minute)
	if err != nil {
		return err
	}
	return nil
}

// 获取验证码
func (c *Client) GetCode(email string) (string, error) {
	code, err := c.Get(email)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (c *Client) UpdateSendMailTime(email string, SendTime int64) error {
	key := fmt.Sprintf("send_time:%s", email)
	err := c.Set(key, strconv.FormatInt(SendTime, 10), time.Duration(config.C.User.Code.MailFrequency)*time.Minute)
	if err != nil {
		return err
	}
	return nil
}

// 获取最近一次发邮件时间
func (c *Client) GetSendMailTime(email string) (int64, error) {
	key := fmt.Sprintf("send_time:%s", email)
	value, err := c.Get(key)
	if err != nil {
		if err == redis.Nil {
			return -1, nil
		}
		return -1, err
	}
	v, _ := strconv.ParseInt(value, 10, 64)
	return v, nil
}
