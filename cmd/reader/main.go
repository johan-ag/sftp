package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var base64Key = []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
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
dpvUVSBAOYZxkLIU1bhik8qx/l0=
-----END OPENSSH PRIVATE KEY-----`)

var (
	theFiles []remoteFiles
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

	// read one file
	readSpecificFile(client, "unencryptedBR_20230112_093014_0000001.meli")
	// read dir
	//reaDir(client, "./")

}

func sshConfig() (*ssh.ClientConfig, error) {
	privateKey, err := ssh.ParsePrivateKeyWithPassphrase(base64Key, []byte("password"))
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


func readSpecificFile(client *sftp.Client, file string) {
		// check it's there
		srcFile, err := client.OpenFile(file, (os.O_RDONLY))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open remote file: %v\n", err)
			return
		}
		defer srcFile.Close()
	
		dstFile, err := os.Create("./temp")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open local file: %v\n", err)
			return
		}
		defer dstFile.Close()
	
		bytes, err := io.Copy(dstFile, srcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to download remote file: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d bytes copied\n", bytes)
}

func reaDir(client *sftp.Client, dir string) {

	fi, err := client.ReadDir(dir)
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
		})
	}

}