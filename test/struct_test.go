package test

import (
	"sync"
	"testing"
)

type Person struct {
	Id   int
	name string
}

func TestStruct(t *testing.T) {
	s := []string{"bw", "jw", "percy", "ma"}
	z := []int{1, 3, 4, 2}
	ch := make(chan Person, len(s))
	m := make(map[string]Person)
	var wg sync.WaitGroup
	wg.Add(len(s))
	for i := 0; i < 4; i++ {
		tempK := i
		go func() {
			defer wg.Done()
			temp := Person{
				z[tempK], s[tempK],
			}
			ch <- temp
		}()
	}
	wg.Wait()
	close(ch)
	for c := range ch {
		m[c.name] = c
	}
	t.Log(m)
}
