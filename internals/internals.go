package internals

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// SSHHost stores the information about the ssh host
type SSHHost struct {
	Host         string
	HostName     string
	Port         string
	User         string
	IdentityFile string
}

func (ssh *SSHHost) configValidator() (bool, error) {
	if len(ssh.Host) != 0 && len(ssh.HostName) != 0 && len(ssh.IdentityFile) != 0 && len(ssh.User) != 0 {
		return true, nil
	} else {
		fmt.Print(ssh)
		return false, errors.New("Invalid config")
	}
}

// ConfigParser reads the ssh config file and parses configs for the ssh host
func ConfigParser() ([]SSHHost, error) {
	configMap := []SSHHost{}
	rawConfig, err := os.Open("config")
	if err != nil {
		log.Fatalf("Error parsing ssh config file %v", err)
	}
	defer rawConfig.Close()

	var sshHost *SSHHost
	scanner := bufio.NewScanner(rawConfig)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Host") {
			sshHost = new(SSHHost)
			sshHost.Host = strings.Split(scanner.Text(), " ")[1]
			scanner.Scan()
			for scanner.Text() != "" {
				keyVal := strings.Split(strings.TrimSpace(scanner.Text()), " ")
				switch keyVal[0] {
				case "HostName":
					sshHost.HostName = keyVal[1]
				case "User":
					sshHost.User = keyVal[1]
				case "Port":
					sshHost.Port = keyVal[1]
				case "IdentityFile":
					sshHost.IdentityFile = keyVal[1]
				}
				scanner.Scan()
			}
			validated, err := sshHost.configValidator()
			if validated {
				configMap = append(configMap, *sshHost)
			} else if err != nil {
				log.Fatal(err)
			}

		}
	}
	return configMap, nil

}
