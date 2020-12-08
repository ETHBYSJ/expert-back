package docx

import (
	"errors"
	"github.com/unidoc/unioffice/document"
)

var ErrSize error = errors.New("尺寸错误")

// 简单封装
func Parse(path string) ([]document.Table, error) {
	doc, err := document.Open(path)
	if err != nil {
		return nil, err
	}
	// 获取表格
	tables := doc.Tables()
	return tables, nil
}

// 获得某单元格内容
func GetCell(table *document.Table, row int, col int) (string, error) {
	// 检查
	rows := table.Rows()
	if len(rows) <= row {
		return "", ErrSize
	}
	cells := rows[row].Cells()
	if len(cells) <= col {
		return "", ErrSize
	}
	cell := cells[col]
	res := ""
	paragraphs := cell.Paragraphs()
	for _, para := range paragraphs {
		for _, run := range para.Runs() {
			res += run.Text()
		}
	}
	return res, nil
}
