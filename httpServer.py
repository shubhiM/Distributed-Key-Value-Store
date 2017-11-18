import json
import sys
from urlparse import urlparse
from BaseHTTPServer import HTTPServer
from BaseHTTPServer import BaseHTTPRequestHandler

from api import API


class MyHttpRequestHandler(BaseHTTPRequestHandler):

    def _do_METHOD_HELPER(self, type):
        parsedPath = urlparse(self.path)
        result = None
        statusCode = None
        api = API(self)
        if type == "GET":
            result, statusCode =  getattr(api, parsedPath.path[1:])()
        else:
            content_len = int(self.headers.getheader('content-length', 0))
            body = self.rfile.read(content_len)
            body = json.loads(body)
            result, statusCode =  getattr(api, parsedPath.path[1:])(**{
                "data": body
            })
        self.send_response(statusCode)
        self.send_header("Content-type", "application/json")
        self.end_headers()
        self.wfile.write(json.loads(json.dumps(result)))
        return

    def do_GET(self):
        """
        Handles all the GET requests
        """
        return self._do_METHOD_HELPER("GET")

    def do_POST(self):
        """
        Handles all the POST requests
        """
        return self._do_METHOD_HELPER("POST")

    def do_PUT(self):
        """
        Handles all the PUT requests
        """
        return self._do_METHOD_HELPER("PUT")


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
