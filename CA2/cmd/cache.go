package main

import (
	"sync"
)

// Cache struct using sync.Map
type Cache struct {
	store sync.Map
}

// Set a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.store.Store(key, value)
}

// Get a value by key from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	val, ok := c.store.Load(key)
	return val, ok
}

// Delete a value by key from the cache
func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}

// func accessCache(c *Cache, id int) {
//     key := fmt.Sprintf("key%d", id)
//     value := rand.Intn(100)

//     // Set a value in the cache
//     c.Set(key, value)
//     fmt.Printf("Goroutine %d set %s to %d\n", id, key, value)

//     // Get a value from the cache
//     if val, ok := c.Get(key); ok {
//         fmt.Printf("Goroutine %d got %s: %d\n", id, key, val)
//     }

//     // Sleep to simulate work
//     time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))

//     // Delete the key
//     c.Delete(key)
//     fmt.Printf("Goroutine %d deleted %s\n", id, key)
// }

// func main() {
//     c := Cache{}

//     // Launch multiple goroutines to access the cache
//     for i := 0; i < 10; i++ {
//         go accessCache(&c, i)
//     }

//     // Wait for all goroutines to finish
//     time.Sleep(5 * time.Second)
// }
