# 使用
--------------

## 使用 go-katran 

go-katran 由两部分组成

1. BPF 代码（ XDP Hook 触发点） 
2. go 代码，BPF 程序代码的控制平面

## 术语

1. **VIP （Virtual IP Address）**: 服务的IP地址。go-katran 扩展了VIP的定义，除了IP地址，还包括端口和协议(TCP或UDP)
2. **Real** : 后端服务器的IP地址，在那里流量将被重定向。

# 案例



