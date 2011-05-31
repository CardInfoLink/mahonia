package mahonia

import "testing"

var nameTests = map[string]string{
	"utf8":       "utf8",
	"ISO 8859-1": "iso88591",
	"Big5":       "big5",
	"":           "",
}

func TestSimplifyName(t *testing.T) {
	for name, simple := range nameTests {
		if simple != simplifyName(name) {
			t.Errorf("%s came out as %s instead of as %s", name, simplifyName(name), simple)
		}
	}
}

var testData = []struct {
	utf8, other, otherEncoding string
}{
	{"Résumé", "Résumé", "utf8"},
	{"Résumé", "R\xe9sum\xe9", "latin-1"},
	{"これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020", "UTF-16LE"},
	{"これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16BE"},
	{"これは漢字です。", "\xfe\xff0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16"},
	{"𝄢𝄞𝄪𝄫", "\xfe\xff\xd8\x34\xdd\x22\xd8\x34\xdd\x1e\xd8\x34\xdd\x2a\xd8\x34\xdd\x2b", "UTF-16"},
	{"Hello, world", "Hello, world", "ASCII"},
	{"Gdańsk", "Gda\xf1sk", "ISO-8859-2"},
	{"Ââ Čč Đđ Ŋŋ Õõ Šš Žž Åå Ää", "\xc2\xe2 \xc8\xe8 \xa9\xb9 \xaf\xbf \xd5\xf5 \xaa\xba \xac\xbc \xc5\xe5 \xc4\xe4", "ISO-8859-10"},
	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "ISO-8859-11"},
	{"latviešu", "latvie\xf0u", "ISO-8859-13"},
	{"Seònaid", "Se\xf2naid", "ISO-8859-14"},
	{"€1 is cheap", "\xa41 is cheap", "ISO-8859-15"},
	{"românește", "rom\xe2ne\xbate", "ISO-8859-16"},
	{"nutraĵo", "nutra\xbco", "ISO-8859-3"},
	{"Kalâdlit", "Kal\xe2dlit", "ISO-8859-4"},
	{"русский", "\xe0\xe3\xe1\xe1\xda\xd8\xd9", "ISO-8859-5"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "ISO-8859-7"},
	{"Kağan", "Ka\xf0an", "ISO-8859-9"},
	{"Résumé", "R\x8esum\x8e", "macintosh"},
	{"Gdańsk", "Gda\xf1sk", "windows-1250"},
	{"русский", "\xf0\xf3\xf1\xf1\xea\xe8\xe9", "windows-1251"},
	{"Résumé", "R\xe9sum\xe9", "windows-1252"},
	{"ελληνικά", "\xe5\xeb\xeb\xe7\xed\xe9\xea\xdc", "windows-1253"},
	{"Kağan", "Ka\xf0an", "windows-1254"},
	{"עִבְרִית", "\xf2\xc4\xe1\xc0\xf8\xc4\xe9\xfa", "windows-1255"},
	{"العربية", "\xc7\xe1\xda\xd1\xc8\xed\xc9", "windows-1256"},
	{"latviešu", "latvie\xf0u", "windows-1257"},
	{"Việt", "Vi\xea\xf2t", "windows-1258"},
	{"สำหรับ", "\xca\xd3\xcb\xc3\u047a", "windows-874"},
	{"русский", "\xd2\xd5\xd3\xd3\xcb\xc9\xca", "KOI8-R"},
	{"українська", "\xd5\xcb\xd2\xc1\xa7\xce\xd3\xd8\xcb\xc1", "KOI8-U"},
	{"Hello 常用國字標準字體表", "Hello \xb1`\xa5\u03b0\xea\xa6r\xbc\u0437\u01e6r\xc5\xe9\xaa\xed", "big5"},
	{"これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "SJIS"},
}

func TestDecode(t *testing.T) {
	utf8 := GetCharset("UTF-8")
	for _, data := range testData {
		cs := GetCharset(data.otherEncoding)
		if cs == nil {
			t.Errorf("Could not create decoder for %s", data.otherEncoding)
			continue
		}

		str := ConvertString(data.other, cs, utf8)

		if str != data.utf8 {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.utf8)
		}
	}
}

func TestEncode(t *testing.T) {
	utf8 := GetCharset("UTF-8")
	for _, data := range testData {
		cs := GetCharset(data.otherEncoding)
		if cs == nil {
			t.Errorf("Could not create encoder for %s", data.otherEncoding)
			continue
		}

		str := ConvertString(data.utf8, utf8, cs)

		if str != data.other {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.other)
		}
	}
}
