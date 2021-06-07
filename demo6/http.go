package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// tips 这是个 interface，引用类型
func greet(w io.Writer, format string) {
	_, _ = fmt.Fprintf(w, "hello, %s", "world")
}

// tips 这是个 struct，值类型
func greet2(w *bytes.Buffer, format string) {
	// w.String() -> obj.toString()
	fmt.Fprintf(w, format, "world")
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	greet(w, "world")
}

func main() {
	// HandlerFunc 是个 func(w IOWriter, r *Http.Request) 类型，   HandlerFunc(funcParam) 的语法，表示将 greetingHandler 强转成 HandlerFunc 类型
	http.ListenAndServe(":5000", http.HandlerFunc(greetingHandler))
}
