package example03

import "testing"

func TestFrTranslation(t *testing.T) {

    got := translate("fr-FR")
    want := "Bonjour "
    if got != want {
        t.Fatalf("Expected %q, got %q", want, got)
    }
}

func testUSTranslation(t *testing.T) {
     got := Greeting("George", "en-US")
     want := "Hello George"
     if got != want {
         t.Fatalf("Expected %q, got %q", want, got)
     }
}
