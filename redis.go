// +----------------------------------------------------------------------
// | redis
// +----------------------------------------------------------------------
// | User: Lengnuan <25314666@qq.com>
// +----------------------------------------------------------------------
// | Date: 2021年06月25日
// +----------------------------------------------------------------------

package gotool

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

// Addr "localhost:6379",
// Password no password set
// DB use default DB
type RedisClient struct {
	Addr     string
	Password string
	DB       int
}

var instance *redis.Client

var once sync.Once

func (r *RedisClient) getInstance() *redis.Client {
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:     r.Addr,
			Password: r.Password,
			DB:       r.DB,
		})
	})
	return instance
}

// mset 批量写入
func (r *RedisClient) StrMSet(pairs ...interface{}) string {
	result, err := r.getInstance().MSet(pairs...).Result()
	if err != nil {
		log.Fatal(pairs, err)
	}

	return result
}

func (r *RedisClient) StrMGet(keys ...string) []interface{} {
	return r.getInstance().MGet(keys...).Val()
}

// set 下的增加数据
func (r *RedisClient) SetAdd(key string, value []string) int64 {
	return r.getInstance().SAdd(key, value).Val()
}

// 获取set下的 value，返回数组
func (r *RedisClient) SetMembersReturnVal(key string) []string {
	return r.getInstance().SMembers(key).Val()
}

func (r *RedisClient) StrSet(key, value string) string {
	res := r.getInstance().Set(key, value, 0)
	return res.Val()
}

func (r *RedisClient) StrGet(key string) string {
	res := r.getInstance().Get(key)
	return string(res.Val())
}

func (r *RedisClient) SortedSetAdd(key string, values ...redis.Z) int64 {
	res := r.getInstance().ZAdd(key, values...)
	return res.Val()
}

func (r *RedisClient) SortedZrangeWithScores(key string) []redis.Z {
	res := r.getInstance().ZRangeWithScores(key, 0, -1).Val()
	return res
}

// 删除key
func (r *RedisClient) DeleteKey(key string) int64 {
	return r.getInstance().Del(key).Val()
}

func (r *RedisClient) GetKeysCount(pattern string) int {
	list := r.getInstance().Keys(pattern).Val()
	return len(list)
}

func (r *RedisClient) MExpire(duration int64, keys ...string) {
	pipe := r.getInstance().Pipeline()
	for _, v := range keys {
		pipe.Expire(v, time.Duration(duration)*24*time.Hour)
	}
	_, err := pipe.Exec()
	log.Println("MExpire:", err)
}