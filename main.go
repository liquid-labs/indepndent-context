package indepndent_context

import (
	"context"
	"time"
)

type copiedContext struct {
	self   context.Context
	values context.Context
}

func NewDerivedContext(source context.Context) *copiedContext {
	newCtx := &copiedContext{
		self:   context.Background(),
		values: source,
	}

	return newCtx
}

func (c *copiedContext) Deadline() (deadline time.Time, ok bool) {
	return c.self.Deadline()
}

func (c *copiedContext) Done() <-chan struct{} {
	return c.self.Done()
}

func (c *copiedContext) Err() error {
	return c.self.Err()
}

func (c *copiedContext) Value(key interface{}) interface{} {
	return c.values.Value(key)
}
