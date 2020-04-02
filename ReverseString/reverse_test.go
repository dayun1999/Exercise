package ReverseString

import(
	"testing"
)

func TestReverseString1(t *testing.T) {
	//name := "郭德纲" //错误
	name := "王大云" //eman
	result := reverseString1(name)
	t.Logf("输出1结果是：%s\n",result)

}