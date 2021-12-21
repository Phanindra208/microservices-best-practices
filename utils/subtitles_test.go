package utils_test

import (
	"strings"
	"testing"

	"github.com/iAmPlus/microservice/utils"
)

func TestFindSubtitleLanguageCodes(t *testing.T) {
	testCases := []struct {
		input         string
		expectedLangs string
		force         bool
	}{
		{
			input:         "lorem ipsum dolor sit amet\r\nSubtitle : Arabic\r\n",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle : Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "With Arabic Subtitle",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtittle arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtittle arabic and English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "Subtitle in Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle in Arabic and French",
			expectedLangs: "Ar,Fr",
		},
		{
			input:         "Subtitle - Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle - Arabic & English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "Subtilte : Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtille: Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle : French & Arabic",
			expectedLangs: "Fr,Ar",
		},
		{
			input:         "Subtitled in Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtiled in Arabic and French",
			expectedLangs: "Ar,Fr",
		},
		{
			input:         "Subtitled in Arabic and English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "with English Subtitles",
			expectedLangs: "En",
		},
		{
			input:         "subtitled in English and Arabic",
			expectedLangs: "En,Ar",
		},
		{
			input:         "with no subtitles",
			expectedLangs: "",
		},
		{
			input:         "subtitle in Arabic and English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "With English and Arabic subtitle",
			expectedLangs: "En,Ar",
		},
		{
			input:         "with  and French Subtitles",
			expectedLangs: "Fr",
		},
		{
			input:         "With english subtitle",
			expectedLangs: "En",
		},
		{
			input:         "Armenian and French with  Subtitles",
			expectedLangs: "",
		},
		{
			input:         "With French & Arabic Subtitles",
			expectedLangs: "Fr,Ar",
		},
		{
			input:         "With Arabic and French subtitles",
			expectedLangs: "Ar,Fr",
		},
		{
			input:         "Subtitle Arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle: Arabic & French",
			expectedLangs: "Ar,Fr",
		},
		{
			input:         "Subtitle Arabic and English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "Arabic subtitle",
			expectedLangs: "Ar",
		},
		{
			input:         "Subtitle : Arabic and English",
			expectedLangs: "Ar,En",
		},
		{
			input:         "Subtiles : English & Arabic",
			expectedLangs: "En,Ar",
		},
		{
			input:         "Withe Arabic and English Subtitle",
			expectedLangs: "Ar,En",
		},
		{
			input:         "With Arabic & English Subtitle",
			expectedLangs: "Ar,En",
		},
		{
			input:         "English Subtitle (Arabic)",
			expectedLangs: "En",
		},
		{
			input:         "(No Arabic Subtitle)",
			expectedLangs: "",
		},
		{
			input:         "Kate Hudson, Candice Bergen\r\nLanguage English with Subtittle arabic",
			expectedLangs: "Ar",
		},
		{
			input:         "Main cast:\r\nVoices of : Dan Aykroyd, Christine Taylor, Justin Timberlake\r\n\r\nSubtitle - Arabic \r\n\r\nGenre: Animation, Adventure \r\n",
			expectedLangs: "Ar",
		},
		{
			input:         "Main cast:\r\nJames McAvoy, Emily blunt, Maggie Smith\r\nDirector: Kelly Asbury Genre: Animation, Adventure\r\nSubtitle - Arabic ",
			expectedLangs: "Ar",
		},
		{
			input:         "Arabic, French,.",
			expectedLangs: "Ar,Fr",
			force:         true,
		},
		{
			input:         "layers of deceit to reveal the terrible truth behind SPECTRE.\r\nSubtitle: Arabic & French.",
			expectedLangs: "Ar,Fr",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			langs := utils.FindSubtitleLanguageCodes(testCase.input, testCase.force)
			joint := strings.Join(langs, ",")

			if joint != testCase.expectedLangs {
				t.Fatalf("languages: expected '%v' but got '%v'", testCase.expectedLangs, joint)
			}
		})
	}
}
