package main
import(
	"log"
	"os"
	// "strings"
)
func main(){
	arg:=os.Args
	str:=""
	for i:=0; i < len(arg[1:]); i++ {
		str=str+" "+arg[i+1]
	}
	log.Println(str)
}