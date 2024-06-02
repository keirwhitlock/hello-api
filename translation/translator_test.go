package translation

import (
	"testing"
)

func TestTranslate(t *testing.T) {

	// Arrage
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "french",
			Translation: "bonjour",
		},
	}

	underTest := NewStaticService()

	for _, test := range tt {
		res := underTest.Translate(test.Word, test.Language)
		if res != test.Translation {
			t.Errorf(`expected "%s" to be "%s" from "%s" but received "%s"`, test.Word, test.Translation, test.Translation, res)
		}
	}
}
