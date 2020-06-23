package main

import("fmt" 
"net/http")

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Hellow world</h1>")
}

func main()  {
	http.HandleFunc("/", index)
	fmt.Println("server starting...")
	http.ListenAndServe(":3000", nil)
}
