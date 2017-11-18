from singleton import Singleton


class Database(object):
    """
    Represents the In Memory Key Value Store.
    This class is a Singleton class. There will be
    on Database object at a time/process.
    """
    __metaclass__ = Singleton

    def __init__(self, **kwargs):
        self.store = {}
        self.file = str(kwargs.get("port")) + '.json'
        self.count = 0

    def query(self, **kwargs):
        """
        Expects Keys in kwargs
        Return key value pairs from database
        """
        pass

    def fetch(self):
        """
        Returns all the key value pairs from database
        """
        self.count = self.count + 1
        return self.count

    def set(self, **kwargs):
        """
        Creates the given key/value pairs if key is not present else
        Updates the given key/Value pairs.
        """
        self.count = self.count + 1
