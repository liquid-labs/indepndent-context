package indepndent_context

import (
	"context"
)

type copiedContext struct {
	context.Context
	values map[interface{}]interface{}
}

func NewIndependentContext(parent context.Context) *copiedContext {
	newCtx := &copiedContext{
		Context: context.Background(),
		values:  make(map[interface{}]interface{}),
	}

	if parent != nil {
		parentValues := getAllContextValues(parent)
		for k, v := range parentValues {
			newCtx.values[k] = v
		}
	}
	return newCtx
}

func (c *copiedContext) Value(key interface{}) interface{} {
	return c.values[key]
}

func getAllContextValues(ctx context.Context) map[interface{}]interface{} {
	values := make(map[interface{}]interface{})
	collectValues(ctx, values)
	return values
}

func collectValues(ctx context.Context, values map[interface{}]interface{}) {
	if ctx == nil {
		return
	}
	if val := ctx.Value(ctx); val != nil {
		values[ctx] = val
	}

	collectValues(getParent(ctx), values)
}

func getParent(ctx context.Context) context.Context {
	if parent, ok := ctx.(interface{ Parent() context.Context }); ok {
		return parent.Parent()
	}
	return nil
}
