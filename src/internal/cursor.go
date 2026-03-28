package internal

import ()

type Vec2i struct {
	X int
	Y int;
}

func Make_Cursor() Vec2i {
	return Vec2i { X: 0, Y: 0 }
}

func X_Bounds(x int, length int) bool {
	if x >= length || x < 0 {
		return false
	} else {
		return true
	}
}

func Y_Bounds(y int, length int) bool {
	if y > length || y < 0 {
		return false
	} else {
		return true
	}
}

func Cursor_In_Bounds(list *List) bool {
	curs := list.Buffers[list.Index].Cursor
	data := list.Buffers[list.Index].Data
	
	if curs.Y > len(data) || curs.Y < 0 || len(data) <= 0 {
		return false
	}

	if curs.X > len(data[curs.Y]) || curs.X < 0 {
		return false
	}
	return true
}

func Cursor_Add_Y(inc int, buf *Buffer){
	curs := &buf.Cursor
	data := buf.Data
	if !Y_Bounds(curs.Y + inc, len(data)) {
		return
	}
	curs.Y = curs.Y + inc
	if !X_Bounds(curs.X, len(data[curs.Y])) && len(data[curs.Y]) > 0 {
		curs.X = len(data[curs.Y]) - 1
	} else if len(data[curs.Y]) <= 0 {
		curs.X = 0
	}
}

func Cursor_Add_X(inc int, buf *Buffer){
	curs := &buf.Cursor
	data := buf.Data
	if !X_Bounds(curs.X + inc, len(data[curs.Y])) {
		return
	}
	curs.X = curs.X + inc
}
