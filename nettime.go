package main
import( 
	"net"
	"log"
	"io"
	"time"
)


func main(){
	listener ,err := net.Listen("tcp","localhost:8888")
	if err != nil{
		log.Fatal(err)
	}
	for {
		conn,err := listener.Accept()
		if err != nil{
			log.Print(nil)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_ , err := io.WriteString( c , time.Now().Format("15:04:05\n"))
		if err != nil{
			return 
		}
		time.Sleep(time.Second)
	}
}

