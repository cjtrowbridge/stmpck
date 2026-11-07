//go:build js && wasm

package storage

import (
	"errors"
	"fmt"
	"syscall/js"
)

func DefaultRoot() string { return "stmpck" }
func (s *Store) read(name string) (data []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("browser storage: %v", e)
		}
	}()
	v := js.Global().Get("localStorage").Call("getItem", s.Root+"/"+name)
	if v.IsNull() {
		return nil, errors.New("no saved games")
	}
	return []byte(v.String()), nil
}
func (s *Store) write(name string, data []byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("browser storage: %v", e)
		}
	}()
	js.Global().Get("localStorage").Call("setItem", s.Root+"/"+name, string(data))
	return nil
}
