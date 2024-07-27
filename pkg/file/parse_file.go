package file

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/sandricoprovo/fran/pkg/util"
)

func ParseFile(f string, lineHandler func(string) string) {
	// Open source file
	sf, err := os.Open(f)
	util.PanicCheck(err)

	defer sf.Close()

	// Create a temp file to later overwrite source file
	tfp := f + ".tmp"
	tf, err := os.Create(tfp)
	util.PanicCheck(err)

	defer tf.Close()

	scanner := bufio.NewScanner(sf)
	tfw := bufio.NewWriter(tf)

	for scanner.Scan() {
		var nl string

		l := scanner.Text()
		nl = lineHandler(l)

		// Write the updated line to the temporary file
		_, err = tfw.WriteString(nl + "\n")
		util.PanicCheck(err)
	}

	scannerErr := scanner.Err()
	util.PanicCheck(scannerErr)

	tfw.Flush()

	err = os.Rename(tfp, f)
	util.PanicCheck(err)
}

func ParseFileRecursive(exts []string, parser func(string) string) {
	rd := "."

	err := filepath.Walk(rd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			for _, e := range exts {
				if ext == e {
					ParseFile(path, parser)
					break
				}
			}
		}

		return nil
	})

	util.PanicCheck(err)
}
