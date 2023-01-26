package internal

import "sort"

func Sort(datas *map[string]string) *[]string {
	var sortedDatas []string
	for key, val := range *datas {
		s := Format(&key, &val)
		sortedDatas = append(sortedDatas, s)
	}
	sort.Strings(sortedDatas)
	return &sortedDatas
}
