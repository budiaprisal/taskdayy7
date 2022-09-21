package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {


route  := mux.NewRouter ()

//path folder public

route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))


//routing

route.HandleFunc("/",home ).Methods("GET")
route.HandleFunc("/contact",contact ).Methods("GET")
route.HandleFunc("/blog",blog ).Methods("GET")
route.HandleFunc("/blog-detail",blogdetail ).Methods("GET")
route.HandleFunc("/form-blog",formaAddBlog ).Methods("GET")
route.HandleFunc("/add-blog",addblog ).Methods("POST")
//menjalan server
fmt.Println("server running")
http.ListenAndServe("localhost:5000", route)
}

func formaAddBlog(w http.ResponseWriter, r *http.Request)  {
w.Header().Set("Content-Type", "text/html; charset=utf-8")
var tmpl, err = template.ParseFiles("view/add-blog.html")

if err != nil{
	w.Write([]byte("message : " + err.Error()))
	return
}

tmpl.Execute(w,nil)
}

func addblog(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Title: " + r.PostForm.Get("inputTitle"))
	fmt.Println("startDate: " + r.PostForm.Get("startDate"))
	fmt.Println("endDate: " + r.PostForm.Get("endDate"))
	fmt.Println("Content: " + r.PostForm.Get("inputContent"))
	fmt.Println("node: " + r.PostForm.Get("node"))
	fmt.Println("react: " + r.PostForm.Get("react"))
	fmt.Println("reacteurope: " + r.PostForm.Get("reacteurope"))
	fmt.Println("js: " + r.PostForm.Get("js"))
	fmt.Println("Image: " + r.PostForm.Get("inputImage"))
	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
	}

func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("view/index.html")
	
	if err != nil{
		w.Write([]byte("message : " + err.Error()))
		return
	}
	
	tmpl.Execute(w,nil)
	}

func contact(w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var tmpl, err = template.ParseFiles("view/contact.html")
		
		if err != nil{
			w.Write([]byte("message : " + err.Error()))
			return
		}
		
		tmpl.Execute(w,nil)
		}

		
func blog(w http.ResponseWriter, r *http.Request)  {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			var tmpl, err = template.ParseFiles("view/blog.html")
			
			if err != nil{
				w.Write([]byte("message : " + err.Error()))
				return
			}
			
			tmpl.Execute(w,nil)
			}


func blogdetail(w http.ResponseWriter, r *http.Request)  {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				var tmpl, err = template.ParseFiles("view/blog-detail.html")
				
				if err != nil{
					w.Write([]byte("message : " + err.Error()))
					return
				}
				id, _ := strconv.Atoi(mux.Vars(r)["id"])
			

				response := map[string]interface{}{
					"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
					"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) disektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian ManpowerGroup.",
					"Id":      id,
					"Duration" : "Duration",
					"Technologies" :"Technologies",
					"Date"   : "22 Agustus 2022",
					
				}


				tmpl.Execute(w,response)
				}