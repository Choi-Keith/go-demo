package qrcode

import (
	"demo01/util"
	"fmt"
	qrcode2 "github.com/skip2/go-qrcode"
	"image/color"
)

func generateQrcode() {
	// 二维码登录参考链接:https://juejin.cn/post/6946536305115267109
	uuid := util.GenUUID("")
	fmt.Printf("uuid: %s", uuid)
	qrcode2.WriteColorFile(uuid, qrcode2.High, 256, color.RGBA{R: 00, G: 00, B: 00, A: 1}, color.White, "./qr.png")

}
