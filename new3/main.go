package main
import "fmt"
func reformat(message string, formatter func(string) string) string {
	res := formatter(message)
	res = formatter(res)
	res = formatter(res)
return "TEXTIO: " + res
}

func main() {
    addDots := func(s string) string {
        return s + "."
    }

    result := reformat("General Kenobi", addDots)

    
    fmt.Println(result)
}
