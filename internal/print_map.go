package internal

import "fmt"

func PrintMap(title string, datas *map[string]string) {
	fmt.Printf("%s:\n", title)
	if len(*datas) == 0 {
		fmt.Println("no any files")
	} else {
		for key, val := range *datas {
			fmt.Printf("%s@sha256:%s\n", key, val)
		}
	}

}
