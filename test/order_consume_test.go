package test

import (
	"crypto/md5"
	"fmt"
	"sync"
	"testing"
)

type ConsumeData struct {
	AppId string
	Sort  int
}

func TestOrderConsume(t *testing.T) {
	data := []*ConsumeData{
		{
			"A", 5,
		}, {
			"C", 4,
		}, {
			"B", 6,
		}, {
			"A", 4,
		}, {
			"D", 5,
		}, {
			"D", 4,
		}, {
			"B", 5,
		}, {
			"A", 3,
		}, {
			"C", 3,
		}, {
			"C", 2,
		}, {
			"A", 2,
		}, {
			"B", 4,
		}, {
			"B", 3,
		}, {
			"D", 3,
		}, {
			"A", 1,
		}, {
			"D", 2,
		}, {
			"C", 1,
		}, {
			"B", 2,
		}, {
			"D", 1,
		}, {
			"B", 1,
		},
	}
	ConsumeV2(data)
}

// 协程保证不了顺序
func Consume(data []*ConsumeData) {
	var wg sync.WaitGroup
	wg.Add(len(data))
	for _, v := range data {
		tempData := v
		go func() {
			defer wg.Done()
			fmt.Printf("该店铺:%s + %d \n", tempData.AppId, tempData.Sort)
		}()
	}
	wg.Wait()
}

// 更改一下
func ConsumeV2(data []*ConsumeData) {
	var wg sync.WaitGroup
	hashMap := make(map[string][]*ConsumeData)
	for _, v := range data {
		tempString := v.AppId
		hash := md5.Sum([]byte(tempString))
		hashString := string(hash[:])
		if _, ok := hashMap[hashString]; ok {
			hashMap[hashString] = append(hashMap[hashString], v)
		} else {
			var TempConsumeSlice []*ConsumeData
			TempConsumeSlice = append(TempConsumeSlice, v)
			hashMap[hashString] = TempConsumeSlice
		}
	}
	wg.Add(len(hashMap))
	for _, m := range hashMap {
		termMap := m
		go func() {
			defer wg.Done()
			for _, s := range termMap {
				fmt.Printf("该店铺:%s + %d \n", s.AppId, s.Sort)
			}
		}()
	}
	wg.Wait()
}
