
package main

import (
    "flag"
    "log"
    "net/http"
	"html/template"
	"fmt"
	"github.com/user/sql"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
	http.Handle("/", http.HandlerFunc(test1))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func test1(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("test1...")
	user.PrintAreaName()
	templ.Execute(w, req.FormValue("s"))
}

func test2(w http.ResponseWriter, req *http.Request) {
    log.Fatal("test2........")
}

func test3(w http.ResponseWriter, req *http.Request) {
    log.Fatal("test3........")
}

func test4(w http.ResponseWriter, req *http.Request) {
    log.Fatal("test4........")
}




const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`


