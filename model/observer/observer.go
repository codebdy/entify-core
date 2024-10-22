package observer

import (
	"context"
	"sync"

	"github.com/codebdy/entify-core/shared"
)

type ModelObserver interface {
	Key() string
	ObjectPosted(object map[string]interface{}, entityName string, ctx context.Context)
	ObjectMultiPosted(objects []map[string]interface{}, entityName string, ctx context.Context)
	ObjectDeleted(object map[string]interface{}, entityName string, ctx context.Context)
	ObjectMultiDeleted(objects []map[string]interface{}, entityName string, ctx context.Context)
}

var ModelObservers sync.Map

func AddObserver(obsr ModelObserver) {
	ModelObservers.Store(obsr.Key(), obsr)
}

func RemoveObserver(key string) {
	ModelObservers.Delete(key)
}

func EmitObjectPosted(object map[string]interface{}, entityName string, ctx context.Context) {
	//newCtx := context.WithValue(context.Background(), consts.CONTEXT_VALUES, contexts.Values(ctx))
	newCtx := context.WithValue(context.Background(), shared.LOADERS, ctx.Value(shared.LOADERS))
	go func() {
		ModelObservers.Range(func(key interface{}, value interface{}) bool {
			value.(ModelObserver).ObjectPosted(object, entityName, newCtx)
			return true
		})
	}()
}

func EmitObjectMultiPosted(objects []map[string]interface{}, entityName string, ctx context.Context) {
	go func() {
		ModelObservers.Range(func(key interface{}, value interface{}) bool {
			value.(ModelObserver).ObjectMultiPosted(objects, entityName, ctx)
			return true
		})
	}()
}

func EmitObjectDeleted(object map[string]interface{}, entityName string, ctx context.Context) {
	go func() {
		ModelObservers.Range(func(key interface{}, value interface{}) bool {
			value.(ModelObserver).ObjectDeleted(object, entityName, ctx)
			return true
		})
	}()
}

func EmitObjectMultiDeleted(objects []map[string]interface{}, entityName string, ctx context.Context) {
	go func() {
		ModelObservers.Range(func(key interface{}, value interface{}) bool {
			value.(ModelObserver).ObjectMultiDeleted(objects, entityName, ctx)
			return true
		})
	}()
}
