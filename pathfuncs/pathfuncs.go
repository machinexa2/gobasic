package pathfunctions
import "strings"

func urler(xpath string) string{
	if !(strings.HasPrefix(xpath, "http://") || strings.HasPrefix(xpath, "https://")){
		return "http://" + xpath
	}
	return xpath
}

func ender(xpath string, ypath string) string{
	if xpath[len(xpath)-1:] != ypath{
		xpath += ypath
	}
	return xpath
}

func unstarter(xpath string, ypath string) string{
	if xpath[:1] == ypath{
		return xpath[1:]
	}
	return xpath
}


