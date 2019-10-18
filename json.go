package main
import(
	"fmt"
	"encoding/json"
	//"os"
)
type res struct {
		Page int
		Fruit []string
}

type response struct{
		Page int `json:"page"`
		Fruit []string `json:"fruit"`
}
func main(){
		bolB, _ :=json.Marshal(true)
		fmt.Println(string(bolB))
		
		intB,_ := json.Marshal(100)
		fmt.Println(string(intB))

		res1D  := &res{
				Page : 1,
				Fruit : []string{"apple","banana","pear"}}
		res1B ,_:= json.Marshal(res1D)
		fmt.Println(string(res1B))
}

