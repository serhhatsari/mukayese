package internal

import (
	"fmt"
	"strings"
)

const ALGORITHM = "sha256"

const FORMAT = "%s@%s:%s"

type Checksum struct {
	path      string
	algorithm string
	hash      string
}

func NewChecksum(path string, algorithm string, hash string) *Checksum {
	return &Checksum{path: path, algorithm: algorithm, hash: hash}
}

func (c *Checksum) String() string {
	return fmt.Sprintf(FORMAT, c.path, c.algorithm, c.hash)
}

func Format(path *string, checksum *string) string {
	return fmt.Sprintf(FORMAT, *path, ALGORITHM, *checksum)
}

func Parse(str string) *Checksum {
	pathAndChecksum := strings.Split(str, "@")
	algAndHash := strings.Split(pathAndChecksum[1], ":")
	return NewChecksum(pathAndChecksum[0], algAndHash[0], algAndHash[1])
}
