package sport

//推程
type SportPush struct {
	//推程Sin
	Sin string
	//推程Cos
	Cos string
	//推程加减速
	Normal string
	//推程五次多项式
	Complex string
}
//回程公式
type SportReturn struct {
	//回程Sin
	Sin string
	//回程Cos
	Cos string
	//回程加减速
	Normal string
	//回程五次多项式
	Complex string
}

//防止转义使用此结构体
type Comment struct {
	//远休止
	Remote string
	//近休止
	Near string
}


func Go() *SportPush {
	//初始化
	data :=&SportPush{}

	//Sin 公式
	data.Sin = `
		if f<=ft
        s(f)=h*((f/ft)-(sin((2*pi*f)/ft))/(2*pi));s=s(f);
        ds(f)=(h*(1-cos(2*pi*f/ft)))/ft;ds=ds(f);`

    data.Cos = `
		if f<=ft
		s(f)=(h*(1-cos(pi*f/ft)))/2;s=s(f);
        ds(f)=(pi*h*sin(pi*f/ft)*f)/(2*ft);ds=ds(f);`

    data.Normal = `
		if f<=ft/2
        s(f) = 2*h*f^2/ft^2; s=s(f);
        ds(f) = 4*h*f*hd/(ft*hd)^2;ds=ds(f);
    elseif f>ft/2 && f<=ft
        s(f)=h-2*h*(ft-f)^2/ft^2;s=s(f);
        ds(f) = 4*h*(ft-f)*hd/(ft*hd)^2;ds=ds(f);`

     data.Complex=`
	if f<=ft
		s(f)=((10*h*(f^3)/(ft^3))-(15*h*(f^4)/(ft^4))+(6*h*(f^5)/(ft^5)));s=s(f);
		ds(f)=(3*10*h/(ft^3)+(4*(-15)*h/(ft^4))+(5*6*h/(ft^5)));ds=ds(f);`

	return data
}
func Back() *SportReturn {
	//初始化
	data :=&SportReturn{}

	data.Sin = `
	elseif f>d1 && f<=d2
        k=f-d1;
        s(f)=h*(1-(k/fh)+(sin(2*pi*k/fh)/(2*pi)));s=s(f);
        ds(f)=(h*(cos(2*pi*k/fh)-1))/fh;ds=ds(f);`

	data.Cos = `
	elseif f>d1 && f<=d2
        k=f-d1;
        s(f)=.5*h*(1+cos(pi*k/fh));s=s(f);
        ds(f)=-.5*pi*h*sin(pi*k/fh)/(fh*hd);ds=ds(f);`

    data.Normal =`
	elseif f>d1 && f<=d1+(fh/2)
        k=f-d1; % k
        s(f)=h-(2*h*(k^2)/(fh^2));s=s(f);
        ds(f)=-4*h*k/(fh^2);ds=ds(f);
    elseif f > d1+(fh/2) && f <= d2
        k=f-d1;
        s(f)=(2*h*((fh-k)^2))/(fh^2);s=s(f);
        ds(f)=(-4*h*(fh-k))/(fh^2);ds=ds(f);`

    data.Complex=`
	elseif f>d1 && f<=d2
		k=f-d1;
		s(f)=h-((10*h*(k^3)/(fh^3))-(15*h*(k^4)/(fh^4))+(6*h*(k^5)/(fh^5)));s=s(f);
		ds(f)=-(3*10*h/(fh^3)+(4*(-15)*h/(fh^4))+(5*6*h/(fh^5)))ds=ds(f);
`
	return data
}

func C() *Comment {
	data :=&Comment{}

	data.Near = `
	elseif f>d2 &&  f<=n
        s=0;ds=0;`

	data.Remote= `
	elseif f>ft && f<=d1
        s=h;ds=0;`

	return data
}