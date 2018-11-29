package check

import (
	"math"
)

//---回程验证--
//k=f-d1;
//        s(f)=h*(1-(k/fh)+(sin(2*pi*k/fh)/(2*pi)));s=s(f);
//        ds(f)=(h*(cos(2*pi*k/fh)-1))/fh;ds=ds(f);
func BackSin(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, fh, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = h * (1 - ((index / fh) - (math.Sin((2 * math.Pi * index) / fh))/(2*math.Pi)))
		ds = (h*(math.Cos((2 * math.Pi * index) / fh)) - 1) / fh
		r[i] = comment(ds, e, s, a)
		i++
	}
	return sendData(r, length)
}

//s(f)=.5*h*(1+cos(pi*k/fh));s=s(f);
// ds(f)=-.5*pi*h*sin(pi*k/fh)/(fh*hd);ds=ds(f)
func BackCos(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, fh, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = (h * (1 + math.Cos(math.Pi*index/fh))) / 2
		ds = -math.Pi*h*math.Sin(math.Pi*index/fh)*index/(2*fh)
		r[i] = comment(ds, e, s, a)
		i++
	}
	return sendData(r, length)
}

func BackNormal(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, fh, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		//等加速
		if index <= fh/2 {
			s = 2 * h * (math.Pow(index, 2)) / (math.Pow(fh, 2))
			ds = (4 * h * index * hd) / (math.Pow(fh*hd, 2))
			r[i] = comment(ds, e, s, a)
			i++
			//等减速
		} else {
			s = h - (2*h*math.Pow(fh-index, 2))/math.Pow(fh, 2)
			ds = 4 * h * (fh - index) * hd / math.Pow(fh*hd, 2)
			r[i] = comment(ds, e, s, a)
			i++
		}

	}
	return sendData(r, length)
}

func BackComplex(begin float64, end float64, e float64, h float64, a float64) float64 {
	length, fh, r := getInfo(begin, end)
	for index := begin; index <= end; index += 1 {
		s = (10 * h * math.Pow(index, 3) / math.Pow(fh, 3)) - (15 * h * math.Pow(index, 4) / math.Pow(fh, 4)) + (6 * h * math.Pow(index, 5) / math.Pow(fh, 5))
		ds = (3 * 10 * h / math.Pow(fh, 3)) + (4 * (-15) * h / math.Pow(fh, 4)) + (5 * 6 * h / math.Pow(fh, 5))
		r[i] = comment(ds, e, s, a)
		i++
	}
	return sendData(r, length)
}
