/*
anon-suggestion.go
Copyright WPI Pep Band 2015

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cgi"
	"net/smtp"
)

type recaptchaResponse struct {
	Success    bool     `success`
	ErrorCodes []string `error-codes`
}

func main() {
	if err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		http.Redirect(w, r, "../Contact Us", http.StatusFound)

		r.ParseForm()
		post := r.PostForm
		response, err := http.Get("https://www.google.com/recaptcha/api/siteverify?secret=6LfmVAETAAAAAIeSMi9xdgBN4A20HrvOfjBnvwS7&response=" +
			post.Get("g-recaptcha-response") + "&remoteip=" + r.RemoteAddr)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var res recaptchaResponse
		json.Unmarshal(body, &res)
		suggheader := "Anonymous Suggestion" + "\n\n"
		if res.Success {
			err := smtp.SendMail("localhost:25", nil, "pep-suggestion@wpi.edu", []string{"pepoff@wpi.edu"}, []byte(suggheader+post.Get("suggestion")+post.Get("Name"))
			if err != nil {
				log.Fatal(err)
			}
		}
	})); err != nil {
		log.Fatal(err)
	}
}
