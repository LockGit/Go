### Chinese and English translation tools in the command line（命令行下中英文翻译工具）
```
据说内置的有道翻译的Key因为调用量过度频繁导致被封了？
所以我就直接抓了Chrome浏览器的有道插件的数据包，请求这个api应该没有什么过度频繁的问题
```
![](https://github.com/LockGit/Go/blob/master/img/dict.gif)


## 一键安装(One key Install)
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