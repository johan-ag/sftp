package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)


var base64Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA57ns63oUCeC5M8UBBmxRZiKxbAvzh+OEyqtnQcbffiedfAHO
S+J4nhOzOS9kh8xwJg1Nw0SpF3lYpwN65OZMxPgX24fIVsreuLVogy30PHjgKHft
os0d2UJGRD1D8LoicV7jCmhEavlrReTjCNkrok1hXVTfV9hicx246mtjF4Ga09wP
XEMh7wcM55jF0JwVDeyhN7qAcX2EF2Tdl17Cq50mUSDXCnrUgNwGCcTxJyGjQSmj
Pq+dN6coLae+miTcBN8PYKR+Ppu3TZUnycqZigg4mDwZwSCr+EPUDezYWZk6fzWc
SvDcDPVORcadgqzTH8P1ndnrstO0RTOICGJvfwIDAQABAoIBAGx8VPjhTGRbexlL
j/FL4MfqUhn9dmQWFmMz38Ghs5xCO66Eweow+rs3Cd8p2uzgv1hxPgi/KlK9Es4O
CQkE8Mf+Rl0Wsqo/jAn5lBZl+0QcawVHME/Zq7G9H1xvOlGHMvzUqYKD0hQq9Rfh
0pCf65eulni1dWKRAXZXYe0NnXSw6o827kmWG8/SBAAjLoX1stBN15ydxEvqDAH6
4uBJBp/xPlFg7OSbuVg6vjtk5524e7/eSWbkcy8f03bbzu574NFnjr8KMk7VcUml
KhR7cBQDyB1hPO50Qdg6YQKri3Uahqwwh3SYhqgY2QmBq9IGOikDbYmtoK5vpNIE
VSh4h4ECgYEA95iHU74NiK6l4aOj2U5cYmcokzf/hTLgbSy3s62AScmZuhQhdqtM
8YZ9S8gNA3/iZKL0lYYH+NgBdK5qPqTSgfiRb03qjNJW9l0sLupPmsT+9j4gyDNU
orvZmsEs304g6R8OtA4d7hVLAvj1LtfrmkEOFg3JyhjUCbd2VtiWuB8CgYEA75d/
zpZQHqrHf0Bb7dnWseJAKK2ZxBzE1t7e/CFnvAIBCMsc7Wq0fKKYE/dvC+KvuLMM
pJ1jlzG3Im21O6eaJXjq87R38z/7h1vbGVUuGGwxsQM0YVv9/tGuWQfCiqchHKNC
IIxBdLIRC401YNGvlQYLvdhcAWT/QVfQDkDi3KECgYBOoW0QxkGsD7L0lrB5Wa5z
Pcmf/1+xoHevlEz+zfH4/QJKGwyJkFtONOTjxTOE0f9G8I12UuuDNq43rRBmtpd3
2UXusDL15/LgKLTYbWc8gqVp9E8a5VpSmA/CWujyPfCruEZGHlmY+8AnwIK1DrNu
+IwNqvFTslDciRwW9o/fBwKBgQCJTjNKe04TJSALRyDn4G1vqA/IwCRdBXexBCya
JbLDRgq9tVkwnnQs8SliV+zlzNWTbdkTBberUCx5qJGyTEzEqNVNMOGFH+fQHM4l
aHFwQaEApquAGC4puJTIJPBScZFCTjV5Xaetbtbh3dz3TXPoXXyhIbsoeW7nTaHI
aVsZwQKBgQCIUis3o9ETjWSQKovd9Ygc0/M+AObs0tIVm2GMwMYGthAIK8xAyn/K
QZrKJ3v4UVVVwEGMUPsVt6x8gRNK71OEJxRzOCEHrw+bZPscdoLMDltAh2OSnvq/
lT+wJ0mKYqw3jEeZ1XPJo7a5aN/z4DaNj36hvsA1411bPe1GZA37kQ==
-----END RSA PRIVATE KEY-----`)


var (
	theFiles []remoteFiles
	path = "file.txt"
)

func main() {

	config, err := sshConfig()
	if err != nil {
		log.Fatal("fail ssh config::", err)
	}

	conn, err := ssh.Dial("tcp", "127.0.0.1:2222", config)
	if err != nil {
		log.Fatal("fail ssh connection::", err)
	}

	// open an SFTP session over an existing ssh connection.
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// walk a directory
	w := client.Walk("/home/user")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	// leave your mark
	f, err := client.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal(err)
	}
	f.Close()

	// check it's there
	// fi, err := client.Lstat("hello.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(fi)

	fi, err := client.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fi {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}

		theFiles = append(theFiles, remoteFiles{
			Name:    name,
			Size:    size,
			ModTime: modTime,
		})

		log.Printf("%+v --- ", remoteFiles{
			Name:    name,
			Size:    size,
			ModTime: modTime,
	 } )
	}
}

func sshConfig() (*ssh.ClientConfig, error) {
	privateKey, err := ssh.ParsePrivateKey(base64Key)
	if err != nil {
		return nil, err
	}

	sshAuthMethod := ssh.PublicKeys(privateKey)
	sshConfig := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			sshAuthMethod,
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return sshConfig, err
}

type remoteFiles struct {
    Name    string
    Size    string
    ModTime string
}