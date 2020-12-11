// Package format provides MCPE level formats written for lav7.
package format

import (
	"reflect"
	"strings"

	"github.com/L7-MCPE/lav7/types"
)

var levelProviders = map[string]Provider{}

// Provider is a interface for level formats.
type Provider interface {
	Init(string)                          // Level name: usually used for file directories
	Loadable(int32, int32) (string, bool) // Path: path to file, Ok: if the chunk is saved on the file
	LoadChunk(int32, int32, string) (*types.Chunk, error)
	WriteChunk(int32, int32, *types.Chunk) error
	SaveAll(map[[2]int32]*types.Chunk) error
}

// RegisterProvider adds level format provider for server.
func RegisterProvider(provider Provider) {
	typname := reflect.TypeOf(provider)
	typsl := strings.Split(typname.String(), ".")
	name := strings.ToLower(typsl[len(typsl)-1])
	if _, ok := levelProviders[name]; !ok {
		levelProviders[name] = provider
	}
}

// GetProvider finds the provider with given name.
// If it doesn't present, returns nil.
func GetProvider(name string) Provider {
	if pv, ok := levelProviders[name]; ok {
		return pv
	}
	return nil
}
