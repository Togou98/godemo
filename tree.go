package main
import "fmt"

type node struct{
		data int
		left , right *node
}
type tree struct{
		root *node
}

func insert(root *node ,data int){
		i  := new(node)
		i.data = data ; i.left =nil ; i.right = nil;
		if root == nil{
				root = i
				return
		}else{
				t := root
				for t != nil {
						if data < t.data {
								if t.left == nil{
										t.left = i
										return
								}else{ t = t.left}
						}else {
								if t.right == nil{
										t.right = i
										return 
								}else{ t = t.right}
						}
				}
		}
}

func (root *node)inorder(){
		if root == nil {
				return
		}
		root.left.inorder()
		fmt.Printf("%d ",root.data)
		root.right.inorder()
}
func main(){
		t := tree{root:nil}
		arr := []int{9,7,3,6,1}
		for _,n := range arr{
				insert(t.root , n)
		}
		t.root.inorder()

}

