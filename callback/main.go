package main

/*
extern void my_callback(void*);
static void my_job(void *p) {
  my_callback(p);
}
*/
import "C"
import "unsafe"

type message struct {
	text string
}

func main() {
	C.my_job(unsafe.Pointer(&message{
		text: "I love golang",
	}))
}

//export my_callback
func my_callback(p unsafe.Pointer) {
	println(((*message)(p)).text)
}
