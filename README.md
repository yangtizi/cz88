# cz88
纯真88 IP数据库

```golang
package main
 
import (
	"fmt"
	"github.com/yangtizi/cz88"
)
 
func main() {
	fmt.Println(cz88.GetAddress("47.56.100.100"))
}
```


```bash
$go mod init cz88
 
$go build -o cz88.exe
 
$./cz88.exe
 
香港 阿里云
```
