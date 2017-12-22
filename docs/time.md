### time
>一.`Sleep` func(d Duration) 休息d时长，阻塞后执行下边
Code:
```
fmt.Println("1")
time.Sleep(time.Second)
fmt.Println("2")
```
OutPut:
```
1
阻塞1秒
2
```

>二.`After` func(d Duration) <-chan time d时长后，不阻塞后执行下边，取出chan阻塞后执行下边
Code:
```
fmt.Println("1")
tc := time.After(time.Second * 2)
fmt.Println("2")
<-tc
fmt.Println("3")
```
OutPut:
```
1
2
阻塞2秒
3
```

>三.`Tick` func(d Duration) <-chan time 每隔d时长，阻塞后执行下边

Code:
```
tc := time.Tick(time.Second * 2)
for {
    <-tc
    fmt.Println("123")
}
```
OutPut:
```
阻塞2秒
123
阻塞2秒
123
```

>四.Time

>1.年(Year)月(Month)日(Day)时(Hour)分(Minute)秒(Second)纳秒(Nanosecond 1秒=1000x1000x1000纳秒)

Code:
```
t := time.Now()
fmt.Println("Year:", t.Year())
fmt.Println("Month:", t.Month())
fmt.Println("Day:", t.Day())
fmt.Println("Hour:", t.Hour())
fmt.Println("Minute:", t.Minute())
fmt.Println("Second:", t.Second())
fmt.Println("Nanosecond:", t.Nanosecond())
```
OutPut:
```
Year: 2017
Month: December
Day: 5
Hour: 17
Minute: 53
Second: 47
Nanosecond: 791425000
```
>2.Parse(layout,value string)(Time,error) 把字符串value转换成(UTC时区)时间格式

Code:
```
fmt.Println(time.Parse("2006/01/02", "2017/12/07"))
```
OutPut:
```
2017-12-07 00:00:00 +0000 UTC <nil>
```
>3.ParseInLocation(layout,value string loc *Location)(Time,error) 把字符串value转换成loc时区时间格式

Code:
```
fmt.Println(time.Parse("2006/01/02", "2017/12/07"))
```
OutPut:
```
2017-12-07 00:00:00 +0000 UTC <nil>
```

> 4.初始化time

```
// 根据年,月,日,时,分,秒,纳秒,时区返回
t := time.Date(2015, 12, 31, 18, 45, 0, 0, time.Now().Location())
fmt.Println(t)
// 根据当前时间返回
t = time.Now()
fmt.Println(t)
// 利用格式化layout和时间字符串，输出时间
t, _ = time.Parse("2006-01-02 15:04:05", "2015-12-31 18:45:00")
fmt.Println(t)
// 返回某个时区的时间
fmt.Println(t.In(time.Now().Location()))
// 利用格式化layout和时间字符串和时区，输出时间
t, _ = time.ParseInLocation("2006-01-02 15:04:05", "2015-12-31 18:45:00", time.Now().Location())
fmt.Println(t)
// 返回离1970-01-01 08:00:00 sec和nsec后的时间
t = time.Unix(1, 999999999)
fmt.Println(t)
// 返回增加duration后的时间
t = t.Add(20)
fmt.Println(t)
// 返回增加years,months,days后的时间
t = t.AddDate(1, 1, 1)
fmt.Println(t)
// 返回增Time local时间
t = t.Local()
fmt.Println(t)
// 返回以duration 四舍五入时间
d, _ := time.ParseDuration("1h")
t = time.Now().Round(d)
fmt.Println(t)
// 返回以duration为单位截取小数位时间
t = time.Now().Truncate(d)
fmt.Println(t)
// 返回以duration为单位截取小数位时间
t = time.Now().UTC()
fmt.Println(t)
```

> 5.比较大小

```
T1 := time.Date(2015, 1, 31, 18, 45, 0, 0, time.UTC)
T2 := time.Date(2015, 2, 1, 2, 45, 0, 0, time.Now().Location())
fmt.Println(T1)
fmt.Println(T2)

// 早于
fmt.Println(T1.Before(T2))
// 晚于
fmt.Println(T1.After(T2))
// 等于
fmt.Println(T1.Equal(T2))
// 是否是1年1月1日(UTC时区)
fmt.Println(T2.IsZero())
```
> 6.格式化

```
// 格式化
fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
```

> 7.编码/解码

```
T1 := time.Now()
T2 := time.Now().Add(time.Hour)
fmt.Println(T1)
fmt.Println(T2)

// gob编码 / 解码
enc, _ := T1.GobEncode()
T2.GobDecode(enc)
fmt.Println(T2)

// Binary编码解码
bina, _ := T1.MarshalBinary()
T2.UnmarshalBinary(bina)
fmt.Println(T2)
// JSONt编码解码
json, _ := T1.MarshalJSON()
T2.UnmarshalJSON(json)
fmt.Println(T2)
// Text编码解码
text, _ := T1.MarshalText()
T2.UnmarshalText(text)
fmt.Println(T2)
```

>五.Duration(持续时长)

>1.`ParseDuration`(s string) (Duration,error) 用s(h,m,s,ms,us,ns)返回时长

Code:
```
fmt.Println(time.ParseDuration("5h"))
fmt.Println(time.ParseDuration("5m"))
fmt.Println(time.ParseDuration("5s"))
fmt.Println(time.ParseDuration("5ms"))
fmt.Println(time.ParseDuration("5us"))
fmt.Println(time.ParseDuration("5ns"))
// 等同于
fmt.Println(time.Hour * 5)
fmt.Println(time.Minute * 5)
fmt.Println(time.Second * 5)
fmt.Println(time.Millisecond * 5)
fmt.Println(time.Microsecond * 5)
fmt.Println(time.Nanosecond * 5)
// Sub两个time之间时长
T1 := time.Date(2016, 1, 31, 18, 45, 0, 0, time.UTC)
T2 := time.Date(2016, 2, 1, 2, 45, 0, 0, time.UTC)
fmt.Println(T1)
fmt.Println(T2)
fmt.Println(T2.Sub(T1))
```
OutPut:
```
5h0m0s <nil>
5m0s <nil>
5s <nil>
5ms <nil>
5µs <nil>
5ns <nil>
8h0m0s
```

>2.`Since`(t time) Duration 从某时间开始到现在时长, 

>`Util`(t time) Duration 从现在到某时间时长

Code:
```
fmt.Println(time.Since(time.Now()))
fmt.Println(time.Until(time.Now()))
```
OutPut:
```
3.443µs
-1.538µs
```
>3 把时长转换为小时(Hours)，分钟(Minutues)，秒(Seconds)，纳秒(Nanoseconds)，string

Code:
```
d, _ := time.ParseDuration("30m")
fmt.Println(d.Hours())
fmt.Println(d.Minutes())
fmt.Println(d.Seconds())
fmt.Println(d.Nanoseconds())
fmt.Println(d.String())
```
OutPut:
```
0.5
30
1800
1800000000000
30m0s
```
>4 func (d Duration) `Round`(m Duration) Duration 将d以m为单位进行`四舍五入`

> func (d Duration) `Truncate`(m Duration) Duration 将d以m为单位`截取小数位`

Code:
```
d, _ := time.ParseDuration("170s")
d2, _ := time.ParseDuration("1m")
fmt.Println(d.Round(d2))
fmt.Println(d.Truncate(d2))
```
OutPut:
```
3m0s
2m0s
```