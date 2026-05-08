// retry-iam reruns a command while it fails with an IAM eventual-consistency
// permission error (AWS or GCP). Any other failure exits immediately.
//
//	go run ./tools/retry-iam tofu apply -auto-approve
//
// Env tunables:
//
//	RETRY_IAM_ATTEMPTS  max attempts (default 12)
//	RETRY_IAM_SLEEP     base sleep seconds, linear backoff capped at 30 (default 5)
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

var iamDenied = regexp.MustCompile(`(?i)Error 403|PERMISSION_DENIED|AccessDenied|AccessDeniedException|UnauthorizedOperation|AuthorizationFailed|is not authorized to perform|does not have .* (permission|authorization)`)

func envInt(key string, def int) int {
	if v, err := strconv.Atoi(os.Getenv(key)); err == nil && v > 0 {
		return v
	}
	return def
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "retry-iam: no command given")
		os.Exit(2)
	}

	attempts := envInt("RETRY_IAM_ATTEMPTS", 12)
	base := time.Duration(envInt("RETRY_IAM_SLEEP", 5)) * time.Second

	for i := 1; ; i++ {
		fmt.Printf("[retry-iam] attempt %d/%d: %v\n", i, attempts, args)

		var buf bytes.Buffer
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &buf)

		err := cmd.Run()
		if err == nil {
			fmt.Printf("[retry-iam] success on attempt %d\n", i)
			os.Exit(0)
		}

		rc := 1
		var ee *exec.ExitError
		if errors.As(err, &ee) {
			rc = ee.ExitCode()
		} else {
			fmt.Fprintf(os.Stderr, "[retry-iam] failed to run %q: %v\n", args[0], err)
			os.Exit(rc)
		}

		if !iamDenied.Match(buf.Bytes()) {
			fmt.Printf("[retry-iam] non-IAM failure (exit %d), not retrying\n", rc)
			os.Exit(rc)
		}
		if i >= attempts {
			fmt.Printf("[retry-iam] still IAM-denied after %d attempts (exit %d)\n", attempts, rc)
			os.Exit(rc)
		}

		sleep := min(base*time.Duration(i), 30*time.Second)
		fmt.Printf("[retry-iam] IAM denial detected, sleeping %v\n", sleep)
		time.Sleep(sleep)
	}
}
