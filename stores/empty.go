package stores

import (
	"github.com/dave/flux"
	"github.com/pubgo/vapper/jsvapper"
)

func NewEmptyStore() *EmptyStore {
	s := &EmptyStore{
	}
	return s
}

type EmptyStore struct {
	app *jsvapper.Vapper
}

func (s *EmptyStore) Init(app *jsvapper.Vapper) {
	s.app = app
}

func (s *EmptyStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	default:
		_ = action
	}
	return true
}
