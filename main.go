package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

const (
	blackBlock = "\033[40m  \033[0m"
	whiteBlock = "\033[47m  \033[0m"
)

func main() {
	qr, err := qrcode.New("deno", qrcode.Low)
	if err != nil {
		log.Fatalf("failed to generate QR code: %s", err)
	}

	bitmap := qr.Bitmap()
	moduleCount := len(bitmap)

	modules := make([][]bool, moduleCount)
	for i := range modules {
		modules[i] = make([]bool, moduleCount)
		for j := range modules[i] {
			modules[i][j] = bitmap[i][j]
		}
	}

	buf := bytes.Buffer{}

	for _, row := range modules {
		for _, col := range row {
			if col {
				buf.WriteString(blackBlock)
			} else {
				buf.WriteString(whiteBlock)
			}
		}
		buf.WriteString("\n")
	}

	fmt.Println(buf.String())
}
