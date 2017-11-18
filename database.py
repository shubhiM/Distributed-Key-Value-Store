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


    def query(self, **kwargs):
        """
        Expects list of keys in kwargs["data"]
        data = [{
                enconding: 'binary',
                'data': ##1111111
                },
                {
                    enconding: 'string',
                    'data': ##222222
                }
            ]
        Return = [
                    {key: {
                            enconding: 'binary',
                            'data': ##1111111
                            },
                            value: True
                        },
                    {key: {
                            enconding: 'string',
                            'data': ##12324
                            },
                            value: False
                        }
                    ]
        """
        result = []
        for key in kwargs.get("data"):
            value = self.store.get(key["data"])
            result.append({
                "key": key,
                "value": True if value else False

            })
        return result


    def fetch(self):
        """
        Returns all the key value pairs from database
        """
        return self.store.values()


    def set(self, **kwargs):
        """
        Creates the given key value pairs if key is not present else
        Updates the given key Value pairs.
        Expects kwargs['data']
        [
            {
                key: {
                            enconding: 'binary',
                            'data': ##1111111
                        },
                value: {
                            enconding: 'binary',
                            'data': "1010101010"
                        }
            },
            {
                key: {
                            enconding: 'string',
                            'data': ##12324
                        },
                value: {
                            enconding: 'string',
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
                    self.store[keyValuePair['key']] = keyValuePair
                    keysAdded += 1
                except Exception:
                    keysFailed.append(keyValuePair)
        return keysAdded, keysFailed
