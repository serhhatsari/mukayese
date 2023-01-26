package internal

import (
	"bufio"
	"os"
	"strings"
)

func ParseSumFile(f *os.File) (map[string]string, error) {
	datas := make(map[string]string)

	scanner := bufio.NewScanner(f)
	for i := 1; scanner.Scan(); i++ {
		s := strings.Split(scanner.Text(), "@")

		path := s[0]
		checksum := s[1]

		c := strings.Split(checksum, ":")

		hash := c[1]

		datas[path] = hash
	}

	//rd := bufio.NewReader(f)
	//for {
	//	line, err := rd.ReadString(EOL)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//
	//		log.Fatalf("read file line error: %v", err)
	//		return nil, err
	//	}
	//
	//	s := strings.Split(line, "@")
	//
	//	path := s[0]
	//	checksum := s[1]
	//
	//	c := strings.Split(checksum, ":")
	//
	//	hash := c[1]
	//
	//	datas[path] = hash
	//}

	return datas, nil
}
