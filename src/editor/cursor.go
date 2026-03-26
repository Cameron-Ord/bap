package text_manipulation

import ("log")

type Vec2i struct {
	x int
	y int;
}

type Editor struct {
	cursor Vec2i
	line_buffer [][]rune

}

func Check_Cursor_Bounds(cursor Vec2i, line_buffer [][]rune) bool {
	x := cursor.x
	y := cursor.y

	if y > len(line_buffer) || y < 0 {
		return false
	}

	line := line_buffer[y]
	if x > len(line) || x < 0 {
		return false
	}

	return true
}

func Delete_At_Position(editor Editor){
	if !Check_Cursor_Bounds(editor.cursor, editor.line_buffer) {
		return
	}
	editor.line_buffer[editor.cursor.y] = append(editor.line_buffer[editor.cursor.y][:editor.cursor.x], editor.line_buffer[editor.cursor.y][editor.cursor.x+1:]...)
}

func Move_Cursor_X(editor *Editor, inc int){
	if editor.cursor.y >= len(editor.line_buffer) || editor.cursor.y < 0 {
		log.Fatalf("Invalid position of Y")
	}

	tmp := editor.cursor.x + inc
	if tmp >= len(editor.line_buffer[editor.cursor.y]) {
		tmp = len(editor.line_buffer) - 1
	}

	if tmp < 0 {
		tmp = 0
	}

	editor.cursor.x = tmp
}

func Move_Cursor_Y(editor *Editor, inc int){
	tmp := editor.cursor.y + inc
	if tmp >= len(editor.line_buffer)  {
		tmp = len(editor.line_buffer) - 1
	}

	if tmp < 0 {
		tmp = 0
	}

	editor.cursor.y = tmp
}
