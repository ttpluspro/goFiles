//练习 3.13： 编写KB、MB的常量声明，然后扩展到YB
package main

import "fmt"

const (
	_   = 1 << (10 * iota) //iota = 0
	KiB                    // 1024 iota = 1
	MiB                    // 1048576 iota =2
	GiB                    // 1073741824
	TiB                    // 1099511627776             (exceeds 1 << 32)
	PiB                    // 1125899906842624
	EiB                    // 1152921504606846976
	ZiB                    // 1180591620717411303424    (exceeds 1 << 64)
	YiB                    // 1208925819614629174706176
)

func main() {
	fmt.Printf("%b\t%d\n", KiB, KiB)
}

// 1 << 0 = 1
// 1 << 10 == 10000000000
//1 << 20 == 100000000000000000000

// 所谓左移符号，就是这样的 最后一位的数字 （当然是1） 移动的位数，然后没有的地方用0补充。

//也就是说 ：