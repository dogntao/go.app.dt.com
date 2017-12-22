### unicode
#### 一、判断类型
```
testStr := "tA \b5Ὂg̀-9,! ℃ᾭG2<董"
for _, v := range testStr {
    fmt.Printf("For %q:\n", v)
    // 控制符 例:\n
    if unicode.IsControl(v) {
        fmt.Println("\t is control rune")
    }
    // 十进制数组 例:1
    if unicode.IsDigit(v) {
        fmt.Println("\t is digit rune")
    }
    // 图形字符 例:字母、数字、标记、符号、空格
    if unicode.IsGraphic(v) {
        fmt.Println("\t is graphic rune")
    }
    // 字母 例:a董(包含汉子)
    if unicode.IsLetter(v) {
        fmt.Println("\t is letter rune")
    }
    // 是否mark(掩码) 例: g̀
    if unicode.IsMark(v) {
        fmt.Println("\t is mark rune")
    }
    // 是否数字 例: 123１２３(不区分全角和半角)
    if unicode.IsNumber(v) {
        fmt.Println("\t is number rune")
    }
    // 是否某种类型 例: 汉字
    rangeTable := []*unicode.RangeTable{unicode.Han}
    if unicode.IsOneOf(rangeTable, v) {
        fmt.Println("\t is 汉字(IsOneOf) rune")
    }
    if unicode.Is(unicode.Han, v) {
        fmt.Println("\t is 汉字(Is) rune")
    }
    if unicode.In(v, unicode.Han, unicode.Hangul) {
        fmt.Println("\t is 汉字(In) rune")
    }
    // 是否可打印字符 例: \t不是
    if unicode.IsPrint(v) {
        fmt.Println("\t is print rune")
    }
    // 是否标点 例:,
    if unicode.IsPunct(v) {
        fmt.Println("\t is punct rune")
    }
    // 是否空白: 例: \t,\n,\v,\f,\r
    if unicode.IsSpace(v) {
        fmt.Println("\t is space rune")
    }
    // 是否符号: 例: <>
    if unicode.IsSymbol(v) {
        fmt.Println("\t is symbol rune")
    }
    // 是否Unicode特殊字符: 例: ᾭ
    if unicode.IsTitle(v) {
        fmt.Println("\t is title rune")
    }
    // 是否大写: 例: A
    if unicode.IsUpper(v) {
        fmt.Println("\t is upper rune")
    }
}
```
#### 二、数据转换
>1.将r转换成小写 `ToLower`(r rune) rune 同理`ToUpper`,`ToTitle`
Code:
```
fmt.Println(string(unicode.ToLower('G')))
```
Output:
```
g
```
>2.转换r `To`(case int, r rune) rune
Code:
```
fmt.Println(string(unicode.To(unicode.UpperCase, 'g')))
fmt.Println(string(unicode.To(unicode.LowerCase, 'G')))
fmt.Println(string(unicode.To(unicode.TitleCase, 'g')))
```
Output:
```
G
g
G
```