/*
Copyright (c) [2022] [开源十年]
[开源十年] is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/
package FeatureJSONTools

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"unsafe"
)

/*
*时间戳和日期的相互转换功能实现
*params:
*app:由主界面传过来的fyne.App实例
 */
func JsonTool(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("JSON工具箱")
	//title
	title := canvas.NewText("JSON工具箱", color.Black)
	content := container.New(layout.NewHBoxLayout(), title, layout.NewSpacer())
	//开始画界面,使用一个大的输入框+按钮组成
	inputJsonStr := widget.NewMultiLineEntry()
	inputJsonStr.SetMinRowsVisible(12)//默认显示三行，强制显示12行
	inputJsonStr.Wrapping=fyne.TextWrapBreak
	errorStr := canvas.NewText("", color.RGBA{255,51,51,255})//用作提示
	btnTurn := widget.NewButton("校验/格式化", func() {
		jsonStr := inputJsonStr.Text
		if len(jsonStr) == 0{
			refreshText(errorStr,"兄弟，你啥也不输入没法给你校验呀？")
		}
		inputJsonStr.SetText(jsonCheckAndFormat(errorStr,inputJsonStr.Text))
		inputJsonStr.Refresh()
	})
	btnArroy := container.New(layout.NewHBoxLayout(),btnTurn,layout.NewSpacer())

	//界面画好了，装进容器里面去显示
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content,inputJsonStr,btnArroy,errorStr))
	myWindow.Resize(fyne.NewSize(600, 300))
	return myWindow
}
//更新提示内容
func refreshText(str *canvas.Text, s string) {
	str.Text = s
	str.Refresh()
}
//对JSON的校验和格式化 传进来的是json字符串
func jsonCheckAndFormat(errorStr *canvas.Text, text string) string {

	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &result)
	if err != nil {
		se, _ := err.(*json.SyntaxError)
		field := string([]byte(text)[:se.Offset])
		if i := strings.LastIndex(field, `":`); i >= 0 {
			field = field[:i]
			if j := strings.LastIndex(field, `"`); j >= 0 {
				field = field[j+1:]
			}
		}
		refreshText(errorStr,err.Error()+" for "+field)
		return text
	}

	js,_ := json.MarshalIndent(result, "", "\t")
	refreshText(errorStr,"Congratulations!")
	return BytesToString(js)

}
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

