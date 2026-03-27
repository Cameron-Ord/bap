package internal

import ()

type Cursor struct {
	x int
	y int;
}

func Cursor_In_Bounds(cursor *Cursor, list *List) bool {
	if cursor.y >= len(list.Buffers[list.Index].Data) || cursor.y < 0 {
		return false
	}
	if cursor.x > len(list.Buffers[list.Index].Data[cursor.y]) || cursor.x < 0 {
		return false
	}
	return true
}
