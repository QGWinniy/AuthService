package terminal

import (
	"io"

	"golang.org/x/crypto/ssh"
)

type Terminal interface {
    // Connect()
    Write([]byte) error
    Read([]byte) (int, error)
    Close() error
    Resize(width, height int) error
}

type SSHTerminal struct {
    Client  *ssh.Client
    Session *ssh.Session
    Stdin   io.WriteCloser
    Stdout  io.Reader
}

func (t *SSHTerminal) Write(data []byte) error {
    _, err := t.Stdin.Write(data)
    return err
}

func (t *SSHTerminal) Read(data []byte) (int, error) {
    return t.Stdout.Read(data)
}

func (t *SSHTerminal) Close() error {
    t.Client.Close()
    return t.Session.Close()
}

func (t *SSHTerminal) Resize(width, height int) error {
    return t.Session.WindowChange(height, width)
}
