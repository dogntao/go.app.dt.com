#### 字符串包含Contains;
1.字符串s是否包含子串substr
> func `Contains`(s, substr string) bool

Code:
```
fmt.Println(strings.Contains("test","te"))
fmt.Println(strings.Contains("test","et"))
fmt.Println(strings.Contains("test",""))
fmt.Println(strings.Contains("",""))
```
Output:
```
true
false
true
true
```
2.字符串s是否包含chars中任何一个字符
> func `ContainsAny`(s, chars string) bool

Code:
```
fmt.Println(strings.ContainsAny("test","i"))
fmt.Println(strings.ContainsAny("test","t"))
fmt.Println(strings.ContainsAny("test","t & i"))
fmt.Println(strings.ContainsAny("test",""))
fmt.Println(strings.ContainsAny("",""))
```
Output:
```
false
true
true
false
false
```

#### 字符串长度Count;
1.计算字符串s中substr长度
> func `Count`(s, substr string) int

> `如果substr是空字符串，返回s长度+1`

Code:
```
fmt.Println(strings.Count("testee","e"))
fmt.Println(strings.Count("test",""))
```
Output:
```
3
5
```

#### 空格拆分Fields;
1.用空格拆分字符串s
> func `Fields`(s string) []string 

Code:
```
fmt.Println(strings.Fields("  this is test   "))
```
Output:
```
[this is test]
```
#### 前后缀
1.字符串s，是否以prefix作为前缀
> func `HasPrefix`(s, prefix string) bool

Code:
```
fmt.Println(strings.HasPrefix("test","t"))
fmt.Println(strings.HasPrefix("Test","t"))
fmt.Println(strings.HasPrefix("test","T"))
fmt.Println(strings.HasPrefix("test",""))
```
Output:
```
true
false
false
true
```
2.字符串s，是否以suffix作为后缀
> func `HasSuffix`(s, suffix string) bool

Code:
```
fmt.Println(strings.HasPrefix("test","t"))
fmt.Println(strings.HasPrefix("Test","t"))
fmt.Println(strings.HasPrefix("test","T"))
fmt.Println(strings.HasPrefix("test",""))
```
Output:
```
true
false
false
true
```

#### 出现位置
1.字符串subStr,在s中第一次出现的位置
> func `Index`(s, subStr string) int

> 如果不包含返回-1

Code:
```
fmt.Println(strings.Index("test", "s"))
fmt.Println(strings.Index("test", "o"))
```
Output:
```
2
-1
```

2.字符串subStr,在s中最后一次出现的位置
> func `LastIndex`(s, subStr string) int

> 如果不包含返回-1

Code:
```
fmt.Println(strings.LastIndex("test", "t"))
fmt.Println(strings.LastIndex("test", "o"))
```
Output:
```
3
-1
```

#### 拼接/拆分
1.把数组a,用sep拼接成字符串
> func `Join`(a []sting, sep string) string

Code:
```
a := []string{"this", "is", "test"}
fmt.Println(strings.Join(a, ","))
```
Output:
```
this,is,test
```

2.把字符串s,用sep拆分成数组
> func `Split`(s, sep string) []string

Code:
```
fmt.Println(strings.Split("this,is,test", ","))
```
Output:
```
[this is test]
```

3.Repeat: 将count个s连接成一个新的字符串
> func `Repeat`(s string, count int) string 

Code:
```
fmt.Println(strings.Repeat("hello", 3))
```
Output:
```
hellohellohello
```

#### 大小写转换
1.首字母大写
> func `Title`(s string) string
Code:
```
fmt.Println(strings.Title("this is test"))
```
Output:
```
This Is Test
```
2.所有字母大写
> 2.1 func `Totitle`(s string) string

Code:
```
fmt.Println(strings.ToTitle("this is test"))
```
Output:
```
THIS IS TEST
```
> 2.2 func `ToUpper`(s string) string

Code:
```
fmt.Println(strings.ToUpper("this is test"))
```
Output:
```
THIS IS TEST
```
3.所有字母小写
> func `ToLower`(s string) string
Code:
```
fmt.Println(strings.ToLower("This IS tEst"))
```
Output:
```
this is test
```

#### 替换
1.Replace:替换s中old为new指定n个(如果n为-1，全部替换)
> func `Replace`(s,old,new string,n int)string

Code:
```
s := "this is test field"
fmt.Println(strings.Replace(s, "s", "t", -1))
```
Output:
```
thit it tett field
```

2.Map:利用func替换s
> func `Map`(mapping func(r rune) rune,s string) string

Code:
```
r := func(r rune) rune {
		if r == 's' {
			return 't'
		}
		return r
	}
s := "this is test field"
fmt.Println(strings.Map(r, s))
```
Output:
```
thit it tett field
```
3.Replacer:根据替换列表进行替换
Code:
```
s := "hello world,hello1 world1,hello2 world2"
r := strings.NewReplacer("hello", "你好", "world", "世界")
rst := r.Replace(s)
fmt.Println(rst)
```
Output:
```
你好 世界,你好1 世界1,你好2 世界2
```

#### 过滤
1. Trim 删除s首尾连续的cutset字符
> func Trim(s string, cutset string) string

Code:
```
fmt.Println(strings.Trim("   hello world    ", " "))
```
Output:
```
hello world
```
2. TrimLeft 删除s字符串左边连续的cutset字符
> func TrimLeft(s string, cutset string) string

Code:
```
fmt.Printf("%s%s", strings.TrimLeft("   hello world    ", " "), "test")
```
Output:
```
hello world    test
```
3. TrimRigt 删除s字符串右边连续的cutset字符
> func TrimRight(s string, cutset string) string

Code:
```
fmt.Printf("%s%s", strings.TrimRight("   hello world    ", " "), "test")
```
Output:
```
   hello worldtest
```
4. TrimFunc 删除s字符串连续的f(rue)的字符，同理TrimLeftFunc,TrimRightFunc
> func TrimFunc(s string, f func(r rune) bool) string

Code:
```
f := func(r rune) bool {
    return r == ' '
}
fmt.Printf("%s%s", strings.TrimFunc("   hello world    ", f), "test")
```
Output:
```
hello worldtest
```
5. TrimPrefix 删除s字符串以prefix开头的字符,同理TrimSuffix
> func TrimPrefix(s, prefix string) string

Code:
```
fmt.Printf("%s%s", strings.TrimPrefix("hello world", "h"), "test")
```
Output:
```
ello world    test
```
6. TrimSpace 删除s字符串首尾连续的空白字符
> func TrimSpace(s string) string

Code:
```
fmt.Printf("%s%s", strings.TrimSpace("   hello world   "), "test")
```
Output:
```
hello worldtest
```

#### 读取字符串
Code:
```
s := "123abc"
reader := strings.NewReader(s)
// offset:偏移量，负数表示反向移动
// whence:从哪里开始移动，0:起始位置;1:当前位置;2:结尾位置
reader.Seek(1, 0)
reader.Seek(1, 1)
rLen := reader.Len()
fmt.Println(rLen)
```
Output:
```
4
```

#### 截取字符串
> 将字符串转换成rune数组，然后获取rune数据，然后转换成字符串
Code:
```
s := "中国test中国"
rs := []rune(s)
fmt.Println(string(rs[3:]))
```
Output:
```
est中国
```

#### 字符串长度
>1.`strings.Count`,算出包含空的，减去1

Code:
```
s := "董涛123abc"
sLen := strings.Count(s, "")
fmt.Println(sLen - 1)
```
Output:
```
8
```

>2.利用[]rune,后调用len

Code:
```
s := "董涛123abc"
fmt.Println(len([]rune(s)))
```
Output:
```
8
```
>3.利用utf8.RuneCountInString,直接算出

Code
```
s := "董涛123abc"
fmt.Println(utf8.RuneCountInString(s))
```
Output:
```
8
```
