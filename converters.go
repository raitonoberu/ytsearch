package youtube_search

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
	items := strings.Split(strings.Split(viewCount.(string), " ")[0], ",")
	str := strings.Join(items, "")
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
