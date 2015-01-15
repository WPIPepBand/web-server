/*
make-page.go
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
	"github.com/russross/blackfriday"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cgi"
	"net/url"
	"os"
	"text/template"
	"time"
)

type PageInfo struct {
	Title      string
	Breadcrumb []string
	Filename   string
	Contents   string
	Year       int
}

func main() {
	err := cgi.Serve(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		header := response.Header()
		header.Set("Content-Type", "text/html; charset=utf-8")

		/* Go uses the URL instead of the query string for request.Form, so we need to get the query string environment
		variable, since the URL is rewritten with mod_rewrite. */
		queryValues, err := url.ParseQuery(os.Getenv("QUERY_STRING"))

		if err != nil {
			log.Fatal(err)
		}

		err = writePage(response, queryValues["page"][0])

		if err != nil {
			log.Fatal(err)
		}
	}))

	if err != nil {
		log.Fatal(err)
	}
}

/*
Get the full HTML of the given page and write it to the given Writer
*/
func writePage(w io.Writer, pageName string) error {
	templ, err := template.New("page").ParseFiles("template.html")

	if err != nil {
		return err
	}

	/* The title, filename, and breadcrumb are stored statically in a JSON file */
	jsonFile, _ := os.Open("pages/pages.json")
	jsonDecoder := json.NewDecoder(jsonFile)

	var pages map[string]PageInfo
	err = jsonDecoder.Decode(&pages)

	if err != nil {
		return err
	}

	var pageInfo PageInfo
	pageInfo = pages[pageName]

	/* The rest of the PageInfo fields are filled in dynamically */
	breadcrumb, contents := parseMarkdown(pageContents(pageName))
	pageInfo.Contents = contents
	pageInfo.Breadcrumb = breadcrumb
	pageInfo.Year = time.Now().Year()

	return templ.ExecuteTemplate(w, "template.html", pageInfo)
}

/*
Get the markdown contents of a page, or of the 404 page if the files does not exit
TODO: Parse the markdown into HTML
*/
func pageContents(page string) string {
	pageContents, err := ioutil.ReadFile("pages/" + page + ".md")
	if err != nil {
		pageContents, _ = ioutil.ReadFile("pages/404.md")
		return string(pageContents) + "\n\n    " + err.Error()
	} else {
		return string(pageContents)
	}
}

/*
Convert the page markdown content into HTML.
TODO: Parse out breadcrumb.
*/
func parseMarkdown(content string) (breadcrumb []string, htmlContent string) {
	output := blackfriday.MarkdownCommon([]byte(content))
	return []string{"Home"}, string(output[:])
}
