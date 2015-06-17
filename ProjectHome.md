The goal of this project is to have a high performance Go wrapper for libpcap.

For now we do this by not copying from C memory but accessing it directly.

To install:
```
go get code.google.com/p/golibpcap/...
```

For usage see [golibpcap/example/tcpdump.go](http://code.google.com/p/golibpcap/source/browse/example/tcpdump.go)

## Hints ##
_Remember that because this uses cgo you cannot cross compile - use a VM or better yet compile on the system where it is intended to run._

_Disable TCP Segmentation Offload (TSO)_
```
sudo ethtool -K eth0 tso off
sudo ethtool -K eth0 gso off
```