package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {

	p("Contains:   ", s.Contains("test", "es"))        // true
	p("Count:      ", s.Count("test", "t"))            // 2
	p("HasPrefix:  ", s.HasPrefix("test", "te"))       // true
	p("HasSuffix:  ", s.HasSuffix("test", "st"))       // true
	p("Index:      ", s.Index("test", "e"))            // 1
	p("Join:       ", s.Join([]string{"a", "b"}, "-")) //a-b
	p("Repeat:     ", s.Repeat("a", 5))                // aaaaa
	p("Replace:    ", s.Replace("foo", "o", "0", -1))  // f00
	p("Replace:    ", s.Replace("foooo", "o", "0", 2)) // f00oo
	p("Split:      ", s.Split("a-b-c-d-e", "-"))       // [a b c d e]
	p("ToLower:    ", s.ToLower("TEsT"))               // test
	p("ToUpper:    ", s.ToUpper("teSt"))               // TEST
	p()
	p("Len:        ", len("hello")) // 5
	p("Char:", "hello"[1])          // 101
}
