package cache

import (
	"fmt"

	"github.com/go-redis/redis/v7"

	"project/common/global"
	"project/utils"
)

// GetUserCache 获取用户缓存
func GetUserCache(keys *[]string, userId int) (cacheMap map[string]*redis.StringCmd) {
	cacheMap = make(map[string]*redis.StringCmd, len(*keys))
	pipe := global.Rdb.TxPipeline()
	for _, k := range *keys {
		cacheMap[k] = pipe.Get(fmt.Sprintf("%s%d", k, userId))
	}
	_, _ = pipe.Exec()
	return
}

// SetUserCache 设置用户缓存
func SetUserCache(userId int, data interface{}, cacheKey string) {
	res, err := utils.StructToJson(data)
	if err != nil {
		return
	}
	global.Rdb.Set(fmt.Sprintf("%s%d", cacheKey, userId), res, 0)
}
