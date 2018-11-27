/**
 * Copyright 2018 godog Author. All rights reserved.
 * Author: Chuck1024
 */

package main_test

import (
	"github.com/chuck1024/godog"
	"github.com/chuck1024/godog/dao/cache"
	"testing"
)

func TestRedis(t *testing.T) {
	URL, _ := godog.AppConfig.String("redis")
	cache.Init(URL)
	key := "key"

	err := cache.Set(key, "value")
	if err != nil {
		t.Logf("redis set occur error:%s", err)
		return
	}

	t.Logf("set success:%s", key)

	value, err := cache.Get(key)
	if err != nil {
		t.Logf("redis get occur error: %s", err)
		return
	}
	t.Logf("get value: %s", value)
}
