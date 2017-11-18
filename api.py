from urlparse import urlparse

from database import Database

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
        print self.database.fetch()

    def query(self, data=None):
        pass

    def set(self, data=None):
        pass
