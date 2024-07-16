package pyrunner

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	memorystorage "Freelance_MassLookingBot_Intermediate-server/internal/app/memoryStorage"
)

type PyRunner struct {
	API_ID      string
	API_HASH    string
	PhoneNumber string
	Chats       []string
	CodeChan    chan interface{}
}

func NewPyRunner(api_id string, api_hash string, phonenumber string, chats []string) *PyRunner {
	return &PyRunner{
		API_ID:      api_id,
		API_HASH:    api_hash,
		PhoneNumber: phonenumber,
		Chats:       chats,
		CodeChan:    make(chan interface{}),
	}
}

func (s *PyRunner) Run() error {
	args := append([]string{"./python/masslook_script/masslookScript.py", s.API_ID, s.API_HASH, s.PhoneNumber}, s.Chats...)
	cmd := exec.Command("python", args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Printf("stdout: %s\\n", text)
			if strings.Contains(text, "Enter the code you received: ") {
				memstorage := memorystorage.GetInstance()
				memstorage.Set(s.PhoneNumber, s.CodeChan)

				code := <-s.CodeChan
				fmt.Printf("Code received from channel: %s\\n", code)

				fmt.Fprintln(stdinPipe, code)

				fmt.Println("Success getted code: ", code)
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Printf("stderr: %s\\n", scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command finished with error: %w", err)
	}

	return nil
}
