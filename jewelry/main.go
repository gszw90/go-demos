package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Gem represents the structure of the data in the Excel file.
type Gem struct {
	Number               int    `json:"Number"`
	Name                 string `json:"Name"`
	Type                 string `json:"Type"`
	PriceRange           string `json:"PriceRange"`
	CulturalConnotation  string `json:"CulturalConnotation"`
	PatternAnalysis      string `json:"PatternAnalysis"`
	CarvingCraftsmanship string `json:"CarvingCraftsmanship"`
	WebRecommendation    string `json:"WebRecommendation"`
}

func main() {
	f, err := excelize.OpenFile("./jewelry.xlsx")
	if err != nil {
		log.Fatalf("无法打开Excel文件: %v", err)
	}
	defer f.Close()

	// 假设数据在第一个工作表中，并且从第一行开始（标题行）
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("无法读取行: %v", err)
	}

	var gems []Gem

	// 跳过标题行，从第二行开始读取数据
	for i, row := range rows {
		if i == 0 {
			continue // 跳过标题行
		}

		if len(row) < 8 {
			log.Printf("行 %d 的列数不足，跳过", i+1)
			continue
		}

		// 转换行数据到Gem结构体
		gem := Gem{
			Number:               parseInt(row[0]),
			Name:                 row[1],
			Type:                 row[2],
			PriceRange:           row[3],
			CulturalConnotation:  row[4],
			PatternAnalysis:      row[5],
			CarvingCraftsmanship: row[6],
			WebRecommendation:    row[7],
		}

		gems = append(gems, gem)
	}

	// 将gems切片序列化为JSON
	jsonData, err := json.MarshalIndent(gems, "", "  ")
	if err != nil {
		log.Fatalf("无法序列化为JSON: %v", err)
	}

	// 将JSON数据写入文件
	file, err := os.Create("gems.json")
	if err != nil {
		log.Fatalf("无法创建文件: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("无法写入文件: %v", err)
	}

	fmt.Println("JSON数据已成功写入gems.json文件")
}

// parseInt 是一个辅助函数，用于将字符串转换为整数，如果转换失败则返回0
func parseInt(s string) int {
	n, _ := strconv.Atoi(s)

	// if n, err := fmt.Sscanf(s, "%d", &n); err == nil && n == 1 {
	// 	return n
	// }
	return n
}
