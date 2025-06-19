package main

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/disintegration/imaging"
	"github.com/skip2/go-qrcode"
)

func main() {
	// Test QR code generation
	// qr, err := createQRCodeWithLogo("https://example.com", "logo.png", "qr.png")
	// if err != nil {
	// 	fmt.Println("Error generating QR code:", err)
	// }
	// fmt.Println("QR code generated:", qr)
	target := "https://www.baidu.com"
	err := openBrowser(target)
	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}

// 冒泡算法
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// create qrcode with logo
func createQRCodeWithLogo(text string, logoPath string, qrPath string) (string, error) {
	// Generate QR code
	qr, err := qrcode.New(text, qrcode.High)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	// Get QR code as image
	qrImage := qr.Image(256)

	// Open and decode logo file
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return "", fmt.Errorf("failed to open logo file: %w", err)
	}
	defer logoFile.Close()

	logo, err := png.Decode(logoFile)
	if err != nil {
		return "", fmt.Errorf("failed to decode logo: %w", err)
	}

	// Calculate logo size (25% of QR code size)
	logoSize := qrImage.Bounds().Dx() / 4
	logo = imaging.Resize(logo, logoSize, logoSize, imaging.Lanczos)

	// Calculate position to center the logo
	x := (qrImage.Bounds().Dx() - logoSize) / 2
	y := (qrImage.Bounds().Dy() - logoSize) / 2

	// Create new image and combine QR code with logo
	finalQR := imaging.New(256, 256, color.White)
	finalQR = imaging.Paste(finalQR, qrImage, image.Point{0, 0})
	finalQR = imaging.Paste(finalQR, logo, image.Point{x, y})

	// Save the final image
	err = imaging.Save(finalQR, qrPath)
	if err != nil {
		return "", fmt.Errorf("failed to save QR code: %w", err)
	}

	return "QR code with logo generated successfully", nil
}

// 使用chromedp启动chrome浏览器打开指定网页
func openBrowser(url string) error {
	ctx := context.Background()
	// 创建一个新的Chrome会话
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// 关闭无头模式，方便调试
		chromedp.Flag("headless", false),
		// 防止监测webdriver
		chromedp.Flag("enable-automation", false),
		// 禁用 blink 特征
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		// 忽略浏览器的风险提示（但好像并没什么用）
		chromedp.Flag("ignore-certificate-errors", true),
		// 关闭浏览器声音（也没用）
		chromedp.Flag("mute-audio", false),
		// 设置浏览器尺寸
		chromedp.WindowSize(1150, 1000),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()
	taskCtx, taskCancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer taskCancel()

	return chromedp.Run(taskCtx, chromedp.Navigate(url))
}

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
