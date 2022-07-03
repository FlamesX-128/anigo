package plugins

import (
	fs "path"
	"strings"

	"github.com/FlamesX-128/anigo/src/utils"
)

func Search(path string) (p []string) {
	for _, file := range utils.ReadDir(path) {
		filename := file.Name()

		if file.Mode().IsDir() {
			Search(fs.Join(path, filename))

			continue
		}

		if strings.HasSuffix(filename, ".so") {
			p = append(p, fs.Join(path, filename))

		}

	}

	return
}
