package rfc4226

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var Secret = "12345678901234567890"

var Ans = map[uint64]string{
	0: "cc93cf18508d94934c64b65d8ba7667fb7cde4b0",
	1: "75a48a19d4cbe100644e8ac1397eea747a2d33ab",
	2: "0bacb7fa082fef30782211938bc1c5e70416ff44",
	3: "66c28227d03a2d5529262ff016a1e6ef76557ece",
	4: "a904c900a64b35909874b33e61c5938a8e15ed1c",
	5: "a37e783d7b7233c083d4f62926c7a25f238d0316",
	6: "bc9cd28561042c83f219324d3c607256c03272ae",
	7: "a4fb960c0bc06e1eabb804e5b397cdc4b45596fa",
	8: "1b3c89f65e6c9e883012052823443f048b4332db",
	9: "1637409809a679dc698207310c8c7fc07290d9e5",
}

func assertSecretCount(t *testing.T, got, want string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func TestEncryptSHA1(t *testing.T) {
	count := uint64(0)
	got := hex.EncodeToString(EncryptSHA1(Secret, count))
	//want, err := hex.DecodeString(Ans[count])
	want := Ans[count]
	assertSecretCount(t, got, want)
}

func TestDT(t *testing.T) {
	got := DT(EncryptSHA1(Secret, 0), 6)
	want := 755224
	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}
