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

func ErrorHandler(err error){
	if err != nil {
		fmt.Println(err);
		panic("ALERT");
	}
}

func ArgumentErrorHandler(err error){
	if err != nil {
		fmt.Println("Use -h/--help");
		fmt.Println(err);
		os.Exit(0);
	}
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

func System(cmd string) {
	output, err := exec.Command(cmd).Output();
	ErrorHandler(err);
	fmt.Println(string(output));
}
