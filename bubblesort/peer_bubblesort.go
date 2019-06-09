//cd %USERPROFILE%\go\src\BubbleSort

package main
import (
	"fmt"
	//"encoding/json"
)

func main() {
	var UserInt [10] int64				//10
	var x int64
	fmt.Printf("Please enter TEN INTEGERES:")
	for i:=0; i<10; i++ {
		fmt.Scan(&x)
		UserInt[i] = x	
	}
	sli := UserInt[0:10]
	BubbleSort(sli)
	fmt.Println("SortedArray:",sli)
}						//21
func BubbleSort(slk []int64) {
	for ii:=0; ii<10; ii++ {
		for iii:=0; iii<9; iii++ {
			if slk[iii] > slk[iii+1] {Swap(slk, iii)}
		}
	}
}

func Swap(slj []int64, j int) {
	t:=(slj)[j]
	(slj)[j] = (slj)[j+1]
	(slj)[j+1] = t
}

