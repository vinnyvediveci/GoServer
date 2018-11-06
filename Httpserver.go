package main
// importing  relevant packages (format and net/http
import(
	"fmt"
	"net/http"

)
//HandleFunc - First parameter takes path, second par. takes function to execute

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
// prints welcome message to website.
		fmt.Fprintf(w, "Welcome to my website!")
	})
//http.FileServer - serves static assets, htp.Dir - shows fileserver where to serve from
	fs := http.FileServer(http.Dir("static/"))
//in order to serve the file correctly, use Stripprefix to strip away directory file lives in 
	http.Handle("/static/", http.StripPrefix("static", fs))

//listen on port and get connections from internet
	http.ListenAndServe(":9001", nil)

}
