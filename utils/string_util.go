package utils

import (
	//	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type StringUtil struct {
}

type RGB struct {
	red, green, blue int64
}

type HEX struct {
	str string
}

func (color RGB) rgb2hex() HEX {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return HEX{r + g + b}
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

//首字母大写
func (this *StringUtil) Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

//字符串截取
func (this *StringUtil) SubString(str string, startIndex int, length int) string {
	rs := []rune(str)
	return string(rs[startIndex : startIndex+length])
}

//替换游戏文本日志里的数字型变量为字符串
func (this *StringUtil) FixJsonString(lineText string) string {
	replaceStrings := [][]string{
		[]string{":", ":\""},     //原来的
		[]string{",", "\","},     //原来的
		[]string{"}", "\"}"},     //原来的
		[]string{":\"\\", ":\\"}, //原来的
		[]string{"\"\"", "\""},   //原来的
		[]string{"}\",{", "},{"}, //最后加的
	}

	for _, arr := range replaceStrings {
		lineText = strings.ReplaceAll(lineText, arr[0], arr[1])
	}

	return lineText
}

func (this *StringUtil) ReplaceHTML2UBB(content string) string {
	//map是无序的 所以每次替换的顺序不一样，这里没有问题，底部不能用map
	tags := [][]string{
		[]string{"<strong>", "[b]"}, //粗体
		[]string{"</strong>", "[/b]"},
		[]string{"<em>", "[i]"}, //斜体
		[]string{"</em>", "[/i]"},
		[]string{"<a ", "[a "},
		[]string{"</a> ", "[/a]"},
	}

	//先批量一批 避免标签嵌套的正则读取问题
	for _, v := range tags {
		content = strings.ReplaceAll(content, v[0], v[1])
	}

	tags = [][]string{
		[]string{"text-decoration: underline;", ""}, //下划线
		[]string{"text-align: center;", ""},
		[]string{" style=\" \"", ""},
		[]string{" style=\"\"", ""},
		[]string{"[a ", "<a "},
		[]string{"[/a]", "</a> "},
	}

	//正则替换的 下划线
	reg := regexp.MustCompile(`text-decoration: underline;[\w\W]*?>([\w\W]+?)<`)
	ret := reg.FindAllStringSubmatch(content, -1)
	for _, result := range ret { //遍历匹配出来的几处，进行替换
		newStr := strings.Replace(result[0], result[1], "[u]"+result[1]+"[/u]", 1)
		//		content = reg.ReplaceAllLiteralString(content, newStr)
		content = strings.Replace(content, result[0], newStr, 1)
	}

	//居中 暂不支持等有需求时再调试，要注意多个居中标签情况的替换
	//	reg = regexp.MustCompile(`text-align: center;[\w\W]*?>([\w\W]+?)</`)
	//	ret = reg.FindAllStringSubmatch(content, -1)
	//	for _, result := range ret {
	//		newStr := strings.Replace(result[0], result[1], "<p align=center>"+result[1]+"</p>", 1)
	//		content = reg.ReplaceAllLiteralString(content, newStr)
	//	}

	//字号
	reg = regexp.MustCompile(`font-size: ([\d]+)px[\w\W]+?>([\w\W]+?)</`)
	ret = reg.FindAllStringSubmatch(content, -1)
	for _, result := range ret {
		fontSize := result[1]
		text := result[2]

		newStr := strings.Replace(result[0], text, "[size="+fontSize+"]"+text+"[/size]", 1)
		content = strings.Replace(content, result[0], newStr, 1) //替换这个片段的代码
		//		tags = append(tags, []string{"font-size: " + fontSize + "px;", ""})
	}

	//颜色
	reg = regexp.MustCompile(`color: ([\w\W]+?);[\w\W]+?>([\w\W]+?)</`) //使用?非贪婪模式
	ret = reg.FindAllStringSubmatch(content, -1)
	for _, result := range ret {
		color := result[1]
		text := result[2]

		rgbStr := strings.ReplaceAll(color, "rgb(", "")
		rgbArr := strings.Split(strings.ReplaceAll(rgbStr, ")", ""), ",")

		r, _ := strconv.Atoi(strings.TrimSpace(rgbArr[0]))
		g, _ := strconv.Atoi(strings.TrimSpace(rgbArr[1]))
		b, _ := strconv.Atoi(strings.TrimSpace(rgbArr[2]))
		color1 := RGB{int64(r), int64(g), int64(b)}

		newStr := strings.Replace(result[0], text, "[color=#"+color1.rgb2hex().str+"]"+text+"[/color]", 1)
		content = strings.Replace(content, result[0], newStr, 1) //替换这个片段的代码
		tags = append(tags, []string{"color: " + color + ";", ""})
	}

	//	tags = append(tags, []string{" style=\"\"", ""})
	tags = append(tags, []string{"<p>", ""})
	tags = append(tags, []string{"</p>", "\n"})
	tags = append(tags, []string{"<br/>", "\n"})
	tags = append(tags, []string{"<span>", ""})
	tags = append(tags, []string{"</span>", ""})

	for _, v := range tags {
		content = strings.ReplaceAll(content, v[0], v[1])
	}

	reg = regexp.MustCompile(`<[\w\W]+?>`) //加一个删除所有尖括号里的标签
	ret = reg.FindAllStringSubmatch(content, -1)

	for _, v := range ret {
		content = strings.ReplaceAll(content, v[0], "")
	}

	return content
}
