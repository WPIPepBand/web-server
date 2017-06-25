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
        return not '.' in self.path
    
    def do_GET(self):
        print('In directory: ' + os.getcwd())
        if self.is_cgi() and not 'cgi-bin' in os.getcwd(): os.chdir(public_html + os.sep + 'cgi-bin')
        elif not self.is_cgi() and 'cgi-bin' in os.getcwd(): os.chdir(os.pardir)
        
        if self.is_cgi():
            query = 'page=' + ('Home' if self.path == '/' else urllib.parse.unquote(self.path[1:]))
            os.environ['QUERY_STRING'] = query
        
        super(GoCGIHandler, self).do_GET()      # Let the superclass actually do the work

if __name__ == '__main__':
    handler = GoCGIHandler
    tester = HTTPServer(('', 8080), handler)
    print('Serving http://pepband.wpi.edu locally on port 8080')
    tester.serve_forever()