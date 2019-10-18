package main
import(
		"fmt"
		"os"
		"io/ioutil"
		"path/filepath"
)
func main(){
		
		f ,_ := ioutil.TempFile("","sample")
	
		fmt.Println("Temp file name :",f.Name())
		defer os.Remove(f.Name())

		_,_ = f.Write([]byte{1 , 2, 3, 4})
		
		dname , _ := ioutil.TempDir("","sampledir")
		fmt.Println("Temp dir name:" ,dname)
		defer os.RemoveAll(dname)

		fname := filepath.Join(dname , "file1")
		_ = ioutil.WriteFile(fname , []byte{1,2},0666)


}

