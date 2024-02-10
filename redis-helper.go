package common

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RtlsRedis struct {
	client  *redis.Client
	options *redis.Options
	ctx     context.Context
}

func PrepareRedisClient(address string, password string, dbNumber int, protocol int) *RtlsRedis {
	rtlsRedis := RtlsRedis{
		options: &redis.Options{
			Addr:     address,
			Password: password, // "" : no password set
			DB:       dbNumber, // 0 : use default DB
			Protocol: protocol, // specify 2 for RESP 2 or 3 for RESP 3
		},
	}
	rtlsRedis.client = redis.NewClient(rtlsRedis.options)
	rtlsRedis.ctx = context.Background()
	return &rtlsRedis
}

func (rr *RtlsRedis) GetAllAnchorsFromRedis() map[string]Anchor {
	var anchorMap = make(map[string]Anchor)

	iter := rr.client.Scan(rr.ctx, 0, "anchor:*", 0).Iterator()
	for iter.Next(rr.ctx) {
		aKey := iter.Val()
		a, err := rr.ReadAnchor(aKey)
		if err == nil {
			anchorMap[aKey] = a
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return anchorMap
}

func (rr *RtlsRedis) ReadAnchor(anchorKey string) (anchor Anchor, err error) {
	jsonVal, err := rr.client.JSONGet(rr.ctx, anchorKey, "ID", "Location", "Range", "Sudoku").Result()
	if err == redis.Nil || err != nil {
		return
	}

	err = json.Unmarshal([]byte(jsonVal), &anchor)
	return
}

func (rr *RtlsRedis) GetAllEntitiesFromRedis() map[string]Entity {
	var entityMap = make(map[string]Entity)

	iter := rr.client.Scan(rr.ctx, 0, "entity:*", 0).Iterator()
	for iter.Next(rr.ctx) {
		eKey := iter.Val()
		e, err := rr.ReadEntity(eKey)
		if err == nil {
			entityMap[eKey] = e
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return entityMap
}

func (rr *RtlsRedis) ReadEntity(entityKey string) (entity Entity, err error) {
	jsonVal, err := rr.client.JSONGet(rr.ctx, entityKey, "ID", "Height", "TagID", "Type", "MaxSpeed").Result()
	if err == redis.Nil || err != nil {
		return
	}

	err = json.Unmarshal([]byte(jsonVal), &entity)
	if err != nil {
		return
	}
	return
}

func (rr *RtlsRedis) GetAllTagsFromRedis() map[string]int {
	var tagMap = make(map[string]int)

	iter := rr.client.Scan(rr.ctx, 0, "tag:*", 0).Iterator()
	for iter.Next(rr.ctx) {
		tKey := iter.Val()
		t, err := rr.ReadTag(tKey)
		if err == nil {
			tagMap[tKey] = t
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return tagMap
}

func (rr *RtlsRedis) ReadTag(tagKey string) (entityId int, err error) {
	entityIdStr, err := rr.client.Get(rr.ctx, tagKey).Result()
	if err == redis.Nil || err != nil {
		return
	}

	entityId, _ = strconv.Atoi(entityIdStr)
	return
}

func (rr *RtlsRedis) WriteMockDataToRedis(anchorCount int, tagCount int) {
	for i := 0; i < anchorCount; i++ {
		a := Anchor{
			ID: 1001 + i,
			Location: Location{
				FloorID: 1,
				Point: Point{
					X: 100 * (i % 4),
					Y: 100 * int(i/4),
					Z: 250,
				},
			},
			Range:  1000,
			Sudoku: i % 10,
		}
		aKey := fmt.Sprintf("anchor:%d", a.ID)
		jsonValue, err := json.Marshal(a)
		if err != nil {
			return
		}
		rr.client.JSONSet(rr.ctx, aKey, "$", jsonValue)
		time.Sleep(1 * time.Millisecond)
	}

	for i := 0; i < tagCount; i++ {
		tagID := 10001 + i
		entityID := 1 + i
		tagKey := fmt.Sprintf("tag:%d", tagID)
		rr.client.Set(rr.ctx, tagKey, entityID, 0)
		time.Sleep(1 * time.Millisecond)
	}

	for i := 0; i < tagCount; i++ {
		e := Entity{
			ID:       1 + i,
			TagID:    10001 + i,
			Type:     2,
			Height:   150,
			MaxSpeed: 300,
		}
		eKey := fmt.Sprintf("entity:%d", e.ID)
		jsonValue, err := json.Marshal(e)
		if err != nil {
			return
		}
		rr.client.JSONSet(rr.ctx, eKey, "$", jsonValue)
		time.Sleep(10 * time.Millisecond)
	}
}
