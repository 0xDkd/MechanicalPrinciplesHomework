package main

import (
	"Matlab/cmd"
	"Matlab/data"
	"Matlab/sport"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"html/template"
	"os"
	"strings"
)


func main() {

	valueData := cmd.GetContent()
	valueData.Sport.Near = template.HTML(strings.Join(strings.Split(sport.C().Near, "<"), "<"))
	valueData.Sport.Remote = template.HTML(strings.Join(strings.Split(sport.C().Remote, "<"), "<"))


	bytes,err := AimerMatlab.Asset("codeTemple/matlab.tpl")
	t := template.New("matlab.tpl")
	t, _ = t.Parse(string(bytes))
	err = t.Execute(os.Stdout, valueData)
	file,err := os.OpenFile("matlab 代码.txt",os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("error", err)
	}
	t.Execute(file,valueData)
	fmt.Println()
	ct.Foreground(ct.Yellow,true)
	fmt.Println("### Matlab代码输出完毕，请到本程序的统计目录可找到 《matlab 代码.txt》 打开就是 matlab 代码 ###")
	fmt.Println()
	fmt.Println("### 或者您也可以直接复制上面的代码,到 matlab 中粘贴, 既可以运行生成您的凸轮 ###")
	fmt.Print()
	fmt.Println("### 如果想要退出程序，请直接关闭窗口即可 ###")
	for{}
}
