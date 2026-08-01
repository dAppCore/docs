package help

func ExampleNewServer() {
	_ = NewServer
}

func ExampleServer_ServeHTTP() {
	_ = (*Server).ServeHTTP
}

func ExampleServer_ListenAndServe() {
	_ = (*Server).ListenAndServe
}
