package models

import (  
    "github.com/garyburd/redigo/redis"  
    "github.com/astaxie/beego"  
    "time"  
)  

var (  
    // 定义常量  
    RedisClient     *redis.Pool  
    REDIS_HOST string  
    REDIS_DB   int  
)  

func init() {  
    // 从配置文件获取redis的ip以及db  
    REDIS_HOST = beego.AppConfig.String("redis::host")  
    REDIS_DB, _ = beego.AppConfig.Int("redis::db")  
    // 建立连接池  
    RedisClient = &redis.Pool{  
        // 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值  
        MaxIdle:     beego.AppConfig.DefaultInt("redis::maxidle", 1),  
        MaxActive:   beego.AppConfig.DefaultInt("redis::maxactive", 10),  
        IdleTimeout: 180 * time.Second,  
        Dial: func() (redis.Conn, error) {  
            c, err := redis.Dial("tcp", REDIS_HOST)  
            if err != nil {  
                return nil, err  
            }  
            // 选择db  
            c.Do("SELECT", REDIS_DB)  
            return c, nil  
        },  
    }  
}  
