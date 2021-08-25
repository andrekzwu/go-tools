package main

import (
	"fmt"

	"github.com/andrezhz/go-tools/xlsx"
)

func main() {
	createXlsx()
}

// createXlsx
func createXlsx() error {
	// 1. 创建XLSX文件实例
	xlsxFile := xlsx.XlsxNewFile()
	// 2. 添加sheet信息
	sheet, err := xlsx.XlsxAddSheet(xlsxFile, "订单汇总")
	if err != nil {
		return err
	}
	// 行数的索引，从0开始
	rowIndex := 0
	// 3. 添加行的头部信息， 第三个参数，true为头部数据，带有灰色底色
	headers := []string{"订单", xlsx.BLANK_STRING, xlsx.BLANK_STRING, xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, headers, true)
	// 头部这一行列合并，对第rowIndex行的第cellIndex列，在水平方向上合并N格
	// 当前示例为，在第0行的第0列开始，水平向右合并三列，即第 0~3列，4个单元格合并为一个单元格
	xlsx.XlsxCellMergeByCellIndex(sheet, rowIndex, 0, 3)
	rowIndex++
	// 3. 添加一行记录，向左对齐
	values := []string{"andrezhz", xlsx.BLANK_STRING, xlsx.BLANK_STRING, xlsx.BLANK_STRING}
	xlsx.XlsxAddRowByLeft(sheet, values)
	// 水平合并单元格
	xlsx.XlsxCellMergeByCellIndex(sheet, rowIndex, 0, 3)
	rowIndex++
	// 4. 添加一行留白数据
	values = []string{xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, values, true)
	xlsx.XlsxCellMergeByCellIndex(sheet, rowIndex, 0, 3)
	rowIndex++
	// 5. 添加一行头部，多个单元格数据，不合并
	values = []string{"订单项目", xlsx.BLANK_STRING, "订单金额", xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, values, true)
	rowIndex++
	// 添加一行记录，多个单元格
	values = []string{"订单金额", xlsx.BLANK_STRING, "10.00", xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, values, false)
	rowIndex++
	// 订单金额
	values = []string{"售后金额", xlsx.BLANK_STRING, "1.00", xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, values, false)
	rowIndex++
	// 总计
	values = []string{"总计", xlsx.BLANK_STRING, "9.00", xlsx.BLANK_STRING}
	xlsx.XlsxAddRow(sheet, values, true)
	rowIndex++
	// 6 添加第二个sheet
	sheet2, err := xlsx.XlsxAddSheet(xlsxFile, "订单数据")
	if err != nil {
		return err
	}
	rowIndex2 := 0
	// 添加一个头部信息
	headers2 := []string{"订单号", "SKUID", "数量", "单价", "金额"}
	xlsx.XlsxAddRow(sheet2, headers2, true)
	rowIndex2++
	// 添加多行记录
	// 记录一下当前行的索引
	startOrderRowIndex := rowIndex2
	//
	values2 := []string{"2dadasd020110414511723367", "4141735", "1", "50", "50"}
	xlsx.XlsxAddRow(sheet2, values2, false)
	rowIndex2++
	//
	values2 = []string{"2dadasd020110414511723367", "3141736", "2", "30", "60"}
	xlsx.XlsxAddRow(sheet2, values2, false)
	rowIndex2++
	// 某一行rowIndex2的某一个单元格cellIndex向下垂直合并多少行
	// 当前示例，将第1行的，第1，2，3，4个单元格，垂直想下合并一个单元格
	xlsx.XlsxRowMerge(sheet2, startOrderRowIndex, []uint32{1, 2, 3, 4}, 1)
	// 总计
	values2 = []string{"金额小计", xlsx.BLANK_STRING, xlsx.BLANK_STRING, xlsx.BLANK_STRING, "110"}
	xlsx.XlsxAddRow(sheet2, values2, true)
	// 7. 保存该xlsx文件到本地
	if err := xlsx.XlsxSaveFile(xlsxFile, "./", "测试文件.xlsx"); err != nil {
		return err
	}
	// 或者将文件内容转化为字节流
	bytes, err := xlsx.XlsxReaderFile(xlsxFile)
	if err != nil {
		fmt.Println(bytes)
		return err
	}
	return nil
}
