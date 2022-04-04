package hello

import "testing"
import "fmt"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Worlde", "French")
		want := "Bonjour, Worlde"
		assertCorrectMessage(t, got, want)
	})
}

/* the example function will not be executed if you remove the comment
	output: Hello, World. Although the function will be compiled, it won't be executed.
 */
func ExampleHello() {
	got := Hello("World","English")
	fmt.Println(got)
	// output: Hello, World
}

// enter command "godoc -http=:6060" to start the local server for the document