package platforms

import "github.com/otiai10/tmof"

func GetPlatform() tmof.Platform {
	return Debian{}
}
