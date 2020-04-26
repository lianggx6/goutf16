# goutf16
* This library provides the method encode string to uint16 array and the method decode uint16 array to string. What's more, this library includes some simple funtion like ```Count```,```Index```,```Join```of uint16 array.
* The encode and decode method is efficient. And the process just need one memory allocation.
## Encode
* encode string to uint16 array like this.
```go
package main

import (
	"fmt"
	"github.com/lianggx6/goutf16"
)

func main() {
	content := goutf16.EncodeStringToUTF16("你好")
	fmt.Println(content)
}
```

## Decode
* decode uint16 array to string like this
```go
package main

import (
	"fmt"
	"github.com/lianggx6/goutf16"
)

func main() {
	content := goutf16.DecodeUTF16ToString([]uint16{20320, 22909})
	fmt.Println(content)
}
```
