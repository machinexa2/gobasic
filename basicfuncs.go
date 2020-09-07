package gobasic

import "fmt"
import "net/url"

func Urldecode(encoded_string string) string {
	decoded_string, err := url.QueryUnescape(encoded_string);
	ehandler(err);
	decoded_string, err = url.QueryUnescape(decoded_string);
	ehandler(err);
	return decoded_string;
}

func Ehandler(err error){
	if err != nil {
		fmt.Println(err);
	}
}
