时间过得真快呀，转眼半个月就过去了。

前段时间接手一个项目，连续5天都是凌晨3点左右才睡觉，中间有一天休息时外面的鸡都开始打鸣了~

就一个感觉，累~

后面趁着周末两天在家好好补了补觉，已经过去四五天了，感觉每天还是睡不醒~

记得高中的时候，一周7天能出去包宿8次0.0

果然是上了年纪不得不认输呀

昨天晚上抽时间把JSON校验及格式化做了一下

虽说这个项目没有给我带来任何收益，但是我还是愿意坚持去维持

在做的过程中总是能给我带来一些意想不到的收获

废话不多说了，下面进入主题

#### 使用Go实现JSON格式校验及格式化输出

这个功能就两个目的：

- 校验JSON的格式正确不正确
- 把输入的JSON按照格式化输出，方便阅读

#### 界面

界面比较简单，只要能满足输入JSON和输出格式化后的JSON以及有一个提示位置即可。

```go
	inputJsonStr := widget.NewMultiLineEntry()
	inputJsonStr.SetMinRowsVisible(12)//默认显示三行，强制显示12行
	inputJsonStr.Wrapping=fyne.TextWrapBreak
```

这是一个输入框，我们也当做输出框使用，NewMultiLineEntry是一个多行的输入框，为了显示效果，我们强制显示12行。

```go
errorStr := canvas.NewText("", color.RGBA{255,51,51,255})//用作提示
```

这里实现一个Text文字显示组件，用于我们校验/格式化过程中出现问题后的提示。

```Go 
btnTurn := widget.NewButton("校验/格式化", func() {
		jsonStr := inputJsonStr.Text
		if len(jsonStr) == 0{
			refreshText(errorStr,"兄弟，你啥也不输入没法给你校验呀？")
		}
		inputJsonStr.SetText(jsonCheckAndFormat(errorStr,inputJsonStr.Text))
		inputJsonStr.Refresh()
	})
```

实现一个按钮NewButton，这个按钮用于处理对输入数据的处理事件触发。

#### 核心代码

核心代码就是jsonCheckAndFormat函数部分了，共计十来行代码,主要部分也就那三行，看代码注释：

```go
//对JSON的校验和格式化 传进来的是json字符串
func jsonCheckAndFormat(errorStr *canvas.Text, text string) string {
	//先把输入的字符串进行map接收，如果不用map会导致最后格式化时出问题
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &result)
  //下面代码的意思是如果JSON格式不正确，我们去定位不正确的位置
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
	//能走到这里，说明JSON校验是通过了，下面是按照JSON格式化输出，关键是MarshalIndent函数的作用
	js,_ := json.MarshalIndent(result, "", "\t")
	refreshText(errorStr,"Congratulations!")
	return BytesToString(js)

}
```

#### 知识点补充

- JSON（JavaScript Object Notation, JS对象简谱）是一种轻量级的数据交换格式，易于人阅读和编写，可以在多种语言之间进行数据交换
- JSON格式能够直接为服务器端代码使用, 大大简化了服务器端和客户端的代码开发量, 但是完成的任务不变, 且易于维护
- 标准库encoding/json包是Golang的标准库之一，常用的函数主要是：Marshal、Unmarshal、HTMLEscape、MarshalIndent、Compact等，掌握这几个后基本就可以熟练使用encoding/json