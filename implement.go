package tmof

type Platform interface {
	Ulimit() int
	FileHandle() (int, int, int)
	Pgrep(string) [][]string
	FileDescriptors(string) []FileDescriptor
}

type FileDescriptor interface {

}
