#!/bin/bash
echo "Reading config." >&2
source $1
echo "Config for the cgi-bin: $cgi_bin" >&2
echo "config for the htdocs: $htdocs" >&2

#git submodule update --remote ./public_html/cgi-bin/pages
go build -o "$cgi_bin"make-page.cgi ./make-page.go
go build -o "$cgi_bin"anon-sugg.cgi ./anon-suggestion.go
