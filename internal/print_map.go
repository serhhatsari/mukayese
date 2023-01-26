package internal

import (
	"fmt"
	"sort"
)

func PrintMap(title string, datas *map[string]string) {
	if 0 != len(title) {
		fmt.Printf("%s:\n", title)
	}

	if len(*datas) == 0 {
		fmt.Println("no any files")
		return
	}

	var sortedDatas []string
	for key, val := range *datas {
		s := Format(&key, &val)
		sortedDatas = append(sortedDatas, s)
	}
	sort.Strings(sortedDatas)

	for _, data := range sortedDatas {
		fmt.Println(data)
	}
}
