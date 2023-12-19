package rateLimiter

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
	limit int
}

func NewCounter(limit int) *Counter {
	return &Counter{
		limit: limit,
	}
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	return c.count
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *Counter) IsLimitExceeded() bool {
	return c.GetCount() > c.limit
}
