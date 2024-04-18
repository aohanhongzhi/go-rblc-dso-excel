package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {

	OperatingManualV2, err := ListDir("/home/eric/Downloads/2024", "xlsx")
	if err != nil {
		panic(err)
	}

	for _, op := range OperatingManualV2 {
		f, err := excelize.OpenFile(op)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		// 获取工作表中指定单元格的值
		sheet := "Operating manual"
		cell, err := f.GetCellValue(sheet, "D85")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(cell)

		f.SetCellStr(sheet, "G3", "04/18/2024")

		f.SetCellStr(sheet, "D85", "HOU Xiaoyi")
		f.SetCellStr(sheet, "D89", "HOU Xiaoyi")
		f.SetCellStr(sheet, "D91", "HOU Xiaoyi")
		f.SetCellStr(sheet, "D93", "HOU Xiaoyi")

		f.SetCellStr(sheet, "I85", "xiaoyi.hou@cn.bosch.com +86-2122185389")
		f.SetCellStr(sheet, "I89", "xiaoyi.hou@cn.bosch.com +86-2122185389")
		f.SetCellStr(sheet, "I91", "xiaoyi.hou@cn.bosch.com +86-2122185389")
		f.SetCellStr(sheet, "I93", "xiaoyi.hou@cn.bosch.com +86-2122185389")

		f.Save()

	}

}
