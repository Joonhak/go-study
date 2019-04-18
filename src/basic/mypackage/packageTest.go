package mypackage

var helloMessage map[string]string

func init() {
	// package import 시 실행되는 함수
	helloMessage = make(map[string]string)
	helloMessage["say"] = "Hello world!!!"
	helloMessage["whisper"] = "hi :)"
}

// public function
func SayHello() string {
	return helloMessage["say"]
}

// private function
func whisper() string {
	return helloMessage["whisper"]
}