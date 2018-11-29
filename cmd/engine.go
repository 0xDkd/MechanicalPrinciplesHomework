package cmd

import (
	"Matlab/sport"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"html/template"
	"strconv"
	"strings"
)

type MatlabValue struct {
	Rb int
	Rt int
	E int
	H int
	Ft int
	Fs int
	Fh int
	Sport SportStatus
}

type SportStatus struct {
	Push template.HTML
	ReturnJourney template.HTML
	PushSportName template.HTML
	ReturnJourneySportName template.HTML
	Remote template.HTML
	Near template.HTML
}
var(
	Rb int
	Rt int
	E int
	H int
	Ft int
	Fs int
	Fh int
	//Push
	P int
	//Back
	B int
	//Config
	C interface{}
)

func init()  {
	ct.Foreground(ct.Green,true)
	welcome :=`
		##########################################
		# Project :解析法画凸轮 Matlab 代码生成器    #
		# Class: Vehicle engineering         	 #
		# Author : RTS                        #
		##########################################`
	fmt.Println(welcome)
	fmt.Print("\n\n")
}

func GetContent() *MatlabValue {
	ct.Foreground(ct.Cyan,false)
	s :=[]string{"等加等减速度运动","余弦加速度运动","正弦加速度运动","五次多项式运动"}
	data := &MatlabValue{}
	fmt.Println("		==请输入基圆半径==")
	fmt.Scanln(&Rb)

	fmt.Println("		==请输入滚子半径==")
	fmt.Scanln(&Rt)

	fmt.Println("		==请输入偏心距==")
	fmt.Scanln(&E)

	fmt.Println("		==请输入推程==")
	fmt.Scanln(&H)

	fmt.Println("		==请输入推程运动角==")
	fmt.Scanln(&Ft)

	fmt.Println("		==请输入远休止角==")
	fmt.Scanln(&Fs)

	fmt.Println("		==请输入回程运动角==")
	fmt.Scanln(&Fh)

	//选择推程
	fmt.Println("		===请输入您所需要的 *推程* 运动的序号，按回车确定(只能选一项)===")
	fmt.Print("		1."+s[0]+"\n")
	fmt.Print("		2."+s[1]+"\n")
	fmt.Print("		3."+s[2]+"\n")
	fmt.Print("		4."+s[3]+"\n")
	fmt.Scanln(&P)

	//选择回程
	fmt.Println("		===请输入您所需要的 *回程* 运动的序号，按回车确定(只能选一项)===")
	fmt.Print("		1."+s[0]+"\n")
	fmt.Print("		2."+s[1]+"\n")
	fmt.Print("		3."+s[2]+"\n")
	fmt.Print("		4."+s[3]+"\n")
	fmt.Scanln(&B)

	//开始对推程和回程进行选择
	switch P {
	case 1:
		data.Sport.Push =  template.HTML(strings.Join(strings.Split(sport.Go().Normal, "<"), "<"))
		data.Sport.PushSportName = "等加等减速运动"
	case 2:
		data.Sport.Push = template.HTML(strings.Join(strings.Split(sport.Go().Cos, "<"), "<"))
		data.Sport.PushSportName = "余弦加速度运动"
	case 3:
		data.Sport.Push =  template.HTML(strings.Join(strings.Split(sport.Go().Sin, "<"), "<"))
		data.Sport.PushSportName = "正弦加速度运动"
	case 4:
		data.Sport.Push =  template.HTML(strings.Join(strings.Split(sport.Go().Complex, "<"), "<"))
		data.Sport.PushSportName = "五次多项式运动"
	default:
		fmt.Print("		输入的编码不正确，程序退出")
	}
	switch B {
	case 1:
		data.Sport.ReturnJourneySportName = "等加速减速运动"
		data.Sport.ReturnJourney =  template.HTML(strings.Join(strings.Split(sport.Back().Normal, "<"), "<"))
	case 2:
		data.Sport.ReturnJourneySportName = "余弦加速度运动"
		data.Sport.ReturnJourney = template.HTML(strings.Join(strings.Split(sport.Back().Cos, "<"), "<"))
	case 3:
		data.Sport.ReturnJourneySportName = "正弦加速度运动"
		data.Sport.ReturnJourney =  template.HTML(strings.Join(strings.Split(sport.Back().Sin, "<"), "<"))
	case 4:
		data.Sport.ReturnJourneySportName = "五次多项式运动"
		data.Sport.ReturnJourney =  template.HTML(strings.Join(strings.Split(sport.Back().Complex, "<"), "<"))
	default:
		fmt.Print("		输入的编码不正确，程序退出")
	}

	//开始赋值给结构体
	data.Ft = Ft
	data.Rb = Rb
	data.E = E
	data.H = H
	data.Fh = Fh
	data.Fs = Fs
	data.Rt = Rt

	backStr := "		##########################################\n"+
		"		#	您的作图需求为:			 #\n"+
		"		#	推程运动为: "+s[P-1]+"	 #\n"+
		"		#	回程运动为: "+s[B-1]+"	 #\n"+
		"		#############<==各项数据为==>#############\n"+
		"		#	基圆半径: "+strconv.Itoa(Rb)+"			 #\n"+
		"		#	滚子半径: "+strconv.Itoa(Rt)+"			 #\n"+
		"		#	偏心距: "+strconv.Itoa(E)+"			 #\n"+
		"		#	最大推程: "+strconv.Itoa(H)+"			 #\n"+
		"		#	推程运动角: "+strconv.Itoa(Ft)+"			 #\n"+
		"		#	远休止角: "+strconv.Itoa(Fs)+"			 #\n"+
		"		#	回程运动角: "+strconv.Itoa(Fh)+"			 #\n"+
		"		##########################################\n"
	fmt.Print(backStr)
	ct.Foreground(ct.Yellow,true)
	fmt.Print("========================生成 MatLab 代码============================\n")
	fmt.Println("")
	ct.Foreground(ct.Blue,true)
	return data

}
