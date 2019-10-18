package main
import(
    "fmt"
    "crypto/sha256"
)

func main(){
	s1arg := []byte("y")
	s2arg := []byte("y")
	s1 := sha256.Sum256(s1arg)
	s2 := sha256.Sum256(s2arg)
	cnt := 0
	for i := 0; i < 32 ; i++{
		if(s1[i]!= s2[i]){
			cnt++
		}
	}
	fmt.Printf("s1 : %x\ns2 : %x\nthere have %d bits differences",s1,s2,cnt)
}