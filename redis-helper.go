package common

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

func PrepareRedisClient(address string, password string, dbNumber int, protocol int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // "" : no password set
		DB:       dbNumber, // 0 : use default DB
		Protocol: protocol, // specify 2 for RESP 2 or 3 for RESP 3
	})
}

func GetAllAnchorsFromRedis() map[string]Anchor {
	var anchorMap = make(map[string]Anchor)

	iter := redisClient.Scan(ctx, 0, "anchor:*", 0).Iterator()
	for iter.Next(ctx) {
		aKey := iter.Val()
		a, err := ReadAnchor(aKey)
		if err == nil {
			anchorMap[aKey] = a
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return anchorMap
}

func ReadAnchor(anchorKey string) (anchor Anchor, err error) {
	jsonVal, err := redisClient.JSONGet(ctx, anchorKey, "ID", "Location", "Range").Result()
	if err == redis.Nil || err != nil {
		return
	}

	err = json.Unmarshal([]byte(jsonVal), &anchor)
	return
}

func GetAllEntitiesFromRedis() map[string]Entity {
	var entityMap = make(map[string]Entity)

	iter := redisClient.Scan(ctx, 0, "entity:*", 0).Iterator()
	for iter.Next(ctx) {
		eKey := iter.Val()
		e, err := ReadEntity(eKey)
		if err == nil {
			entityMap[eKey] = e
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return entityMap
}

func ReadEntity(entityKey string) (entity Entity, err error) {
	jsonVal, err := redisClient.JSONGet(ctx, entityKey, "ID", "Height", "TagID", "Type", "MaxSpeed").Result()
	if err == redis.Nil || err != nil {
		return
	}

	err = json.Unmarshal([]byte(jsonVal), &entity)
	if err != nil {
		return
	}
	return
}

func GetAllTagsFromRedis() map[string]int {
	var tagMap = make(map[string]int)

	iter := redisClient.Scan(ctx, 0, "tag:*", 0).Iterator()
	for iter.Next(ctx) {
		tKey := iter.Val()
		t, err := ReadTag(tKey)
		if err == nil {
			tagMap[tKey] = t
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return tagMap
}

func ReadTag(tagKey string) (entityId int, err error) {
	entityIdStr, err := redisClient.Get(ctx, tagKey).Result()
	if err == redis.Nil || err != nil {
		return
	}

	entityId, _ = strconv.Atoi(entityIdStr)
	return
}
