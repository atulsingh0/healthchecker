package run

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"syscall"
	"time"
)

func Command(ctx context.Context, cmdArgs []string) ([]byte, error) {
	//#nosec:G204 // this is intentionally running a command
	c := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	type result struct {
		Output []byte
		Error  error
	}

	finished := make(chan result)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				finished <- result{Error: fmt.Errorf("command run panicked :%v", r)}
			}
		}()
		out, err := runAndCapture(c)
		finished <- result{
			Output: out,
			Error:  err,
		}
	}()

	// wait for the context being cancelled or the process to have finished.
	select {
	case <-ctx.Done():
		if c.Process == nil {
			return nil, errors.New("no process to terminate")
		}

		// Send a termination signal to the process we ran
		if err := syscall.Kill(c.Process.Pid, syscall.SIGTERM); err != nil {
			return nil, err
		}
		select {
		case <-time.After(time.Second * 10):
			// If the process didn't exit in reasonable time then kill the whole group.
			// (The negative pid means signal everything in the process group)
			if err := syscall.Kill(-c.Process.Pid, syscall.SIGKILL); err != nil {
				return nil, err
			}
		case res := <-finished:
			return res.Output, res.Error
		}
	case res := <-finished:
		return res.Output, res.Error
	}
	// If the context was cancelled before the process finished and
	// the timeout was hit and we sig-killed the group then we should still wait for completion.
	// Essentially we only get here if no other path has seen the process finished.
	res := <-finished
	return res.Output, res.Error
}

func runAndCapture(c *exec.Cmd) ([]byte, error) {
	// (see the os/exec.StdoutPipe documentation about the ordering of reading from the pipe and waiting)
	pipe, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}
	// push all stderr into the pipe as well
	c.Stderr = c.Stdout

	err = c.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start command %s: %w", append([]string{c.Path}, c.Args...), err)
	}

	b, _ := io.ReadAll(pipe)

	// return any exit error from the command
	err = c.Wait()
	if err != nil {
		return b, fmt.Errorf("failed to run command %s: %w\n%s", append([]string{c.Path}, c.Args...), err, b)
	}
	return b, nil
}
