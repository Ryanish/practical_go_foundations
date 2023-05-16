package nlp

import (
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

// var tokenizeCases = []struct { //anonymous struct
// 	text   string
// 	tokens []string
// }{
// 	{"Who's on first?", []string{"who", "s", "on", "first"}},
// 	{"", nil},
// }

// excercise: read test cases from the toml file

type tokenizeCase struct {
	Text   string
	Tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	// data, err := ioutil.ReadFile("testdata/tokenize_cases.toml")
	// require.NoError(t, err)

	var testCases struct {
		Cases []tokenizeCase
	}

	//err = toml.Unmarshal(data, &testCases)
	filePath := `/Users/ryanbulloch/go/src/gofoundation/Day4/testdata/tokenize_cases.toml`
	_, err := toml.DecodeFile(filePath, &testCases)
	require.NoError(t, err, "UnMarshal TOML")
	return testCases.Cases
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "what's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	/* before Testify */
	//if !reflect.DeepEqual(expected, tokens) {
	//	t.Fatalf("expected %#v, got %#v", expected, tokens)
	//}

}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})

}
