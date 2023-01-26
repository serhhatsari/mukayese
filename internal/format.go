package internal

import (
	"fmt"
	"strings"
)

const format = "%s"

type Type struct {
	path      string
	algorithm string
	checksum  string
}

func Format(path *string, algorithm *string, checksum *string) string {
	return fmt.Sprintf("%s@%s:%s", *path, *algorithm, *checksum)
}

func Parse(str *string) Type {
	s := strings.Split(*str, "@")
	return Type{
		path:      s[0],
		algorithm: s[1],
		checksum:  s[2],
	}
}
