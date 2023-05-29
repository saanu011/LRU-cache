package main

import (
	"fmt"

	"container/list"
)

type Cache struct {
	list    *list.List
	maxSize int
}

func main() {
	cache := NewCache()
	str := []string{"A", "Z", "X", "B", "D", "A", "C", "A", "B"}
	for _, e := range str {
		cache.LRU(e)
		cache.printAll()
	}
}

func NewCache() Cache {
	return Cache{
		list:    list.New(),
		maxSize: 5,
	}
}

func existsInList(ls *list.List, a *list.Element) (bool, *list.Element) {
	temp := ls.Front()
	for temp != nil {
		if a.Value == temp.Value {
			return true, temp
		}
		temp = temp.Next()
	}

	return false, nil
}

func (c *Cache) removeLast() {
	temp := c.list.Front()
	a := temp
	for temp != nil {
		a = temp
		temp = temp.Next()
	}
	c.list.Remove(a)
}

func (c *Cache) LRU(a string) {
	if ok, e := existsInList(c.list, &list.Element{Value: a}); ok {
		c.list.Remove(e)
		c.list.PushFront(a)
	} else {
		if c.list.Len() < c.maxSize {
			c.list.PushFront(a)
		} else {
			c.removeLast()
			c.list.PushFront(a)
		}
	}
}

func (c *Cache) printAll() {
	temp := c.list.Front()

	for temp != nil {
		fmt.Printf(`"%v" `, temp.Value)
		temp = temp.Next()
	}

	fmt.Println()
}
