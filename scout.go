package main

import (
	"fmt"
	"log"

	"github.com/OmairK/scout/internals"
)

func main() {
	configMap, err := internals.ConfigParser()
	if err != nil {
		log.Fatalf("error %v", err)
	}
	var hostChoice int
	for index, configs := range configMap {
		fmt.Printf("%v %v\n", index, configs.Host)
	}
	fmt.Scanln(&hostChoice)

	if hostChoice < 0 && hostChoice >= len(configMap) {
		fmt.Println("Invalid input")
		return
	}
}

/*
// An SSH client is represented with a ClientConn.
		// To authenticate with the remote server you must pass at least one
		// implementation of AuthMethod via the Auth field in ClientConfig,
		// and provide a HostKeyCallback.
key, err := ioutil.ReadFile("openvpn.pem")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "13.233.164.20:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	for a := 0; a < 10; a++ {
		var command string
		fmt.Scanln(&command)
		if err := session.Run(command); err != nil {
			log.Fatal("Failed to run: " + err.Error())
		}
		fmt.Println(b.String())
	}
*/
