package gobasic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/url"
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
	}
}

func InputRead() string {
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("> ");
	text, _ := reader.ReadString('\n');
	text = strings.Replace(text, "\n", "", -1);
	return text
}
