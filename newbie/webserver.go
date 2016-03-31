// Web server
package main
import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/togoogle",redirectToGoogle)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))	
}

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hello world in webserver")
	fmt.Fprint(w,"URL.Path = %q\n",r.URL.Path)	
}

func redirectToGoogle(w http.ResponseWriter,r *http.Request){
	http.Redirect(w,r,"http://google.com",301)
}