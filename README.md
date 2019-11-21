# ditime

![travisCI](https://secure.travis-ci.org/digital-idea/ditime.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/digital-idea/ditime)](https://goreportcard.com/report/github.com/digital-idea/ditime)

업무툴 제작에 사용되는 시간관련 Go 라이브러리이다.

### Use in go

```go
package main

import (
    "fmt"
    "github.com/digital-idea/ditime"
)

func main() {
    d := "2019年1月1日"
    t, _ := ditime.ToFullTime(10, d)
    fmt.Println(t)
    // 2019-01-01T10:00:00+09:00
}
```

### License: BSD 3-Clause License