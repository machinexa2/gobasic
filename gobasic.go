package gobasic

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"net/url"
	"github.com/chzyer/readline"
)

func Urldecode(encoded_string string) string {
	decoded_string, err := url.QueryUnescape(encoded_string);
	ErrorHandler(err);
	decoded_string, err = url.QueryUnescape(decoded_string);
	ErrorHandler(err);
	return decoded_string;
}

func InputRead() string {
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("> ");
	text, err := reader.ReadString('\n');
	ErrorHandler(err);
	text = strings.Replace(text, "\n", "", -1);
	return text
}

func PrefilledInputRead(preinput string) string {
	stdin, _ := readline.New("> ");
	defer stdin.Close();
	stdin.WriteStdin([]byte(preinput));
	value, err := stdin.Readline()
	ErrorHandler(err);
	return value
}
