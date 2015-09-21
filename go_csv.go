package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.xls")
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.Write([]string{"编号", "姓名", "年龄"})
	w.Write([]string{"1", "张三", "23"})
	w.Write([]string{"2", "李四", "24"})
	w.Write([]string{"3", "王五", "25"})
	w.Flush()
}
