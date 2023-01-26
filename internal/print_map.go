package internal

import (
	"fmt"
)

func PrintMap(title string, datas *map[string]string) {
	if 0 != len(title) {
		fmt.Printf("%s:\n", title)
	}

	if len(*datas) == 0 {
		fmt.Println("no any files")
		return
	}

	sortedDatas := Sort(datas)

	for _, data := range *sortedDatas {
		fmt.Println(data)
	}
}
