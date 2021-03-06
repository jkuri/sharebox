package vexclient

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type endpoint struct {
	Host string
	Port int
}

func (e *endpoint) String() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

// InitClient initializes vex connection to the remote server.
func InitClient(localServer, remoteServer string, localPort, remotePort int) {

	// local service to be forwarded
	var localEndpoint = endpoint{
		Host: localServer,
		Port: localPort,
	}

	// remote SSH server
	var serverEndpoint = endpoint{
		Host: remoteServer,
		Port: remotePort,
	}

	// remote forwarding port (on remote SSH server network)
	var remoteEndpoint = endpoint{
		Host: "localhost",
		Port: randomPort(11000, 65000),
	}

	sshConfig := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH remote server using serverEndpoint
	serverConn, err := ssh.Dial("tcp", serverEndpoint.String(), sshConfig)
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}

	connectionDone := make(chan struct{})
	go func() {
		err = keepAliveTicker(serverConn, connectionDone)
		if err != nil {
			log.Fatalln(fmt.Printf("Cannot initialize keep alive on connection: %s", err))
		}
	}()

	go func() {
		session, err := serverConn.NewSession()
		if err != nil {
			log.Fatalln(fmt.Printf("Cannot create session error: %s", err))
		}

		stdout, err := session.StdoutPipe()
		if err != nil {
			log.Fatalln(fmt.Printf("Unable to setup stdout for session: %s", err))
		}

		go io.Copy(os.Stdout, stdout)
	}()

	// Listen on remote server port
	listener, err := serverConn.Listen("tcp", remoteEndpoint.String())
	if err != nil {
		log.Fatalln(fmt.Printf("Listen open port ON remote server error: %s", err))
	}
	defer listener.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			// Open a (local) connection to localEndpoint whose content will be forwarded to serverEndpoint
			local, err := net.Dial("tcp", localEndpoint.String())
			if err != nil {
				log.Fatalln(fmt.Printf("Dial INTO local service error: %s", err))
			}

			client, err := listener.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			go handleClient(client, local)
		}
	}()

	<-c
}

func handleClient(client net.Conn, remote net.Conn) {
	defer client.Close()
	chDone := make(chan bool)

	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}

func keepAliveTicker(cl *ssh.Client, done <-chan struct{}) error {
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			_, _, err := cl.SendRequest("keepalive", true, nil)
			if err != nil {
				return errors.Wrap(err, "failed to send keep alive")
			}
		case <-done:
			return nil
		}
	}
}

func randomPort(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
