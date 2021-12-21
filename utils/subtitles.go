package utils

import (
	"encoding/json"
	"strings"

	"github.com/sahilm/fuzzy"
)

// "subtile" matches with both "subtitle"s and typos of it.
// Some inputs have "no subtitles" so we also cover that case using
// "no subtile".
// Sample size is used for searching around "subtitle" word.
const (
	matchExpr         = "subtile"
	skipExpr          = "no subtile"
	noExpr            = "no "
	minSampleSizeHalf = 40
)

const langsData = `{"abkhazian":"Ab","afar":"Aa","afrikaans":"Af","akan":"Ak","albanian":"Sq","amharic":"Am","arabic":"Ar","aragonese":"An","armenian":"Hy","assamese":"As","avaric":"Av","avestan":"Ae","aymara":"Ay","azerbaijani":"Az","bambara":"Bm","bashkir":"Ba","basque":"Eu","belarusian":"Be","bengali":"Bn","bihari":"Bh","bislama":"Bi","bokmål":"Nb","bosnian":"Bs","breton":"Br","bulgarian":"Bg","burmese":"My","castilian":"Es","catalan":"Ca","central khmer":"Km","chamorro":"Ch","chechen":"Ce","chewa":"Ny","chichewa":"Ny","chinese":"Zh","chuang":"Za","chuvash":"Cv","cornish":"Kw","corsican":"Co","cree":"Cr","croatian":"Hr","czech":"Cs","danish":"Da","dhivehi":"Dv","divehi":"Dv","dutch":"Nl","dzongkha":"Dz","english":"En","esperanto":"Eo","estonian":"Et","ewe":"Ee","faroese":"Fo","fijian":"Fj","finnish":"Fi","flemish":"Nl","french":"Fr","frisian":"Fy","fulah":"Ff","gaelic":"Gd","galician":"Gl","ganda":"Lg","georgian":"Ka","german":"De","gikuyu":"Ki","greek":"El","greenlandic":"Kl","guarani":"Gn","gujarati":"Gu","haitian":"Ht","haitian creole":"Ht","hausa":"Ha","hebrew":"He","herero":"Hz","hindi":"Hi","hiri motu":"Ho","hungarian":"Hu","icelandic":"Is","ido":"Io","igbo":"Ig","indonesian":"Id","inuktitut":"Iu","inupiaq":"Ik","irish":"Ga","italian":"It","japanese":"Ja","javanese":"Jv","kalaallisut":"Kl","kannada":"Kn","kanuri":"Kr","kashmiri":"Ks","kazakh":"Kk","khmer":"Km","kikuyu":"Ki","kinyarwanda":"Rw","kirghiz":"Ky","komi":"Kv","kongo":"Kg","korean":"Ko","kuanyama":"Kj","kurdish":"Ku","kwanyama":"Kj","kyrgyz":"Ky","lao":"Lo","latin":"La","latvian":"Lv","letzeburgesch":"Lb","limburgan":"Li","limburger":"Li","limburgish":"Li","lingala":"Ln","lithuanian":"Lt","luba-katanga":"Lu","luxembourgish":"Lb","macedonian":"Mk","malagasy":"Mg","malay":"Ms","malayalam":"Ml","maldivian":"Dv","maltese":"Mt","manx":"Gv","maori":"Mi","marathi":"Mr","marshallese":"Mh","moldavian":"Ro","moldovan":"Ro","mongolian":"Mn","nauru":"Na","navaho":"Nv","navajo":"Nv","ndonga":"Ng","nepali":"Ne","niuean":"","northern sami":"Se","norwegian":"No","nuosu":"Ii","nyanja":"Ny","nynorsk":"Nn","occitan":"Oc","ojibwa":"Oj","oriya":"Or","oromo":"Om","ossetian":"Os","ossetic":"Os","pali":"Pi","panjabi":"Pa","pashto":"Ps","persian":"Fa","polish":"Pl","portuguese":"Pt","punjabi":"Pa","pushto":"Ps","quechua":"Qu","romanian":"Ro","romansh":"Rm","rundi":"Rn","russian":"Ru","sami":"Se","samoan":"Sm","sandawe":"","sango":"Sg","sanskrit":"Sa","sardinian":"Sc","scottish gaelic":"Gd","serbian":"Sr","shona":"Sn","sichuan yi":"Ii","sindhi":"Sd","sinhala":"Si","sinhalese":"Si","slovak":"Sk","slovenian":"Sl","somali":"So","sotho":"St","spanish":"Es","sukuma":"","sundanese":"Su","swahili":"Sw","swati":"Ss","swedish":"Sv","tagalog":"Tl","tahitian":"Ty","tajik":"Tg","tamil":"Ta","tatar":"Tt","telugu":"Te","thai":"Th","tibetan":"Bo","tigre":"","tigrinya":"Ti","tsonga":"Ts","tswana":"Tn","turkish":"Tr","turkmen":"Tk","twi":"Tw","uighur":"Ug","ukrainian":"Uk","urdu":"Ur","uyghur":"Ug","uzbek":"Uz","valencian":"Ca","venda":"Ve","vietnamese":"Vi","volapük":"Vo","walloon":"Wa","welsh":"Cy","western frisian":"Fy","wolof":"Wo","xhosa":"Xh","yiddish":"Yi","yoruba":"Yo","zhuang":"Za","zulu":"Zu"}`

var langsToCodes map[string]string

func init() {
	if err := json.Unmarshal([]byte(langsData), &langsToCodes); err != nil {
		panic(err)
	}
}

// FindSubtitleLanguageCodes finds subtitle language codes from given input.
func FindSubtitleLanguageCodes(input string, force bool) (langs []string) {
	words := strings.Split(input, " ")
	matches := fuzzy.Find(matchExpr, words)

	var index int
	switch {
	case len(matches) == 0 && !force:
		return

	case len(matches) == 0 && force:

	default:
		found := matches[0].Str
		index = strings.LastIndex(input, found)
		if index < 0 && !force {
			return
		}
	}

	length := len(input)
	maxIndex := index + minSampleSizeHalf
	if maxIndex > length {
		maxIndex = length
	}
	minIndex := index - minSampleSizeHalf
	if minIndex < 0 {
		minIndex = 0
	}

	sample := strings.ToLower(input[minIndex:maxIndex])
	matches = fuzzy.Find(skipExpr, []string{sample})
	if len(matches) > 0 && strings.Contains(sample, noExpr) {
		return
	}

	return langCodesFromSample(sample)
}

func langCodesFromSample(sample string) (langs []string) {
	sample = strings.ToLower(strings.TrimSpace(sample))

	// Avoid audio language in sample which comes before "with".
	if i := strings.Index(sample, "with"); i > 0 {
		sample = sample[i+5:]
	}

	sample = strings.ReplaceAll(sample, ".", "")
	sample = strings.ReplaceAll(sample, ",", "")

	for _, word := range strings.Split(sample, " ") {
		code, ok := langsToCodes[word]
		if ok {
			langs = append(langs, code)
		}
	}

	return
}
