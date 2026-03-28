package internal 

type Buffer struct {
	Data [][]rune
	Name string
	Cursor Vec2i
}

type List struct {
	Buffers []Buffer
	Index int
}

func Index_In_Bounds(x int, length int) bool {
	if(x < length && x >= 0){
		return true
	} else {
		return false
	}
}

func Make_List() List {
	return List { Buffers: []Buffer{}, Index: 0 }
}

func Make_Buffer(_Data [][]rune, _Name string) Buffer {
	return Buffer { Data: _Data, Name: _Name, Cursor: Make_Cursor() }
}

func Buffer_Delete_Rune(list *List){
	if !Index_In_Bounds(list.Index, len(list.Buffers)){
		return
	}

	if !Cursor_In_Bounds(list) {
		return
	}
	
//	buffer := list.Buffers[list.Index].Data
//	cursor := list.Buffers[list.Index].Cursor
}

func Buffer_Add_Rune(codepoint rune, list *List){
	if !Index_In_Bounds(list.Index, len(list.Buffers)){
		return
	}

	if !Cursor_In_Bounds(list) {
		return
	}

	//list.Buffers[list.Index].Data[cursor.Y] = append(list.Buffers[list.Index].Data[cursor.Y][:cursor.X], codepoint)
}

func Append_Buffer(Data [][]rune, Name string, list *List) {
	// Linear search over Buffers to find if a buffer with the same Name already exists
	for _, entry := range list.Buffers {
		if entry.Name == Name {
			return
		}
	}
	list.Buffers = append(list.Buffers, Make_Buffer(Data, Name))
}

func Change_Buffer(inc int, list *List){
	tmp := list.Index + inc
	if(!Index_In_Bounds(tmp, len(list.Buffers))){
		return
	}
	list.Index = tmp
}

func Delete_Current_Buffer(list *List){
	tmp := append(list.Buffers[:list.Index], list.Buffers[list.Index + 1:]...)
	if len(tmp) > 0 {
		list.Index = len(tmp) - 1
	} else {
		list.Index = 0
	}
	list.Buffers = tmp
}

func Go_To_Buffer(Name string, list *List) {
	for i, entry := range list.Buffers {
		if Name == entry.Name {
			list.Index = i
			return
		}
	}
}
