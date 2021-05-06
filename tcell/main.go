// Copyright 2020 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/mattn/go-runewidth"
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func displayHelloWorld(s tcell.Screen, dx int) {
	w, h := s.Size()
	s.Clear()
	emitStr(s, w/2-7+dx, h/2, tcell.StyleDefault, "Hello, World!")
	emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
	s.Show()
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	displayHelloWorld(s, 0)

	dx := 0
	mleft := true

	go func() {
		for {
			<-time.After(120 * time.Millisecond)
			if dx < -5 {
				mleft = false
			}
			if dx > 5 {
				mleft = true
			}
			if mleft {
				dx--
			} else {
				dx++
			}
			displayHelloWorld(s, dx)
		}
	}()
	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			displayHelloWorld(s, dx)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Clear()
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
