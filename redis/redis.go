package redis

import (
    "errors"
    "gopkg.in/redis.v5"
    "time"
)

type RedisEntry struct {
    Addr        string
    Password    string
    DB          int
    MaxRetries  int
    PoolSize    int
    IdleTimeout time.Duration
}

// EXPORT
//
var Redis *redis.Client

// RegisterRedis
func RegisterRedis(redisEntry *RedisEntry) error {
    if redisEntry == nil {
        panic("register redis err,redis entry nil")
    }
    // new client
    client := redis.NewClient(&redis.Options{
        Addr:        redisEntry.Addr,
        Password:    redisEntry.Password,
        DB:          redisEntry.DB,
        MaxRetries:  redisEntry.MaxRetries,
        IdleTimeout: redisEntry.IdleTimeout,
        PoolSize:    redisEntry.PoolSize,
    })
    if _, err := client.Ping().Result(); err != nil {
        panic("failed to connect redis:" + err.Error())
    }
    // redis
    Redis = client
    return nil
}

// RedisExist
func RedisExist(key string) (bool, error) {
    if Redis == nil {
        return false, nil
    }
    boolCmd := Redis.Exists(key)
    isExist, err := boolCmd.Result()
    if err != nil || !isExist {
        return false, err
    }
    return true, nil
}

// RedisSet
func RedisSet(key, value string, expiration time.Duration) error {
    if Redis == nil {
        return errors.New("redis is empty points")
    }
    return Redis.Set(key, value, expiration).Err()
}

// RedisGet
func RedisGet(key string) (string, error) {
    if Redis == nil {
        return "", errors.New("redis is empty points")
    }
    // is exist
    boolCmd := Redis.Exists(key)
    isExist, err := boolCmd.Result()
    if err != nil || !isExist {
        return "", nil
    }
    return Redis.Get(key).Result()
}

// RedisDel
func RedisDel(key string) error {
    if Redis == nil {
        return errors.New("redis is empty points")
    }
    return Redis.Del(key).Err()
}

// RedisSetNx
func RedisSetNx(key, value string, expiration time.Duration) bool {
    if Redis == nil {
        return false
    }
    bResult, err := Redis.SetNX(key, value, expiration).Result()
    if err == nil || bResult {
        return true
    }
    return false
}

// RedisIncr
func RedisIncr(key string) (int64, error) {
    if Redis == nil {
        return 0, errors.New("redis is empty points")
    }
    return Redis.Incr(key).Result()
}

// RedisExpire
func RedisExpire(key string, expiration time.Duration) (bool, error) {
    if Redis == nil {
        return false, errors.New("redis is empty points")
    }
    return Redis.Expire(key, expiration).Result()
}
