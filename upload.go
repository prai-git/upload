// routes.go
package upload

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func render(w http.ResponseWriter, tmpl string, context map[string]interface{}) {
	tmpl_list := []string{fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func mainfun() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", Upload).Methods("GET")
	r.HandleFunc("/fileUpload", UploadFile).Methods("POST")

	fmt.Println("Listining port 8080")
	http.ListenAndServe(":8080", r)
}

func Upload(w http.ResponseWriter, q *http.Request) {

	render(w, "form", nil)

}

func UploadFile(w http.ResponseWriter, q *http.Request) {
	fmt.Println("myselfP")
	err := q.ParseForm()

	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	uploadImage(q, "files")

	http.Redirect(w, q, "/upload", http.StatusFound)

}

func uploadImage(q *http.Request, nameInForm string) {
	fmt.Println(q.FormFile)
	q.ParseMultipartForm(32 << 20)

	file, _, err := q.FormFile(nameInForm)
	fileName := q.Header.Get("X_FILENAME")
	totalNumberOfFile := q.Header.Get("totalNumberOfFiles")

	fmt.Println(q)
	fmt.Println(fileName)
	fmt.Println(totalNumberOfFile)
	ext := path.Ext(fileName)

	if err != nil {
		fmt.Println(err)

	}
	defer file.Close()

	f, err := os.OpenFile("/scm/uploadimages/"+uuid.NewV4().String()+ext, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)

	}
	defer f.Close()
	io.Copy(f, file)
}
