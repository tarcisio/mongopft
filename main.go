package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/tarcisio/mongopft/pkg"
)

var ctx context.Context
var cancel context.CancelFunc

func main() {
	http.HandleFunc("/", perfHandler)
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

var home string = `
<html>
<form method="post">
<label>DSN: <input type="text" name="dsn" value="mongodb://<host>:<port>" width="200"></label><br>
<label>Número de threads: <input type="text" name="n" value="10" width="200"></label><br>
<input type="submit" name="acao" value="init"><br>
<input type="submit" name="acao" value="stop">
</form>
</html>
`

func perfHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		acao := r.FormValue("acao")
		dsn := r.FormValue("dsn")
		nstr := r.FormValue("n")
		nint, err := strconv.Atoi(nstr)

		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		switch acao {
		case "init":
			ctx, cancel = context.WithCancel(context.Background())

			for i := 0; i < nint; i++ {
				go pkg.TestThread(ctx, dsn)
			}

			fmt.Fprintln(w, "init")
		case "stop":
			cancel()
			fmt.Fprintln(w, "stop")
		}
		return
	default:
		fmt.Fprintln(w, home)
	}
}
