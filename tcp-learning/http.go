func sayhelloName(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello I am astaxie!”)
}
func main() {
http.HandleFunc("/", sayhelloName)
err := http.ListenAndServe(":9090", nil)
if err != nil {
log.Fatal("ListenAndServe: ", err)
}
}
