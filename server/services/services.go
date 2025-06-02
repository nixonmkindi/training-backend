package services

import (
	"errors"

	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

const aclSuffix = "-acl"
const ResponseMsgKey = "respmsg"

type ResponseMessage struct {
	Error   bool     `json:"error"`
	Message []string `json:"message"`
}

func Init() {
	Cache = cache.New(5*time.Minute, 15*time.Minute)

}

func GetCache(key string) (interface{}, bool) {
	return Cache.Get(key)
}

func ClearCache(key string) {
	Cache.Set(key, nil, 1*time.Second)
	/*if Cache == nil {
		return nil
	}
	Cache.Flush()
	return nil*/

}

func GetACLKey(email string) string {
	return email + aclSuffix
}

func HasPermission(email, res string) (bool, error) {
	key := GetACLKey(email)
	_, ok := GetCache(key)
	if !ok {
		return false, errors.New("error getting cached permission")
	}
	return true, nil
}
