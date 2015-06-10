package tmof

type Implement interface {
	Ulimit() int
	FileHandle() (int, int, int)
	Pgrep(string) (string, string)
	FileDescriptors(string) map[int][]string
}
