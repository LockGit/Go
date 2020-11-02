# Go  
```
一些小项目，算法，以及有学习价值的问题记录：
```
* 中英文翻译互转 (抓的有道chrome扩展数据包)
* 命令行下的dns信息工具 (golang简单实现类似dig的功能)
* 实现socks5代理来访问看上去不能访问的资源 （网络边界突破，实现类似lcx,ngrok的功能……）
* 两种cache淘汰策略 lru.go (nginx,redis,squid中都有lru的身影）
* 内存对齐在golang中的表现 (证明cpu读取内存是按块读取的)
* golang shell tools (golang实现的交互式shell，实现一个简易版的交互终端)
* golang实现一致性hash (从单节点走向分布式节点)
* etcd watch机制的补充 (微服务)
* 微服务系统中的优雅升级
* 讨论rpc
* 敏感词的检测 (tire树,ac自动机)
* 布隆过滤器 && hash表 && Btree
* golang的map,interface,defer,return,goroutine,slice
* golang的两个nil可能不相等
* golang的gc机制
* leetcode with golang (一些算法的golang实现，algorithm文件夹下)
* diff算法（寻找diff的过程抽象成表示为图的搜索，类似git diff的差异比较，复杂!)

### Chinese and English translation tools in the command line（命令行下中英文翻译工具）
```
dict is a command line tool by go
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

### 实现socks5代理来访问看上去不能访问的资源
```
假设有如下网络情况：
a----->b<------>c

a可以访问b
b不能访问a
b与c之间可以互相访问
实现c访问a上的服务

具体实现：
proxy_server.go 为服务端，proxy_client.go 为客户端
在b上运行：go run proxy_server.go 8000 7999 （其中8000,7999为自定义监听的端口）
在a上运行：go run proxy_client.go b的ip:8000  (其实是a主动连接b的8000端口，为后续的流转发铺垫）
然后c使用b ip的7999端口作为socks5代理即可访问a上的内容
更多场景，可自行探索
```

### 两种cache淘汰策略 lru.go
```
1,如果数据最近被访问过，那么将来被访问的几率也更高（访问后，把当前节点调整到链表头，新加入的Cache直接加到链表头中）
2,如果访问的频率越高，将来被访问的几率更高。（需要一个计数器，计数器排序后，调整链表位置，淘汰无用Cache）
go run lru.go
```

### 内存对齐在golang中的表现
```
思考：声明一个结构体的时候，结构体的内部成员顺序是否会对内存的占用产生影响？
首先了解一下:bool,int8,uint8,int32,uint32,int64,uint64,byte,string 类型占用多少字节。
fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
fmt.Printf("uint8 size: %d\n", unsafe.Sizeof(uint8(0)))
fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
fmt.Printf("uint32 size: %d\n", unsafe.Sizeof(uint32(0)))
fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
fmt.Printf("uint64 size: %d\n", unsafe.Sizeof(uint64(0)))
fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
fmt.Printf("string size: %d\n", unsafe.Sizeof("lock"))
得到如下：
bool size: 1
int8 size: 1
uint8 size: 1
int32 size: 4
uint32 size: 4
int64 size: 8
uint64 size: 8
byte size: 1
string size: 16

那么下面这个结构体理论上共占用 1+4+1+8+1=15 byte
type Test struct {
	a bool  // 1 byte
	b int32 // 4 byte
	c int8  // 1 byte
	d int64 // 8 byte
	e byte  // 1 byte
}

执行如下：
test := Test{}
fmt.Println(unsafe.Sizeof(test))
//output 32 
最终输出为占用32个字节,这与前面所预期的结果完全不一样,这充分地说明了先前的计算方式是错误的!

更改结构体成员顺序：
type Test2 struct {
	e byte  // 1 byte
	c int8  // 1 byte
	a bool  // 1 byte
	b int32 // 4 byte
	d int64 // 8 byte
}

执行：
test2 := Test2{}
fmt.Println(unsafe.Sizeof(test2))
//output 16 
最终输出为占用16个字节！两次输出结果不一样。

```
分析:
CPU读取内存是一块一块读取的，块的大小可以为 2、4、6、8、16 字节等大小。

针对第一个输出32byte的实验：

成员变量 | 类型 | 偏移量 | 自身占用
---|---|---|---
a | bool | 0 | 1
字节对齐 | 无 | 1 | 3
b | int32 | 4 | 4 
c | int8 | 8 | 1
字节对齐 | 无 | 9 | 7
d | int64 | 16 | 8
e | byte | 24 | 1
字节对齐 | 无 | 25 | 7
总占用大小 | - | - | 32

针对第二个输出16byte的实验：

成员变量 | 类型 | 偏移量 | 自身占用
---|---|---|---
e | byte  | 0 | 1
c | int8  | 1 | 1
a | bool  | 2 | 1
字节对齐 | 无 | 3 | 1
b | int32 | 4 | 4
d | int64 | 8 | 8
总占用大小 | - | - | 16

以上过程由于字节对齐原因，
也就是说CPU读取内存是一块一块读取的，而不是一次读取一个offset，所以造成了两次结果不一致。

### golang shell tools
+ [run a golang shell in the command](https://github.com/LockGit/Go/blob/master/shell.go)
```
使用golang实现的shell工具
```


### golang实现一致性hash (从单节点走向分布式节点)
+ ![](https://github.com/LockGit/Go/blob/master/img/hash.png)
```
解决数据倾斜问题，引入虚拟节点，hash环
虚拟节为一个真实节点对应多个虚拟节点。虚拟节点扩充了节点的数量，解决了节点较少的情况下数据容易倾斜的问题。               
```

### etcd watch机制的补充 (微服务)
```
将etcd作为服务发现和配置中心已成为常态，我们可能经常需要在应用中起一个协程watch etcd的某个kv变化,用于尽可能实时的通知当前主进程。
但是在实际应用场景中我们发现经常在协程watch一段时间后，etcd中的value发生了变化，但是应用并没有watch到value的变化。
通过查看etcd源码的watcher interface大致可以得知:
1,正常情况下，起监听的协程如果cotenxt是context.Background/TODO，WatchChan管道不会关闭的，除非etcd集群出了问题不可用。
2,起监听的协程如果cotenxt使用了cetcd.WithRequireLeader包装的上下文,这个时候etcd的WatchChan管道会因为etcd集群leader的变更而导致WatchChan通道关闭，那么此时的watch机制自然失效。
理论上在接收到WatchChan通道关闭信号的时候，应用程序应该再次调用watch,也就是用for{}包装下，我们微服务系统又引入了go-micro的Config包，有二次修改搞了一些自定义的东西，粗略观察是没有这种for{}包装的
鉴于此，又由于网络或者磁盘IO等原因，我们自己应用中WatchChan管道的关闭都会导致watch失效,所以我们可以在微服务启动后，起一个goroutine定时轮询etcd,拿到最新的kv。
在WatchChan管道关闭watch失效时，将定时轮询得到的最新kv作为watch机制失效的补充。（定时不用太频繁，我设置的是1分钟，一些配置的1分钟误差可以接受，对于etcd集群来说也不会有太大的压力）

还有一种watch优化机制是在应用和etcd集群之间增加一个grpc代理，但是这个grpc代理会有cache，弄不好就会出问题，使用时也需小心。
```
+ ![](https://github.com/LockGit/Go/blob/master/img/etcd-watch.png)


### 敏感词的检测 (tire树,ac自动机)
```
这个在我的Py项目中也有介绍，这是风控系统的一层基础设施，简而言之,trie树利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的，带上AC自动机版的
会更强劲。在实际应应用中，我们将敏感词微服务化，用于粗略过滤一些简单的敏感信息，敏感词的词库在不断增加，应用中会起一个go协程增量更新新增的词。但更高级的
可能会基于机器学习来做，现实中，tire树版基本够用。
```

### 布隆过滤器 && hash表 && Btree
```
判断元素是否存在某个集合中：
hash表用于定位一个元素是否存在的复杂可以认为是O(1),主要取决去hash算法，如果出现hash冲突，一般会在对应hash值后面挂上一个链表，这里有一个问题，如果数据
足够多，那么hash表也会变的越来越大，hash冲突的可能性也会变大，一些hash算法本身可能会有rehash的过程来扩容hash表的大小。对于少量的数据，hash表是一个不错
的选择。

布隆过滤器的也可以用于判断一个元素是否存在与一个集合中，布隆过滤器与hash表有着本质数据结构区别,布隆过滤器会存在一定的误差，主要原因为hash冲突，所以选择合适的hash函数尤为重要。
网上有bloom filter的golang版实现，https://github.com/willf/bloom，这个库采用了murmurhash,murmurhash实现了尤其对大块的数据，具有较高的平衡性与低碰撞率，Redis在实现字典时也借鉴了这种方法。
布隆过滤器是一个很长的二进制向量和一系列随机映射函数，对于判断一个元素是否存在，布隆过滤器会先对key执行多次不同hash函数得到不同的hash值,在分别查找对应hash值的标志位。全部满足，则元素存在。由于hash冲突的原因，所以布隆过滤器说某个集合中不存在某个元素，那么一定不存在，说某个元素在,可能会被误判。
高版本的redis还支持bloom filter扩展，启动redis-server前加载对应so文件，则redis便很快可以提供一个布隆过滤器的功能，而相关细节均被屏蔽，对于使用者而言只需要设置下相关参数，误差率即可。

Btree，二叉树，这种树形结构应用也很广泛，比如mysql索引就采用了这种数据结构，而数据库本身有很多范围查找的场景，使用这种结构就尤其符合。与之相关的还有红黑树，多用于内存排序。
这些树或者树的变种均为解决数据平衡，提高查找效率而演变出来的结构。
```

### golang的map,interface,defer,return,goroutine,slice
```
1,结构体可以比较，但是包含map的结构体不可比较。
2,当程序运行到defer函数时，不会执行函数实现，但会将defer函数中的参数代码进行执行。
3,使用 i.(type) ,i必须是interface{}
4,map需要初始化,最好用make([]int,x)完成初始化
5,defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
6,真正实现并行至少需要2个操作系统硬件线程并至少有两个goroutine时才能实现真正的并行,每个goroutine在一个单独的系统线程上执行指令。
7,slice存在对array的引用，对slice的修改在slice未扩容前会影响到array，当slice发生扩容行为之后，这时候内部就会重新申请一块内存空间，
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

```golang的两个nil可能不相等
接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V。一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）。

两个接口值比较时，会先比较 T，再比较 V。
接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较。
func main() {
	var p *int = nil
	var i interface{} = p
	fmt.Println(i == p) // true
	fmt.Println(p == nil) // true
	fmt.Println(i == nil) // false
}
上面这个例子中，将一个 nil 非接口值 p 赋值给接口 i，此时，i 的内部字段为(T=*int, V=nil)，i 与 p 作比较时，将 p 转换为接口后再比较，因此 i == p，p 与 nil 比较，直接比较值，所以 p == nil。

但是当 i 与 nil 比较时，会将 nil 转换为接口 (T=nil, V=nil)，与i (T=*int, V=nil) 不相等，因此 i != nil。因此 V 为 nil ，但 T 不为 nil 的接口不等于 nil。
```


```golang的gc机制
最常见的垃圾回收算法有标记清除(Mark-Sweep) 和引用计数(Reference Count)，Go 语言采用的是标记清除算法。并在此基础上使用了三色标记法和写屏障技术，提高了效率。
标记清除收集器是跟踪式垃圾收集器，其执行过程可以分成标记（Mark）和清除（Sweep）两个阶段：

标记阶段 — 从根对象出发查找并标记堆中所有存活的对象；
清除阶段 — 遍历堆中的全部对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表。
标记清除算法的一大问题是在标记期间，需要暂停程序（Stop the world，STW），标记结束之后，用户程序才可以继续执行。为了能够异步执行，减少 STW 的时间，Go 语言采用了三色标记法。

三色标记算法将程序中的对象分成白色、黑色和灰色三类。

白色：不确定对象。
灰色：存活对象，子对象待处理。
黑色：存活对象。
标记开始时，所有对象加入白色集合（这一步需 STW ）。首先将根对象标记为灰色，加入灰色集合，垃圾搜集器取出一个灰色对象，将其标记为黑色，并将其指向的对象标记为灰色，加入灰色集合。重复这个过程，直到灰色集合为空为止，标记阶段结束。那么白色对象即可需要清理的对象，而黑色对象均为根可达的对象，不能被清理。

三色标记法因为多了一个白色的状态来存放不确定对象，所以后续的标记阶段可以并发地执行。当然并发执行的代价是可能会造成一些遗漏，因为那些早先被标记为黑色的对象可能目前已经是不可达的了。所以三色标记法是一个 false negative（假阴性）的算法。

三色标记法并发执行仍存在一个问题，即在 GC 过程中，对象指针发生了改变。比如下面的例子：

1
A (黑) -> B (灰) -> C (白) -> D (白)
正常情况下，D 对象最终会被标记为黑色，不应被回收。但在标记和用户程序并发执行过程中，用户程序删除了 C 对 D 的引用，而 A 获得了 D 的引用。标记继续进行，D 就没有机会被标记为黑色了（A 已经处理过，这一轮不会再被处理）。

1
2
3
A (黑) -> B (灰) -> C (白) 
  ↓
 D (白)
为了解决这个问题，Go 使用了内存屏障技术，它是在用户程序读取对象、创建新对象以及更新对象指针时执行的一段代码，类似于一个钩子。垃圾收集器使用了写屏障（Write Barrier）技术，当对象新增或更新时，会将其着色为灰色。这样即使与用户程序并发执行，对象的引用发生改变时，垃圾收集器也能正确处理了。

一次完整的 GC 分为四个阶段：

1）标记准备(Mark Setup，需 STW)，打开写屏障(Write Barrier)
2）使用三色标记法标记（Marking, 并发）
3）标记结束(Mark Termination，需 STW)，关闭写屏障。
4）清理(Sweeping, 并发)
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
