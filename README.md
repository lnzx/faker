# faker

## useragent

从这个网站取的数据：useragentstring.com, 爬取数据脚本`ua.js`

数据更新日期: **2023/04/11**

## usage

```go
import "github.com/lnzx/faker/useragent"

func main() {
	ua := useragent.New()
	fmt.Println(ua.Chrome())
	fmt.Println(ua.Firefox())
	fmt.Println(ua.Edge())
	fmt.Println(ua.Random())
}
```



