package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Split(src string, dest string, size int) error {
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0444)

	if err != nil {
		return err
	}

	chunks := 0
	block := make([]byte, size)
	destDir := filepath.Join(dest, filepath.Base(src))

	if err = os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	for {
		if n, err := srcFile.Read(block); n == 0 {

			if err == io.EOF {
				break
			} else if err == nil {
				continue
			} else {
				return err
			}
		} else if n < size {
			block = block[:n]
		}

		if err = ioutil.WriteFile(filepath.Join(destDir, strconv.Itoa(chunks)), block, 0666); err != nil {
			return err
		}

		chunks++
	}

	return srcFile.Close()
}
