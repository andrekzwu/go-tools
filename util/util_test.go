package util

import (
	"fmt"
	// "github.com/andrezhz/go-tools/redis"
	"testing"
)

// // TestUUID
// func TestUUID(t *testing.T) {
//     // register redis
//     if err := redis.RegisterRedis(&redis.RedisEntry{
//         Addr:        "xxxxxx:6379",
//         Password:    "123456",
//         DB:          2,
//         PoolSize:    50,
//         IdleTimeout: 60,
//         MaxRetries:  3,
//     }); err != nil {
//         t.Errorf("refister redis err:%v", err)
//         return
//     }
//     uuidStr, err := GenerateUUID()
//     if err != nil {
//         t.Errorf("generate uuid info err:%v", err)
//         return
//     }It's a nice day today
//     fmt.Println(uuidStr)
// }

// // TestIsSqlInjectionAttack
// func TestIsSqlInjectionAttack(t *testing.T) {
//     fmt.Println("---------", IsSqlInjectionAttack("网易严选（苏打优选）,400-0368-163"))
// }

// TestHanzi2Pinyin
func TestHanzi2Pinyin(t *testing.T) {
	fmt.Println(Hanzi2Pinyin("我是邬可正andre019z"))
}
