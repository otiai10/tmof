package platforms

import "github.com/otiai10/tmof"

type Debian struct{}

func (deb Debian) Ulimit() int {
	return 1024
}

func (deb Debian) FileHandle() (int, int, int) {
	return 1600, 0, 30000
}

func (deb Debian) Pgrep(pname string) [][]string {
	return [][]string{[]string{"798", "/bin/foo"}}
}

func (deb Debian) FileDescriptors(pid string) []tmof.FileDescriptor {
	return []tmof.FileDescriptor{
		FileDescriptor{},
	}
}
