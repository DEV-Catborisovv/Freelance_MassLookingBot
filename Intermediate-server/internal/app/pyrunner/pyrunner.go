package pyrunner

import (
	"fmt"
	"os/exec"
)

type PyRunner struct {
	API_ID   string
	API_HASH string
	Chats    []string
}

func NewPyRunner(api_id string, api_hash string, chats []string) *PyRunner {
	return &PyRunner{
		API_ID:   api_id,
		API_HASH: api_hash,
		Chats:    chats,
	}
}

// __________ Logic of Pyrunner __________

func (s *PyRunner) Run() error {
	args := append([]string{"./python/masslook_script/masslookScript.py", s.API_ID, s.API_HASH}, s.Chats...)
	cmd := exec.Command("python3.11", args...)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(out))
	return nil
}
