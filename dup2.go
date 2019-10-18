package main
import(
	"bufio"
	"fmt"
	"os"
)

func main(){
	cnt := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0{
		countlines(os.Stdin,cnt)
	} else {
		for _, arg := range file{
			f , err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2: 5v \n ",err)
				continue
			}
			countlines(f ,cnt)
			f.Close()
		}
	}
	for line , n := range cnt{
		if n > 1{
			fmt.Printf("%d \t %s \n", n , line)
		}
	}
}

func countlines(f *os.File ,cnt map[string]int ){
	input := bufio.NewScanner(f)
	for input.Scan(){
		cnt[input.Text()]++
	}
}
