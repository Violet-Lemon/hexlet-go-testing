// файл hello/hello.go
package hello

func Hello(name string) string {
	if name == "" {
		return "Hello, world!"
	}

	return "Hello, " + name + "!"
}

func IsEven(n int) bool {
	return n%2 == 0
}
