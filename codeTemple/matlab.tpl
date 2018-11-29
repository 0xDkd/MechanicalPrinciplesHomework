%定义参数
rb ={{ .Rb}};%基圆半径
rt = {{ .Rt}};%滚子半径
e ={{ .E}} ;%偏心距
h = {{ .H}};%推程最大值
ft ={{ .Ft}};%推程运动角
fs = {{ .Fs}};%远休止角
fh ={{ .Fh}};%回程运动角
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
    %推程加速状态（{{ .Sport.PushSportName}})
        {{ .Sport.Push}}
        %远休止状态
    {{ .Sport.Remote}}
         %回程运动状态({{ .Sport.ReturnJourneySportName}})
        {{ .Sport.ReturnJourney}}
        %近休止状态
    {{ .Sport.Near}}
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
plot(e+rt*cos(ct),se+rt*sin(ct),'m')
plot(xp,yp,'b')