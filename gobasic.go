package gobasic

import "bufio"
import "fmt"
import "os"
import "strings"
//import "net/url"
import "github.com/chzyer/readline"
/*
func Urldecode(encoded_string string) string {
	decoded_string, _ := url.QueryUnescape(encoded_string);
	decoded_string, _ = url.QueryUnescape(decoded_string);
	return decoded_string;
}
*/
func InputRead() string {
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("> ");
	text, _ := reader.ReadString('\n');
	text = strings.Replace(text, "\n", "", -1);
	return text
}

func PrefilledInputRead(preinput string) string {
	stdin, _ := readline.New("> ");
	defer stdin.Close();
	stdin.WriteStdin([]byte(preinput));
	value, _ := stdin.Readline()
	return value
}

func InArray(check string, list []string) bool {
        for _, individual := range list {
                if individual == check {
                        return true
                }
        }
        return false
}

