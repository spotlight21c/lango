# lango
Simple translation package by golang


# How to use

en.json

```js
{
    "hello":"Hello",
    "how are you":"How are you?",
    "my name is":"My name is %s"
}
```

ko.json

```js
{
    "hello":"안녕하세요"
}
```

```golang
import "github.com/spotlight21c/lango"

lango.Init("path/to/locale", "en")
lango.SetLocale("ko")

fmt.Println(lango.Get("hello"))
// 안녕하세요

fmt.Println(lango.Get("how are you"))
// How are you?

fmt.Println(lango.Get("nice to meet you"))
// nice to meet you

fmt.Println(lango.Get("my name is", "John"))
// My name is John
```
