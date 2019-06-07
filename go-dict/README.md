### Chinese and English translation tools in the command line（命令行下中英文翻译工具）
```
dict is a command line tool by go 
```
![](https://github.com/LockGit/Go/blob/master/img/dict.gif)


## 一键安装(Onekey Install)
```
wget https://raw.githubusercontent.com/LockGit/Go/master/go-dict/dict.go && go build -o /usr/local/bin/dict dict.go 
```


## 使用(Usage)
### 英译中(English To Chinese)
```bash
➜  ~ dict I love you
您的输入: I love you
*******英汉翻译:*******
我爱你。
*******网络释义:*******
I Love You : 我爱你
I really love you : 真的爱你
I Do love you : 我是爱你的
```

### 中译英(Chinese To English)
```bash
➜  ~ dict 我爱你
您的输入: 我爱你
*******英汉翻译:*******
I love you
*******网络释义:*******
我爱你 : I Love You
我也爱你 : I Love You Too
我就爱你 : The Arrangement
```
