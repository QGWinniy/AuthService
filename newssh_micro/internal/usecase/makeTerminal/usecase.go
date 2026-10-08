package maketerminal

import (
	terminal "ssh/internal/Terminal"
	"ssh/internal/domain"


	"golang.org/x/crypto/ssh"
)

type Config struct {
	// RepoUserConnect repository.UserServerSessionRepository по идеи нахуй не надо
}

type UseCase struct {
	// config Config
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) MakeTerminal(server *domain.UserServer) (terminal.Terminal, error) {
    config := &ssh.ClientConfig{
        User: server.UserName,
        Auth: []ssh.AuthMethod{
            ssh.Password(server.Password), // тут в будущем разобраться как делать конект по ssh-ключу
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), // эту хуету в будущем обезопасить
    }

    client, err := ssh.Dial(
        "tcp",
        server.Addr,
        config,
    )

    if err != nil {
        return nil, err
    }

    session, err := client.NewSession()
    if err != nil {
        client.Close()
        return nil, err
    }

    stdin, err := session.StdinPipe()
    if err != nil {
        session.Close()
        client.Close()
        return nil, err
    }

    stdout, err := session.StdoutPipe()
    if err != nil {
        session.Close()
        client.Close()
        return nil, err
    }


    if err := session.RequestPty(
        "xterm-256color",
        120,
        30,
        ssh.TerminalModes{
            ssh.ECHO:          1,
            ssh.TTY_OP_ISPEED: 14400,
            ssh.TTY_OP_OSPEED: 14400,
        },
    ); err != nil {
        session.Close()
        client.Close()
        return nil, err
    }

    
    
    if err := session.Shell(); err != nil {
        session.Close()
        client.Close()
        return nil, err
    }

    // if err := session.Start(""); err != nil {
    //     session.Close()
    //     client.Close()
    //     return nil, err
    // }

    return &terminal.SSHTerminal{
        Client: client,
        Session: session,
        Stdin:   stdin,
        Stdout:  stdout,
    }, nil
}

