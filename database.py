from singleton import Singleton


class Database(object):
    """
    Represents the In Memory Key Value Store.
    This class is a Singleton class. There will be
    on Database object at a time/process.
    """
    __metaclass__ = Singleton

    def __init__(self, **kwargs):
        """
        self.store = {
            "11111111": {
                "key": {
                    "encoding": "string",
                    "data": "11111111"
                },
                "value":{
                    "encoding": "string",
                    "data": "abcedefgh"
                },
            },
            "1010101010": {
                "key": {
                    "encoding": "binary",
                    "data": "1010101010"
                },
                "value":{
                    "encoding": "binary",
                    "data": "1111100000"
                },
            }
        }

        """
        self.store = {}


    def query(self, **kwargs):
        """
        Expects list of keys in kwargs["data"]
        data = [{
                encoding: 'binary',
                'data': ##1111111
                },
                {
                    encoding: 'string',
                    'data': ##222222
                }
            ]
        Return = [
                    {key: {
                            encoding: 'binary',
                            'data': ##1111111
                            },
                            value: True
                        },
                    {key: {
                            encoding: 'string',
                            'data': ##12324
                            },
                            value: False
                        }
                    ]
        """
        result = []
        getAllData = kwargs.get("getAllData")
        for key in kwargs.get("data"):
            value = self.store.get(key["data"])
            if getAllData:
                result.append(value)
            else:
                result.append({
                "key": key,
                "value": True if value else False

            })
        return result


    def fetch(self, **kwargs):
        """
        Returns all the key value pairs from database
        """
        data = kwargs.get("data")
        if data:
            self.query(**{"data": data, "getAllData": True})
        return self.store.values()


    def set(self, **kwargs):
        """
        Creates the given key value pairs if key is not present else
        Updates the given key Value pairs.
        Expects kwargs['data']
        [
            {
                key: {
                            encoding: 'binary',
                            'data': ##1111111
                        },
                value: {
                            encoding: 'binary',
                            'data': "1010101010"
                        }
            },
            {
                key: {
                            encoding: 'string',
                            'data': ##12324
                        },
                value: {
                            encoding: 'string',
                            'data': 'abcdefg'
                        }
            }
        ]
        """
        keysAdded  = 0
        keysFailed = []
        data = kwargs.get('data')
        for keyValuePair in data:
                try:
                    self.store[keyValuePair["key"]["data"]] = keyValuePair
                    keysAdded += 1
                except Exception:
                    keysFailed.append(keyValuePair)
        return keysAdded, keysFailed
