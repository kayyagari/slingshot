package jnlp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunJar(dir string, mainClsName string, vmArgs []string, args []string) error {
	finalArgs := make([]string, 0)
	if vmArgs != nil {
		finalArgs = append(finalArgs, vmArgs...)
	}

	finalArgs = append(finalArgs, mainClsName)
	if args != nil {
		finalArgs = append(finalArgs, args...)
	}

	cmd := exec.Command("java", finalArgs...)
	cmd.Dir = dir
	cp := buildClasspath(dir)
	if cp == "" {
		return fmt.Errorf("no jar files found under %s", dir)
	}

	cp = "CLASSPATH=" + cp
	cmd.Env = []string{cp}
	return cmd.Start()
}

func buildClasspath(dir string) string {
	f, err := os.Open(dir)
	if err != nil {
		return ""
	}

	names, _ := f.Readdirnames(-1)

	return strings.Join(names, string(os.PathListSeparator))
}
