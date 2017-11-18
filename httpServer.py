import sys
import urlparse
from BaseHTTPServer import HTTPServer
from BaseHTTPServer import BaseHTTPRequestHandler

from api import API


class MyHttpRequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        """
        Handles all the GET requests
        """
        # TODO:check for the call validity
        api = API(self)

    def do_POST(self):
        """
        Handles all the POST requests
        """
        # TODO:check for the call validity
        pass

    def do_HEAD(self):
        """
        """
        # TODO:check for the call validity
        pass

    def do_PUT(self):
        pass


if __name__ == "__main__":
    host = None
    port = None
    try:
        host = sys.argv[1]
        port = sys.argv[2]
    except:
        host = 'localhost'
        port = 8080
    finally:
        print "Server is running at " + host + ":" + str(port)
        server = HTTPServer((host, int(port)), MyHttpRequestHandler)
        server.serve_forever()
