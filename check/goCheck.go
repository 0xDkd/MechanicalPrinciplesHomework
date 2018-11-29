package check

import (
	"math"
)

//----推程验证----

//begin 推程开始角度
//end 推程结束角度
//e 偏心距
//h 最大推程
//a 推程许用压力角
//推程 Sin 验证
func GoSin(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, ft, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = h * ((index / ft) - (math.Sin((2 * math.Pi * index) / ft))/(2*math.Pi))
		ds = (h * (1 - (math.Cos((2 * math.Pi * index) / ft)))) / ft
		r[i] = comment(ds, e, s, a)
		i++
	}
	//fmt.Println(r)
	return sendData(r, length)
}

//s(f)=(h*(1-cos(pi*f/ft)))/2;s=s(f);
//ds(f)=(pi*h*sin(pi*f/ft)*f)/(2*ft);ds=ds(f);
//推程 Cos 验证
func GoCos(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, ft, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = (h * (1 - math.Cos(math.Pi*index/ft))) / 2
		ds = (math.Pi * h * math.Sin(math.Pi*index/ft) * index) / (2 * ft)
		r[i] = comment(ds, e, s, a)
		i++
	}
	return sendData(r, length)
}

//if f<=ft/2
//        s(f) = 2*h*f^2/ft^2; s=s(f);
//        ds(f) = 4*h*f*hd/(ft*hd)^2;ds=ds(f);
//    elseif f>ft/2 && f<=ft
//        s(f)=h-2*h*(ft-f)^2/ft^2;s=s(f);
//        ds(f) = 4*h*(ft-f)*hd/(ft*hd)^2;ds=ds(f);
// 推程 等加速等减速验证
func GoNormal(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, ft, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		//等加速
		if index <= ft/2 {
			s = 2 * h * (math.Pow(index, 2)) / (math.Pow(ft, 2))
			ds = (4 * h * index * hd) / (math.Pow(ft*hd, 2))
			r[i] = comment(ds, e, s, a)
			i++
			//等减速
		} else {
			s = h - (2*h*math.Pow(ft-index, 2))/math.Pow(ft, 2)
			ds = 4 * h * (ft - index) * hd / math.Pow(ft*hd, 2)
			r[i] = comment(ds, e, s, a)
			i++
		}

	}
	return sendData(r, length)
}

//s(f)=((10*h*(k^3)/(fh^3))-(15*h*(k^4)/(fh^4))+(6*h*(k^5)/(fh^5)));s=s(f);
//ds(f)=(3*10*h/(fh^3)+(4*(-15)*h/(fh^4))+(5*6*h/(fh^5)));ds=ds(f);
//五次多项式
func GoComplex(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, ft, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = (10 * h * math.Pow(index, 3) / math.Pow(ft, 3)) - (15 * h * math.Pow(index, 4) / math.Pow(ft, 4)) + (6 * h * math.Pow(index, 5) / math.Pow(ft, 5))
		ds = (3 * 10 * h / math.Pow(ft, 3)) + (4 * (-15) * h / math.Pow(ft, 4)) + (5 * 6 * h / math.Pow(ft, 5))
		r[i] = comment(ds, e, s, a)
		i++
	}
	return sendData(r, length)
}