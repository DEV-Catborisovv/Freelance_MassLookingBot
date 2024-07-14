package pyrunner

import (
	"fmt"
	"os/exec"
)

type PyRunner struct {
	API_ID   string
	API_HASH string
}

func NewPyRunner(api_id string, api_hash string) *PyRunner {
	return &PyRunner{
		API_ID:   api_id,
		API_HASH: api_hash,
	}
}

// __________ Logic of Pyrunner __________

func (s *PyRunner) Run() error {
	cmd := exec.Command("python3.11", "./python/masslook_script/masslookScript.py", s.API_ID, s.API_HASH)

	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
	}

	fmt.Println(string(out))

	return nil
}
