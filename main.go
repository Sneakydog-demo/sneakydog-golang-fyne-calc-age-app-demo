package main

import (
	"fmt"
	"os"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	os.Setenv("FYNE_FONT", "./Alibaba-PuHuiTi-Regular.ttf")
	os.Setenv("FYNE_THEME", "light")


	a := app.New()
	w := a.NewWindow("计算年龄")

	w.Resize(fyne.Size{Width: 300, Height: 400})
	w.CenterOnScreen()

	input := widget.NewEntry()
	input.SetPlaceHolder("请输入年龄..")
	hello := widget.NewLabel("")

	w.SetContent(container.NewVBox(input, hello, widget.NewButton("计算", func() {

		int, err := strconv.Atoi(input.Text)

		if err != nil {
			hello.SetText("你好像已经不是人了！！！！")
		} else {

			if int > 0 {
				hello.SetText("你已经" + input.Text + "岁了！！！！	")
			} else{
				hello.SetText(fmt.Sprintf("你已经嗝屁%d年！！！！", int * -1))
			}
 
		}

		input.SetText("")

	})))

	w.ShowAndRun()

}
