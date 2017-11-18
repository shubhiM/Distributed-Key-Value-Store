from urlparse import urlparse
from database import Database


HTTP_SUCCESS = 200
HTTP_NOT_FOUND = 404
HTTP_BAD_REQUEST = 400
HTTP_SERVER_ERROR = 500


class API(object):
    """
    Defines the public APIs for the http Server
    """
    def __init__(self, request):
        """
        """
        parsedPath = urlparse(request.path)
        self.database = Database()
        getattr(self, parsedPath.path[1:])()

    def fetch(self, data=None):
        """
        Path: /fetch
        Method: GET
        Result:
        """
        result = self.database.fetch()
        return result, HTTP_SUCCESS

    def query(self, data=None):
        keys = self.database.query(**{"data": data})
        keysFound = filter(lambda k: k["value"] == True, keys)
        statusCode = HTTP_SUCCESS if len(keysFound) == len(keys) else HTTP_NOT_FOUND
        return keys, statusCode

    def set(self, data=None):
        keysAdded, keysFailed = self.database.set(**{"data": data})
        result = {
            "keys_added": keysAdded,
            "keys_failed": keysFailed
        }
        statusCode = HTTP_BAD_REQUEST if keysFailed else HTTP_SUCCESS
        return result, statusCode
