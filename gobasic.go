package gobasic

import "fmt"
import "net/url"

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
