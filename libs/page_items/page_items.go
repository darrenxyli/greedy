// Package page_items contains parsed result by PageProcesser.
// The result is processed by Pipeline.
package page_items

import (
	"github.com/darrenxyli/greedy/libs/request"
)

// PageItems represents an entity save result parsed by PageProcesser and will be output at last.
type PageItems struct {

	// The req is Request object that contains the parsed result, which saved in PageItems.
	req *request.Request

	// The items is the container of parsed result.
	items map[string]string

	// The skip represents whether send ResultItems to scheduler or not.
	skip bool
}

// NewPageItems returns initialized PageItems object.
func NewPageItems(req *request.Request) *PageItems {
	items := make(map[string]string)
	return &PageItems{req: req, items: items}
}

// GetRequest returns request of PageItems
func (pItem *PageItems) GetRequest() *request.Request {
	return pItem.req
}

// AddItem saves a KV result into PageItems.
func (pItem *PageItems) AddItem(key string, item string) {
	pItem.items[key] = item
}

// GetItem returns value of the key.
func (pItem *PageItems) GetItem(key string) (string, bool) {
	t, ok := pItem.items[key]
	return t, ok
}

// GetAll returns all the KVs result.
func (pItem *PageItems) GetAll() map[string]string {
	return pItem.items
}

// SetSkip set skip true to make pItem page not to be processed by Pipeline.
func (pItem *PageItems) SetSkip(skip bool) *PageItems {
	pItem.skip = skip
	return pItem
}

// GetSkip returns skip label.
func (pItem *PageItems) GetSkip() bool {
	return pItem.skip
}
