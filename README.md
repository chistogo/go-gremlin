# go-gremlin

## Project Description

The goal of this project is to create a one way file transport program using UDP with error
detection and correction capabilities for reliable data transfer through a gremlin proxy
server. The sending client should prompt the user for a file to send. There should be no
restrictions on the file type. Once the user has selected a file to send, the client program
will send the file to the destination (using UDP) through a proxy server. The proxy server
will forward (or drop) data received from the sending host to the destination host. The
drop rate percentage on the proxy should be configurable between 0% and 100%, where
this drop rate is unknown to the sending/receiving hosts. You can assume the proxy does
not modify received packets, only forward to receiving host or drop. Error detection and
correction should be implemented to ensure successful file transfer over UDP through the
proxy. The proxy should print to the screen information on its current drop rate and how
it handles each packet it receives in real time.
