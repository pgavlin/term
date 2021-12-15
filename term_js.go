// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js
// +build js

package term

import (
	"syscall/js"
)

type state struct {
	obj js.Value
}

var terminal = js.Global().Get("terminal")

func isTerminal(fd int) bool {
	rv := terminal.Call("isTerminal", fd)
	return rv.Bool()
}

func makeRaw(fd int) (s *State, err error) {
	defer func() { catch(recover(), &err) }()

	old := terminal.Call("makeRaw", fd)
	return &State{state{old}}, nil
}

func getState(fd int) (s *State, err error) {
	defer func() { catch(recover(), &err) }()

	old := terminal.Call("getState", fd)
	return &State{state{old}}, nil
}

func restore(fd int, state *State) (err error) {
	defer func() { catch(recover(), &err) }()

	terminal.Call("restore", fd, state.state.obj)
	return nil
}

func getSize(fd int) (width, height int, err error) {
	defer func() { catch(recover(), &err) }()

	sz := terminal.Call("getSize", fd)
	return sz.Get("width").Int(), sz.Get("height").Int(), nil
}

func readPassword(fd int) (p []byte, err error) {
	defer func() { catch(recover(), &err) }()

	passwd := terminal.Call("readPassword", fd)
	return []byte(passwd.String()), nil
}

func catch(x interface{}, err *error) {
	if x != nil {
		e, ok := x.(error)
		if !ok {
			panic(x)
		}
		*err = e
	}
}
