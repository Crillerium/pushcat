# 猫猫推送
一个有着猫猫图标的通知推送应用  

# 安装教程
## 一、通过GO INSTALL安装
执行命令:
```
go install github.com/Crillerium/pushcat
```
## 二、通过GO BUILD构建
**(注意: 此方法需要用户自行添加PATH)**
1. 克隆仓库并进入
```
git clone https://github.com/Crillerium/pushcat.git
cd pushcat
```
2. 构建
```
go build -ldflags "-s -w"
```

## 三、通过Releases页下载
**(注意: 此方法需要用户自行添加PATH)**
前往(Releases页)[https://github.com/Crillerium/pushcat/releases]下载已编译文件

# 使用方法
```
pushcat -t 通知标题 -c 通知内容
```