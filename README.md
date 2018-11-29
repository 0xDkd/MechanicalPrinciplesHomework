<center>
# 机械原理大作业解析法分析
</center>

[TOC]

## 工具

本次的作业使用的画图的工具为 Matlab，Matlab 的代码可以帮助我们更好的去理解整个的运动过程和运动曲线的是如何形成的

## 分析运动状态

按照学号分配题目，我是20 号，下面是我的原始数据

| 名称|数据|
| --- | :-: |
| 近休止状态 | 0˚ - 30˚ |
| 推程运动状态 | 30˚ - 200˚  |
| 远休止状态 | 200˚ - 270˚ |
| 回程运动状态 | 270˚ - 360˚ |
| 初选基圆半径 | 20 |
| 偏心距 | +10 |
| 滚子半径 | 14 |
| 最大推程 | 30 |


推程运动为: 正弦加速度运动
回程运动为: 等加速等减速运动

可以确定

**推程**的运动方程为:
`s＝h[δ/δ0-sin(2πδ/δ0)/2π]`
**回程**的运动方程为:
`s ＝h-2hδ2/δ’20` (前半段)
`s ＝2h(δ’0-δ)2/δ’20` (后半段)

近休止和远休止都不会影响滚子推杆的位移。
于是我们可以编写如下的 matlab 代码

```matlab
rb =50;%基圆半径
rt = 14;%滚子半径
e =10 ;%偏心距
h = 30;%推程最大值
ft =170;%推程运动角
fs = 70;%远休止角
fh =90;%回程运动角
hd =pi /180;
du = 180/pi;%定义° 的含义
se =sqrt(rb^2-e^2);
d1 = ft+fs;
d2 = ft+fs+fh;
n=360;

%开始运动
if f<=ft
    s(f)=h*((f/ft)-(sin((2*pi*f)/ft))/(2*pi));s=s(f);
    ds(f)=(h*(1-cos(2*pi*f/ft)))/ft;ds=ds(f);
    %远休止状态

elseif f>ft && f<=d1
    s=h;ds=0;
    %回程运动状态(等加速减速运动)

elseif f>d1 && f<=d1+(fh/2)
     k=f-d1; % k
    s(f)=h-(2*h*(k^2)/(fh^2));s=s(f);
    ds(f)=-4*h*k/(fh^2);ds=ds(f);
elseif f > d1+(fh/2) && f <= d2
    k=f-d1;
    s(f)=(2*h*((fh-k)^2))/(fh^2);s=s(f);
    ds(f)=(-4*h*(fh-k))/(fh^2);ds=ds(f);
    %近休止状态

elseif f>d2 &&  f<=n
    s=0;ds=0;
end
    
```

接下来只需要把凸轮的轮廓方程带入

```matlab
%计算轮廓
    xx(f)=(se+s)*sin(f*hd)+e*cos(f*hd);x=xx(f);
    yy(f)=(se+s)*cos(f*hd)-e*sin(f*hd);y=yy(f);
    dx(f)=(ds-e)*sin(f*hd)+(se+s)*cos(f*hd);dx=dx(f);
    dy(f)=(ds-e)*cos(f*hd)-(se+s)*sin(f*hd);dy=dy(f);
    xp(f)=x+rt*dy/sqrt(dx^2+dy^2);xxp=xp(f);
    yp(f)=y-rt*dx/sqrt(dx^2+dy^2);yyp=yp(f);

end
```

计算出了凸轮的轮廓，接下来就是画图了，使用 `plot` 函数即可把图像画出

```matlab
plot(xx,yy,'r-.')
axis([-(rb+h-10) (rb+h+10) -(rb+h+10) (rb+rt+10)])

axis equal
text(rb+h+3,0,'X')
text(0,rb+rt+3,'Y')
text(-5,5,'O')
title('偏置移动从动件盘型凸轮设计')
hold on;

plot([-(rb+h) (rb+h)],[0 0],'k')
plot([0 0],[-(rb+h) (rb+rt)],'k')
plot([e e],[0 (rb+rt)],'k--')
ct = linspace(0,2*pi);
plot(e+rt*cos(ct),se+rt*sin(ct),'m')
plot(xp,yp,'b')
```

至此，完整的代码为

```matlab
%定义参数
rb =50;%基圆半径
rt = 14;%滚子半径
e =10;%偏心距
h = 30;%推程最大值
ft =170;%推程运动角
fs = 70;%远休止角
fh =90;%回程运动角
hd =pi /180;
du = 180/pi;%定义° 的含义
se =sqrt(rb^2-e^2);
d1 = ft+fs;
d2 = ft+fs+fh;
n=360;

%% 初始化坐标
s=zeros(n);
ds = zeros(n);
x= zeros(n);
y = zeros(n);
dx =zeros(n);
dy =zeros(n);
xx = zeros(n);
yy = zeros(n);
yp = zeros(n);

%% 开始计算

for f=1:n
    
    %推程加速状态（正弦运动）
    if f<=ft
        % σ_0 ==> ft  ||| σ ==> f
        s(f)=h*((f/ft)-(sin((2*pi*f)/ft))/(2*pi));s=s(f);
        ds(f)=(h*(1-cos(2*pi/ft)))/ft;ds=ds(f);
        %远休止状态
    elseif f>ft && f<=d1
        s=h;ds=0;
        %回程运动状态(加速度) ft==> pi*f-ft-fs
    elseif f>d1 && f<=d1+(fh/2)
        k=f-d1; % k
        s(f)=h-(2*h*(k^2)/(fh^2));s=s(f);
        ds(f)=-4*h*k/(fh^2);ds=ds(f);
        %回程运动状态(减速度)
    elseif f > d1+(fh/2) && f <= d2
        k=f-d1;
        s(f)=(2*h*((fh-k)^2))/(fh^2);s=s(f);
        ds(f)=(-4*h*(fh-k))/(fh^2);ds=ds(f);
        %近休止状态
    elseif f>d2 &&  f<=n
        s=0;ds=0;
    end
    
    
    %计算轮廓
    xx(f)=(se+s)*sin(f*hd)+e*cos(f*hd);x=xx(f);
    yy(f)=(se+s)*cos(f*hd)-e*sin(f*hd);y=yy(f);
    dx(f)=(ds-e)*sin(f*hd)+(se+s)*cos(f*hd);dx=dx(f);
    dy(f)=(ds-e)*cos(f*hd)-(se+s)*sin(f*hd);dy=dy(f);
    xp(f)=x+rt*dy/sqrt(dx^2+dy^2);xxp=xp(f);
    yp(f)=y-rt*dx/sqrt(dx^2+dy^2);yyp=yp(f);
    
end

plot(xx,yy,'r-.')
axis([-(rb+h-10) (rb+h+10) -(rb+h+10) (rb+rt+10)])

axis equal
text(rb+h+3,0,'X')
text(0,rb+rt+3,'Y')
text(-5,5,'O')
title('偏置移动从动件盘型凸轮设计')
hold on;

plot([-(rb+h) (rb+h)],[0 0],'k')
plot([0 0],[-(rb+h) (rb+rt)],'k')
plot([e e],[0 (rb+rt)],'k--')
ct = linspace(0,2*pi);
plot(rb*cos(ct),rb*sin(ct),'g')
plot(e*cos(ct),e*sin(ct),'c--')
plot(e+rt*cos(ct),se+rt*sin(ct),'m')
plot(xp,yp,'b')
```
## 生成器分析

请看[我的博客](https://0w0.tn)
