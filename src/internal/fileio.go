package internal

import ( 
	"fmt"
	"log"
	"os"
	"io"
)

func WriteFile() {

}

func ReadFile(file *os.File) []byte {
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return data
}

func OpenFile(path string) *os.File {
	fmt.Printf("Attempting to open: %s\n", path)
	file, err := os.OpenFile(path, os.O_RDONLY | os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return file
}

func CloseFile(file *os.File) {
	if file != nil {
		file.Close()
	}
}


