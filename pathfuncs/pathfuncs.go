package pathfunctions
import "strings"

func Urler(xpath string) string{
	if !(strings.HasPrefix(xpath, "http://") || strings.HasPrefix(xpath, "https://")){
		return "http://" + xpath
	}
	return xpath
}

func Ender(xpath string, ypath string) string{
	if xpath[len(xpath)-1:] != ypath{
		xpath += ypath
	}
	return xpath
}

func Unstarter(xpath string, ypath string) string{
	if xpath[:1] == ypath{
		return xpath[1:]
	}
	return xpath
}


