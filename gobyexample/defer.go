package main

import "fmt"
import "os"

/* Defer (연기하다)

프로그램 실행 이후에 함수 호출이 되는 것을 보장한다.
다른 언어의 finally 와 비슷.
*/

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // main이 끝나기 직전에 실행된다.
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
