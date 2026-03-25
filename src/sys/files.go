package sys

import (
	"os"
)

func ReadFile(path String)  {
	File *f = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer f.Close()


}

func WriteFile(path String)
