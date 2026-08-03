package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type SshConnection struct {
	client *ssh.Client
}

// hostKeyCallback verifies the device against the user's known_hosts file.
//
// This connection carries the router password and returns the full exported
// configuration, so accepting any host key would let anything on the path
// between here and the device collect both. There is deliberately no flag to
// turn the check off: record the key once with ssh or ssh-keyscan instead.
func hostKeyCallback() (ssh.HostKeyCallback, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot locate the home directory to read known_hosts: %v", err)
	}

	path := filepath.Join(home, ".ssh", "known_hosts")
	cb, err := knownhosts.New(path)
	if err != nil {
		return nil, fmt.Errorf("cannot verify the host key from %v: %v\n"+
			"connect once with ssh, or run: ssh-keyscan -H <host> >> %v", path, err, path)
	}

	return cb, nil
}

func NewSsh(host, username, password string) (*SshConnection, error) {
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	callback, err := hostKeyCallback()
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: callback,
		Timeout:         10 * time.Second,
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial, %v", err)
	}

	conn := &SshConnection{
		client: client,
	}

	return conn, nil
}

func (c *SshConnection) Close() error {
	return c.client.Close()
}

func (c *SshConnection) Run(cmd string) (string, error) {
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session, %v", err)
	}
	defer func() { _ = session.Close() }()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("failed to run, %v", err.Error())
	}
	return b.String(), nil
}

func GetMikrotikConfig(conn *SshConnection) (string, error) {
	return conn.Run("/export terse")
}

func GetResourceId(conn *SshConnection, path string, requiredFields []string) string {
	var id string
	var filter = strings.Join(requiredFields, " ")

	cmd := fmt.Sprintf(":put [%v get [ find %v ]]", path, filter)
	log.Debug("Searching id in ", path, " with command: ", cmd)
	res, err := conn.Run(cmd)
	if err != nil {
		log.Error("Error running command: ", err)
		return "?"
	}

	ss := reId.FindStringSubmatch(res)
	log.Debug("ss is ", ss)
	if len(ss) == 2 {
		id = ss[1]
		log.Info("Found id ", id, " for ", filter)
		return id
	}

	// Let's try with dynamic=no
	cmd = fmt.Sprintf(":put [%v get [ find %v dynamic=no ]]", path, filter)
	log.Debug("Searching id in ", path, " with command: ", cmd)
	res, err = conn.Run(cmd)
	if err != nil {
		log.Error("Error running command: ", err)
		return "?"
	}

	ss = reId.FindStringSubmatch(res)
	log.Debug("ss is ", ss)
	if len(ss) == 2 {
		id = ss[1]
		log.Info("Found id ", id, " for ", filter)
		return id
	}

	if id == "" {
		log.Warnf("Id not found for %v (filter: %v)", requiredFields, filter)
		return "?"
	}

	return id
}
