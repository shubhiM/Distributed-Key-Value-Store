import socket
import sys

#in bytes
DATA_TRANSMISSION_RATE = 16

class Server(object):
    """
    Server class
    """
    def __init__(self, host='localhost', port=8080):
        self.host = host
        self.port = port

    def openSocket(self):
        """
        Opens a socket to listen to given Host IP/Name and port
        """
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_address = (self.host, self.port)
        print >> sys.stderr, 'starting up on %s port %s' % server_address
        sock.bind(server_address)
        sock.listen(1)
        while True:
            print >> sys.stderr, 'waiting for a connection'
            connection, client_address = sock.accept()
            try:
                print >> sys.stderr, 'connection from', client_address
                while True:
                    data = connection.recv(DATA_TRANSMISSION_RATE)
                    print >>sys.stderr, 'received "%s"' % data
                    #TODO: Add checks for receiving and sending data
            finally:
                connection.close()


if __name__ == "__main__":
    host = None
    port = None
    try:
        host = sys.argv[1]
        port = sys.argv[2]
    except Exception:
        print >> sys.stderr, "Running with default Host(Name/IP) and port"
    finally:
        server = Server(host , int(port)) if host and port else Server()
        server.openSocket()
