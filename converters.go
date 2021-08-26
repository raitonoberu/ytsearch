package ytsearch

import (
	"math"
	"strconv"
	"strings"
)

// durationToInt converts the duration string ("4:20") to seconds (260).
func durationToInt(duration interface{}) int {
	items := strings.Split(duration.(string), ":")
	result := 0
	for i := 0; i < len(items); i++ {
		durationInt, _ := strconv.Atoi(items[i])
		result += durationInt * int(math.Pow(60, float64(len(items)-i-1)))
	}
	return result
}

// viewCountToInt converts the view count string ("32,519 views") to int (32519).
func viewCountToInt(viewCount interface{}) int {
	// viewCount can be either (1) "100,000 views" or (2) "100 000 views"
	str := strings.Split(viewCount.(string), " ")[0]
	// (1) separated by comma (U+002C)
	str = strings.ReplaceAll(str, ",", "")
	// (2) separated by no-break space (U+00A0)
	str = strings.ReplaceAll(str, "\u00A0", "")

	result, _ := strconv.Atoi(str)
	return result
}

// descriptionToStr converts the description parts ([{"text": "foo "}, {"text": "bar"}]) to string ("foo bar").
func descriptionToStr(description []interface{}) string {
	result := make([]string, len(description))
	for i := 0; i < len(description); i++ {
		result[i] = description[i].(map[string]interface{})["text"].(string)
	}
	return strings.Join(result, "")
}
