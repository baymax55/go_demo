package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize"
)

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.gif", ""); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{
        "x_scale": 0.5,
        "y_scale": 0.5
    }`); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}
	// 保存文件
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}
