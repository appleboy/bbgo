// Code generated by "callbackgen -type MarketDataStore"; DO NOT EDIT.

package bbgo

import (
	"github.com/c9s/bbgo/types"
)

func (store *MarketDataStore) OnUpdate(cb KLineCallback) {
	store.updateCallbacks = append(store.updateCallbacks, cb)
}

func (store *MarketDataStore) EmitUpdate(kline types.KLine) {
	for _, cb := range store.updateCallbacks {
		cb(kline)
	}
}
