package pid

import (
	"os"
	"fmt"
	"bufio"
	"path/filepath"
	"bytes"
)

// comment for /proc/pid/cmdline
// This read-only file holds the complete command line for the process, unless the process is a zombie.  In the latter
// case, there is nothing in this file: that is, a read on this file will return 0 characters.  The command-line argu‐
// ments  appear  in  this file as a set of strings separated by null bytes ('\0'), with a further null byte after the
// last string.

type CmdlineInfo struct {
	Cmdline string
}

func Cmdline(pid string) (CmdlineInfo, error) {
	var ci CmdlineInfo

	f, err := os.Open(filepath.Join(Ctx.Procfs, pid, "cmdline"))
	if err != nil {
		return ci, fmt.Errorf("Cmdline: open %v/%v/cmdline failed: %v\n", Ctx.Procfs, pid, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	cmdline, _ := r.ReadBytes('\n')
	if len(cmdline) > 1 {
		cmdline := bytes.Replace(cmdline[:len(cmdline)-1], []byte{0}, []byte{' '}, -1) 

		ci.Cmdline = string(cmdline)
	}
	
	return ci, nil
}
