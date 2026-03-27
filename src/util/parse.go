package util

func ParseBytes(buffer []byte) [][]rune {
	const delimiter = '\n'
	const skip = '\r'

	line_buffer := make([][]rune, 0)
	byte_str := make([]byte, 0)

	for _, byte := range buffer {
		switch byte {
			default: {
				byte_str = append(byte_str, byte)
			}

			case delimiter: {
				line_buffer = append(line_buffer, []rune(string(byte_str)))
				byte_str = nil
			}

			case skip: {
				continue
			}
			}
	}

	if len(byte_str) > 0 {
		line_buffer = append(line_buffer, []rune(string(byte_str)))
	}

	return line_buffer
}
