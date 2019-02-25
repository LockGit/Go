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

### 命令行下的dns信息工具
#### 一键安装(OneKey Install)
```
wget https://raw.githubusercontent.com/LockGit/Go/master/fdns/FindDnsRecordFast.go && go build -o /usr/local/bin/fdns FindDnsRecordFast.go 
```

#### 使用(Usage)
![](https://github.com/LockGit/Go/blob/master/img/fdns.gif)


### 两种cache淘汰策略 lru.go
```
1,如果数据最近被访问过，那么将来被访问的几率也更高（访问后，把当前节点调整到链表头，新加入的Cache直接加到链表头中）
2,如果访问的频率越高，将来被访问的几率更高。（需要一个计数器，计数器排序后，调整链表位置，淘汰无用Cache）
go run lru.go
```

### golang shell tools
+ [run a golang shell in the command](https://github.com/LockGit/Go/blob/master/shell.go)
```
使用golang实现的shell工具
```

### 这些问题都不是问题
```
1,结构体可以比较，但是包含map的结构体不可比较。
2,当程序运行到defer函数时，不会执行函数实现，但会将defer函数中的参数代码进行执行。
3,使用 i.(type) ,i必须是interface{}
4,map需要初始化,最好用make([]int,x)完成初始化
5,interface {}(*int) nil != nil 
6,defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
7,真正实现并行至少需要2个操作系统硬件线程并至少有两个goroutine时才能实现真正的并行,每个goroutine在一个单独的系统线程上执行指令。
8,slice存在对array的引用，对slice的修改在slice未扩容前会影响到array，当slice发生扩容行为之后，这时候内部就会重新申请一块内存空间，
将原本的元素拷贝一份到新的内存空间上。此时其与原本的数组就没有任何关联关系了，再进行修改值也不会变动到原始数组。
```
```go
func caseOne() *int {
	var i int
	defer func(i *int) {
		time.Sleep(time.Second * 2)
		*i = 3
	}(&i)
	return &i
}

func caseTwo() int {
	var i int
	defer func(i int) {
		time.Sleep(time.Second * 2)
		i = 3
	}(i)
	return i
}
func main() {
	fmt.Println(*caseOne())
	fmt.Println(caseTwo())
}
```

### leetcode with golang 
```
cat *.go |grep Time -A 1 |grep -v Time
 * 给定两个二进制字符串，返回它们的和（也是一个二进制字符串）。 例如， a = "11" b = "1" 返回"100"。
--
 * 将大于10的数的各个位上的数字相加，若结果还大于0的话，则继续相加，直到数字小于10为止
--
 * 将获得两个非空链接列表，表示两个非负整数。 数字以相反的顺序存储，并且它们的每个节点包含单个数字。 添加两个数字并将其作为链接列表返回。
--
 * 找出字符串中所有的变位词,重复出现的
--
 * 高度平衡二叉树判断
--
 * 中序遍历二叉树
--
 * 反转二叉树
--
 * 层序遍历二叉树
--
 * 在二叉树中找一条路径，使得该路径的和最大。该路径可以从二叉树任何结点开始，也可以到任何结点结束。
--
 * 返回所有根到叶节点的路径
--
 * 二叉树的后序遍历
--
 * 二叉树前序遍历
--
 * 二叉树锯齿形层次遍历
--
 * 等价于 求 m 与 n 二进制编码中 同为1的前缀. [5,7] , 5 & 6 & 7 = 4
--
 * 爬N阶的楼梯，每次可以爬1阶或者2阶，一共有多少种爬法? (递归，动态规划)
--
 * 回溯算法|递归实现
--
 * 给一个数组，其中数组在下标i处的值为A[i]，坐标(i,A[i])和坐标(i,0)构成一条垂直于坐标轴x的直线。现任取两条垂线和x轴组成四边形容器。问其中盛水量最大为多少？
--
 * 求小于给定非负数n的质数个数
--
 * git diff 算法
--
 * 给定字符串T和S，求S的子串中有多少与T相等。
--
 * 动态规划
--
 * 求一个整数的阶乘末尾有多少个0
--
* 从一个未经排序的数组中找出第k大的元素。注意是排序之后的第k大，而非第k个不重复的元素。
--
 * 在旋转排序数组中搜索(指定的目标)
--
 * 二分查找
--
 * 给定一个数组[a,b,c,d,e],一个target值，找出满足a+b+c+d==target的不重复集合
--
 * 无向图复制 DFS,BFS
--
 * 群组错位词分类
--
 * 实现一个http代理
--
 * 实现字符串位置查找
--
 *  同构字符串
--
 * 跳跃游戏---动态规划
--
 * 求一个字符串最后一个字符的长度
--
 * 字符串最长公共前缀
--
 * 从一个数组中找出出现半数以上的元素
--
 * 动态规划---最大面积
--
 * 求连续的数组值，加和最大
--
* 寻找连续数组中缺失的数
--
 * 将数组中的0移动数组末尾
--
 * 海明重量---位操作
--
 * 实现一个正则匹配算法
--
 * 一个有序数组，返回去除重复元素后数组的长度。不允许申请新的空间。
--
 * 给定一个数组和一个值，在这个数组中移除指定值和返回移除后新的数组长度。不要为其他数组分配额外空间,O(1)复杂度
--
 * golang反转链表
--
 * 反转数组
--
 * 一个list中只有一个数出现了1次，其他数2次，找出这个出现一次的数字
--
 * 在一个数组中，找到3个元数之和等于0的元素列表组合
--
 * 给一个整数数组，找到三个数的和与给定target的值距离最短的那个和
--
 * 两数之和
--
 * 是否是丑陋数，所谓丑陋数就是其质数因子只能是2,3,5。
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
## 使用(Usage)


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