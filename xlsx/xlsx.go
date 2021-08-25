package xlsx

import (
    "bytes"
    "github.com/tealeg/xlsx"
    "path/filepath"
)

// type
// file
type File = xlsx.File
type Sheet = xlsx.Sheet
type Row = xlsx.Row

const (
    BLANK_STRING = ""
)

// xlsxSetStyleFont
// set cell style font
func xlsxSetStyleFont(style *xlsx.Style, isBold bool) {
    // set font
    font := xlsx.NewFont(10, "Arial")
    font.Bold = isBold
    style.Font = *font
}

// xlsxSetStyleFill
// set cell style fill
func xlsxSetStyleFill(style *xlsx.Style) {
    // set fill
    fill := xlsx.NewFill("solid", "00D3D3D3", "FF000000")
    style.Fill = *fill
}

// xlsxSetStyleBorder
// set cell style border
func xlsxSetStyleBorder(style *xlsx.Style) {
    // set border
    border := xlsx.NewBorder("thin", "thin", "thin", "thin")
    style.Border = *border
}

// xlsxSetStyleAlignment
// set cell style alignment
func xlsxSetStyleAlignment(style *xlsx.Style) {
    // Alignment
    align := xlsx.DefaultAlignment()
    align.Horizontal = "center"
    align.Vertical = "center"
    align.WrapText = true
    style.Alignment = *align
}

//xlsxSetStyleAlignmentForLeft
func xlsxSetStyleAlignmentForLeft(style *xlsx.Style) {
    // Alignment
    align := xlsx.DefaultAlignment()
    align.Horizontal = "left"
    align.Vertical = "center"
    align.WrapText = true
    style.Alignment = *align
}

// xlsxAddCell
func xlsxAddCell(row *Row, value string) *xlsx.Cell {
    cell := row.AddCell()
    cell.Value = value
    return cell
}

// xlsxSetHeaderStyle
func xlsxSetHeaderStyle(cell *xlsx.Cell) {
    style := xlsx.NewStyle()
    defer cell.SetStyle(style)
    xlsxSetStyleFont(style, true)
    xlsxSetStyleFill(style)
    xlsxSetStyleBorder(style)
    xlsxSetStyleAlignment(style)
}

// xlsxSetStyle
func xlsxSetStyle(cell *xlsx.Cell) {
    style := xlsx.NewStyle()
    defer cell.SetStyle(style)
    xlsxSetStyleFont(style, false)
    xlsxSetStyleBorder(style)
    xlsxSetStyleAlignment(style)
}

// xlsxSetStyleForLeft
func xlsxSetStyleForLeft(cell *xlsx.Cell) {
    style := xlsx.NewStyle()
    defer cell.SetStyle(style)
    xlsxSetStyleFont(style, false)
    xlsxSetStyleBorder(style)
    xlsxSetStyleAlignmentForLeft(style)
}

// EXPORT
//
// XlsxNewFile
func XlsxNewFile() *xlsx.File {
    return xlsx.NewFile()
}

// XlsxAddSheet
func XlsxAddSheet(file *xlsx.File, sheetName string) (*Sheet, error) {
    sheet, err := file.AddSheet(sheetName)
    return sheet, err
}

// XlsxAddRow
// xlsx add a row and set cell style
func XlsxAddRow(sheet *Sheet, values []string, isHeader bool) *Row {
    row := sheet.AddRow()
    for _, v := range values {
        cell := xlsxAddCell(row, v)
        // set style
        if isHeader == true {
            xlsxSetHeaderStyle(cell)
        } else {
            xlsxSetStyle(cell)
        }
    }
    return row
}

// XlsxAddRowByLeft
func XlsxAddRowByLeft(sheet *Sheet, values []string) *Row {
    row := sheet.AddRow()
    for _, v := range values {
        cell := xlsxAddCell(row, v)
        // set style
        xlsxSetStyleForLeft(cell)
    }
    return row
}

// set col width
func XlsxSetColWidth(sheet *Sheet, colNum int, width float64) {
    sheet.SetColWidth(0, colNum-1, width)
}

// set row merge
func XlsxRowMerge(sheet *Sheet, rowIndex int, mergeIndexs []uint32, vcells int) {
    row := sheet.Rows[rowIndex]
    for _, index := range mergeIndexs {
        row.Cells[index].Merge(0, vcells)
    }
}

// XlsxCellMergeByCellIndex
func XlsxCellMergeByCellIndex(sheet *Sheet, rowIndex, cellIndex, hcells int) {
    row := sheet.Rows[rowIndex]
    row.Cells[cellIndex].Merge(hcells, 0)
}

// XlsxSaveFile
// create file
func XlsxSaveFile(file *File, filePath string, fileName string) error {
    path := filepath.ToSlash(filePath + fileName)
    err := file.Save(path)
    return err
}

// XlsxReaderFile
func XlsxReaderFile(file *File) ([]byte, error) {
    buffer := new(bytes.Buffer)
    if err := file.Write(buffer); err != nil {
        return nil, err
    }
    return buffer.Bytes(), nil
}
