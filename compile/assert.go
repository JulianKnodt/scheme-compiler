package compile

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var chtestdata sync.Once

var spaceReplacer = strings.NewReplacer(" ", "_")

func Assert(in, expected, testname string) error {
	chtestdata.Do(func() {
		if err := os.Chdir("./testdata"); err != nil {
			panic(err)
		}
	})
	testname = spaceReplacer.Replace(testname)
	filename := fmt.Sprintf("%s.s", testname)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Compile to file
	if err := Compile(in, dst); err != nil {
		return err
	}

	// Link to main
	if info, err := exec.Command("gcc",
		fmt.Sprintf("-o%s", testname),
		filename,
		"runtime.c").CombinedOutput(); err != nil {
		fmt.Println(string(info))
		return err
	}
	defer os.Remove(testname) // Clean up executables

	// Execute
	output, err := exec.Command(fmt.Sprintf("./%s", testname)).CombinedOutput()
	if err != nil {
		return err
	}
	if expected != string(output) {
		return fmt.Errorf("Assertion Failed, expected: %s, got: %s", expected, output)
	}
	return nil
}
