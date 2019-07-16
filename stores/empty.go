package stores

import (
	"github.com/dave/flux"
	"github.com/pubgo/vapper/vapper"
)

func NewEmptyStore() *EmptyStore {
	s := &EmptyStore{
	}
	return s
}

type EmptyStore struct {
	app *vapper.Vapper
}

func (s *EmptyStore) Init(app *vapper.Vapper) {
	s.app = app
}

func (s *EmptyStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	default:
		_ = action
	}
	return true
}
