package redisdb

import (
	"context"
	"encoding/json"
	"nuryanto2121/cukur_in_capster/pkg/setting"

	"github.com/mitchellh/mapstructure"
)

// Register :
type Register struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	ButtonLink   string `json:"button_link"`
	AktivasiLink string `json:"aktivasi_link"`
}

// StoreRegister :
func StoreRegister(ctx context.Context, data interface{}) error {
	var Register Register

	err := mapstructure.Decode(data, &Register)
	if err != nil {
		return err
	}

	bRegister, err := json.Marshal(Register)
	if err != nil {
		return err
	}

	mRegister := map[string]interface{}{
		"email_type": "register",
		"data":       string(bRegister),
	}

	dRegister, err := json.Marshal(mRegister)
	if err != nil {
		return err
	}

	_, err = rdb.SAdd(ctx, setting.FileConfigSetting.RedisDBSetting.Key, string(dRegister)).Result()
	if err != nil {
		return err
	}

	return nil
}
