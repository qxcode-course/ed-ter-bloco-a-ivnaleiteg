package main

import (
	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() {
		e.cursor = e.cursor.Prev()
	} else if e.line != e.lines.Front() {
		e.line = e.line.Prev()
		e.cursor = e.line.Value.End()
	}
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
	} else if e.line.Next() != e.lines.End() {
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
	}
}

func (e *Editor) KeyEnter() {
	line := e.line.Value
	newLine := NewList[rune]()
	for it := e.cursor; it != line.End(); {
		next := it.Next()
		newLine.Insert(newLine.End(), it.Value)
		line.Erase(it)
		it = next
	}
	newNode := e.lines.Insert(e.line.Next(), newLine)
	e.line = newNode
	e.cursor = newLine.Front()
}

func (e *Editor) getColumn() int {
	i := 0
	for it := e.line.Value.Front(); it != e.cursor && it != e.line.Value.End(); it = it.Next() {
		i++
	}
	return i
}

func (e *Editor) getCharAt(line *List[rune], col int) *Node[rune] {
	it := line.Front()
	for i := 0; it != line.End() && i < col; i++ {
		it = it.Next()
	}
	return it
}

func (e *Editor) KeyUp() {
	if e.line.Prev() != e.lines.End() {
		col := e.getColumn()
		e.line = e.line.Prev()
		e.cursor = e.getCharAt(e.line.Value, col)
	}
}

func (e *Editor) KeyDown() {
	if e.line.Next() != e.lines.End() {
		col := e.getColumn()
		e.line = e.line.Next()
		e.cursor = e.getCharAt(e.line.Value, col)
	}
}

func (e *Editor) KeyBackspace() {
	line := e.line.Value
	if e.cursor != line.Front() {
		e.cursor = line.Erase(e.cursor.Prev())
	} else if e.line.Prev() != e.lines.End() {
		prevNode := e.line.Prev()
		prevLine := prevNode.Value
		pos := prevLine.End()
		for it := line.Front(); it != line.End(); {
			next := it.Next()
			prevLine.Insert(prevLine.End(), it.Value)
			it = next
		}
		e.lines.Erase(e.line)
		e.line = prevNode
		e.cursor = pos
	}
}

func (e *Editor) KeyDelete() {
	line := e.line.Value
	if e.cursor != line.End() {
		e.cursor = line.Erase(e.cursor)
	} else if e.line.Next() != e.lines.End() {
		nextNode := e.line.Next()
		nextLine := nextNode.Value
		for it := nextLine.Front(); it != nextLine.End(); {
			next := it.Next()
			line.Insert(line.End(), it.Value)
			it = next
		}
		e.lines.Erase(nextNode)
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	screen, _ := tcell.NewScreen()
	screen.Init()
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.End()
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x, y := 0, 0
	cursorX, cursorY := -1, -1
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = ' '
			}
			style := e.style
			if char == e.cursor {
				cursorX, cursorY = x, y
				style = style.Reverse(true)
			}
			e.screen.SetContent(x, y, data, nil, style)
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	if cursorX != -1 {
		e.screen.ShowCursor(cursorX, cursorY)
	}
	e.screen.Show()
}

func main() {
	editor := NewEditor()
	defer editor.screen.Fini()
	editor.Draw()
	for {
		ev := editor.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				editor.KeyEnter()
			case tcell.KeyLeft:
				editor.KeyLeft()
			case tcell.KeyRight:
				editor.KeyRight()
			case tcell.KeyUp:
				editor.KeyUp()
			case tcell.KeyDown:
				editor.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				editor.KeyBackspace()
			case tcell.KeyDelete:
				editor.KeyDelete()
			default:
				if ev.Rune() != 0 {
					editor.InsertChar(ev.Rune())
				}
			}
			editor.Draw()
		case *tcell.EventResize:
			editor.screen.Sync()
			editor.Draw()
		}
	}
}
