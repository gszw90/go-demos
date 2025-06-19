package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx"
)

// 将map[string]interface转换成指定的结构体
func MapToStruct(m map[string]interface{}, s interface{}) error {
	// 检查 m 是否为 nil
	if m == nil {
		return fmt.Errorf("input map cannot be nil")
	}

	// 检查 s 是否为 nil
	if s == nil {
		return fmt.Errorf("output struct cannot be nil")
	}

	// 检查 s 是否是指针类型
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("output struct must be a pointer")
	}

	// 检查指针是否为 nil
	if v.IsNil() {
		return fmt.Errorf("output struct pointer cannot be nil")
	}

	// 将 map 转换为 JSON 字节
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %w", err)
	}

	// 将 JSON 字节解析到结构体中
	if err := json.Unmarshal(b, s); err != nil {
		return fmt.Errorf("failed to unmarshal into struct: %w", err)
	}

	return nil
}

// 判断一个map是否为空
func IsMapEmpty(m map[string]interface{}) bool {
	return len(m) == 0
}

// 读取指定的excel文件，第一行是标题，从第二行开始读取数据，返回一个map
func ReadExcelToMap(filePath string) (map[string]interface{}, error) {
	// 读取excel文件
	rows, err := ReadExcel(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read excel file: %w", err)
	}
	// 判断rows是否为空
	if len(rows) == 0 {
		return nil, fmt.Errorf("excel file is empty")
	}
	// 创建一个map
	m := make(map[string]interface{})
	// 获取第一行数据
	firstRow := rows[0]
	// 从第二行开始读取数据，并将数据存入map中
	for i := 1; i < len(rows); i++ {
		// 获取当前行数据
		row := rows[i]
		// 遍历当前行数据
		for j, cell := range row {
			// 将当前行数据作为key，第一行数据作为value，存入map中
			m[firstRow[j]] = cell
		} // 将map存入map中
	}
	// 返回map
	return m, nil
}

func ReadExcel(filePath string) ([][]string, error) {
	// 打开excel文件
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	// 获取第一个sheet
	sheet := file.Sheets[0]
	// 创建一个二维切片
	rows := make([][]string, len(sheet.Rows))
	// 遍历sheet的行
	for i, row := range sheet.Rows {
		// 创建一个一维切片
		cells := make([]string, len(row.Cells))
		// 遍历行的单元格
		for j, cell := range row.Cells {
			// 将单元格的值作为字符串存入一维切片中
			cells[j] = cell.String()
		}
		// 将一维切片存入二维切片中
		rows[i] = cells
	}
	// 返回二维切片
	return rows, nil
}

// 将字符串转化为字节切片
func StringToBytes(s string) []byte {
	return []byte(s)
}

// 将字符串传华为bool,如果字符串为空,0,false,False,FALSE,f等就返回false，否则返回true
func StringToBool(s string) bool {
	// 将字符串转换为小写
	s = strings.ToLower(s)
	// 判断字符串是否为空,0,false,False,FALSE等
	if s == "" || s == "0" || s == "f" || s == "false" {
		return false
	}
	return true
}

// 实现一个in_array的方法,使用泛型来实现
func InArray[T comparable](item T, array []T) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

// 利用泛型实现一个map_reduce
func MapReduce[T any, R any](data []T, mapFunc func(T) R, reduceFunc func(R, R) R) R {
	result := make([]R, len(data))
	for i, v := range data {
		result[i] = mapFunc(v)
	}
	finalResult := result[0]
	for _, v := range result[1:] {
		finalResult = reduceFunc(finalResult, v)
	}
	return finalResult
}
