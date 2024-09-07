package redisdb

import (
	"context"
	"fmt"
	"time"
)

// GetList :
func GetList(ctx context.Context, key string) ([]string, error) {
	list, err := rdb.SMembers(ctx, key).Result()
	return list, err
}

// RemoveList :
func RemoveList(ctx context.Context, key string, val interface{}) error {
	_, err := rdb.SRem(ctx, key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// AddList :
func AddList(ctx context.Context, key, val string) error {
	_, err := rdb.SAdd(ctx, key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// TurncateList :
func TurncateList(ctx context.Context, key string) error {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

// AddSession :
//
//	func AddSession(key string, val interface{}, mn int) error {
//		// ss := 1 * time.Hour
//		var (
//			tm = time.Minute
//		)
//		if mn > 0 {
//			tm := time.Duration(mn) * time.Minute
//			fmt.Println(tm)
//		} else {
//			tm = 0
//		}
//		set := rdb.Set(key, val, tm)
//		fmt.Println(set)
//		return nil
//	}
func AddSession(ctx context.Context, key string, val interface{}, mn time.Duration) error {
	set, err := rdb.Set(ctx, key, val, mn).Result()
	if err != nil {
		return err
	}
	fmt.Println(set)
	return nil
}

// GetSession :
func GetSession(ctx context.Context, key string) interface{} {
	value := rdb.Get(ctx, key).Val()
	fmt.Println(value)
	return value
}
