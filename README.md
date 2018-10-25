# Go  

### Chinese and English translation tools in the command line（命令行下中英文翻译工具）
```
据说内置的有道翻译的Key因为调用量过度频繁导致被封了？
所以我就直接抓了Chrome浏览器的有道插件的数据包，请求这个api应该没有什么过度频繁的问题
```
![](https://github.com/LockGit/Go/blob/master/img/dict.gif)

## 一键安装(OneKey Install)
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


### leetcode with golang 
```
cat *.go |grep Time -A 1 |grep -v Time
```
![](https://github.com/LockGit/Go/blob/master/img/leetcode.gif)

### http-proxy.go 使用go实现一个http代理
```
执行：go run http-proxy.go -h 查看帮助

执行：go run http-proxy.go --debug=true  打开调式模式

默认监听：8080 端口，把浏览器的代理设置成127.0.0.1 8080 端口，那么浏览器所访问资源将会走go代理脚本
```
![](https://github.com/LockGit/Go/blob/master/img/http-proxy.png)
```
* 如果是https类型，CONNECT方法 （要求http协议>=1.1)
    * 客户端通过CONNECT方法请求代理服务器创建一条到达任意目的服务器和端口的TCP链接，代理服务器仅对客户端和服务器之间的后续数据进行盲转发（只是转发，不关心，也不懂发送的内容是什么）。
* 如果是http类型，直接代理转发

代理https时的大致过程如下：
1）客户端通过HTTP协议发送一条CONNECT方法的请求给代理服务器，告知代理服务器需要连接的主机和端口。
CONNECT www.xxx.com:443 HTTP/1.1
User-agent: Mozilla/5.0

2）代理服务器一旦建立了和目标主机（上例的www.xxx.com:443）TCP连接，就会回送一条HTTP 200 Connection Established应答给客户端。
example: HTTP/1.1 200 Connection Established

3）此时隧道就建立起来了。客户端通过该HTTP隧道发送的所有数据都会被代理服务器（通过之前建立起来的与目标主机的TCP连接)原封不动的转发给目标服务器。目标服务器发送的所有数据也会被代理服务器原封不动的转发给客户端。

4) 后续可完善认证体系，在vps上部署代理
```



### diff.go 实现类似git中diff功能
```
执行: go run diff.go 1.md 2.md

```
![](https://github.com/LockGit/Go/blob/master/img/go-diff.png)

```
关于算法：
Myers算法，原作者1953年生于美国Idaho州，University of California at Berkeley教授，
美国科学院工程部成员，贡献了模式匹配和计算生物学的基本算法。
算法将寻找diff的过程，抽象成表示为图的搜索。

在网上找了一些关于该算法文档描述如下：
以两个字符串，src=ABCABBA，dst=CBABAC为例，根据这两个字符串我们可以构造下面一张图，横轴是src内容，纵轴是dst内容。
那么，图中每一条从左上角到右下角的路径，都表示一个diff。向右表示“删除”，向下表示”新增“，对角线则表示“原内容保持不动“。
图：
```
![](https://github.com/LockGit/Go/blob/master/img/diff.png)

```
最优路径如红线：
```
![](https://github.com/LockGit/Go/blob/master/img/diff2.png)

```
首先，定义参数d和k，d代表路径的长度，k代表当前坐标x - y的值。定义一个”最优坐标“的概念，
最优坐标表示d和k值固定的情况下，x值最大的坐标。x大，表示向右走的多，表示优先删除。
还是用上面那张图为例。我们从坐标(0, 0)开始，此时，d=0，k=0，然后逐步增加d，计算每个k值下对应的最优坐标。
因为每一步要么向右（x + 1），要么向下（y + 1），要么就是对角线（x和y都+1)，所以，当d=1时，k只可能有两个取值，要么是1，要么是-1。
当d=1，k=1时，最优坐标是(1, 0)。
当d=1，k=-1时，最优坐标是(0, 1)。
因为d=1时，k要么是1，要么是-1，当d=2时，表示在d=1的基础上再走一步，k只有三个可能的取值，分别是-2，0，2。
当d=2，k=-2时，最优坐标是(2, 4)。
当d=2，k=0时，最优坐标是(2, 2)。
当d=2，k=2时，最优坐标是(3, 1)。
以此类推，直到我们找到一个d和k值，达到最终的目标坐标(7, 6)。

算法过程：
迭代d，d的取值范围为0到n+m，其中n和m分别代表源文本和目标文本的长度（这里我们选择以行为单位）
每个d内部，迭代k，k的取值范围为-d到d，以2为步长，也就是-d，-d + 2，-d + 2 + 2…
使用一个数组v，以k值为索引，存储最优坐标的x值（这里使用hash也行，但是用数组效率更高一些，因为Go不支持使用负数做索引，所以需要创建一个自定义类型）
将每个d对应的v数组存储起来，后面回溯的时候需要用
当我们找到一个d和k，到达目标坐标(n, m)时就跳出循环
使用上面存储的v数组（每个d对应一个这样的数组），从终点反向得出路径

git真正用的是标准Myers算法的一个变体，标准的算法空间消耗很大。
在某些情况下，变体产生的diff会和标准算法有所不同。
也就是说，如果你按照上面的算法实现的程序，出来的结果和git diff的结果有所不同是正常的。
```