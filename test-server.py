# test-server.py
# A simple server that runs locally and serves the Pep Band website from your machine
#
# Usage: python test-server.py
# Notes:
# - Point your browser to 'localhost:8080' to get to the Pep Band website Home page
# - From there, all navigation works as normal
#
# Author:   Tanuj Sane
# Since:    6/24/2017
# Version:  1.0
#
# Changelog:
# - 1.0 Initial commit

import os, urllib
from http.server import CGIHTTPRequestHandler, HTTPServer

public_html = os.path.normpath('public_html')

class GoCGIHandler(CGIHTTPRequestHandler):
    def is_cgi(self):
        self.cgi_info = '', 'make-page.cgi'     # Any request will always be to make-page.cgi
        return True
    
    def do_GET(self):
        os.chdir(public_html + os.sep + 'cgi-bin' if self.is_cgi() else public_html)
        
        query = 'page=' + ('Home' if self.path == '/' else urllib.parse.unquote(self.path[1:]))
        print(query)
        os.environ['QUERY_STRING'] = query
        
        super(GoCGIHandler, self).do_GET()      # Let the superclass actually do the work

if __name__ == '__main__':
    handler = GoCGIHandler
    tester = HTTPServer(('', 8080), handler)
    print('Serving http://pepband.wpi.edu locally on port 8080')
    tester.serve_forever()