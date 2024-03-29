package models

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"kabao/logs"
	"strconv"
	"time"
)

var (
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
	REDIS_AUTH  string
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis::host")
	REDIS_DB, _ = beego.AppConfig.Int("redis::db")
	REDIS_AUTH = beego.AppConfig.String("redis::auth")
	if REDIS_HOST == "" {
		REDIS_HOST = "172.98.201.182:6379"
	}
	if REDIS_AUTH == "" {
		REDIS_AUTH = "19901124ubuntu"
	}
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST, redis.DialPassword(REDIS_AUTH))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

//检查需要注册的手机号共发送了几次验证码了
func CheckPhoneTimes(phone string) (int, error) {
	rc := RedisClient.Get()
	defer rc.Close()
	count, err := redis.Int(rc.Do("INCR", "requests:"+phone))
	if err != nil {
		return 0, err
	}
	//设置超时12小时
	if count == 1 {
		if ok, _ := redis.Int(rc.Do("EXPIRE", "requests:"+phone, 12*3600)); ok != 1 {
			logs.Error("Redis: Set Expire Error " + phone)
		}
	}
	return count, nil
}

//写入验证码
func SetSMSCode(phone string, code string) error {
	rc := RedisClient.Get()
	defer rc.Close()
	if _, err := rc.Do("SET", "sms:"+phone, code); err != nil {
		return err
	}
	if ok, _ := redis.Int(rc.Do("EXPIRE", "sms:"+phone, 300)); ok != 1 {
		logs.Error("Redis: SetSMSCode Set Expire Error " + phone)
	}
	return nil
}

//获取验证码
func GetSMSCode(phone string) (string, error) {
	rc := RedisClient.Get()
	defer rc.Close()
	code, err := redis.String(rc.Do("GET", "sms:"+phone))
	if err != nil {
		return "", err
	}
	return code, nil
}

//设置登录操作的token
func SetUserTokenRedis(tokenid int, token string) error {
	rc := RedisClient.Get()
	defer rc.Close()
	if _, err := rc.Do("SET", "token:"+strconv.Itoa(tokenid), token); err != nil {
		return err
	}
	if ok, _ := redis.Int(rc.Do("EXPIRE", "token:"+strconv.Itoa(tokenid), 24*3600)); ok != 1 {
		logs.Error("Redis: SetUserTokenRedis Set Expire Error " + strconv.Itoa(tokenid))
	}
	return nil
}

//获取登录的token
func GetUserTokenRedis(tid int) (string, error) {
	rc := RedisClient.Get()
	defer rc.Close()
	token, err := redis.String(rc.Do("GET", "token:"+strconv.Itoa(tid)))
	if err != nil {
		return "", err
	}
	return token, nil
}
