// +build darwin

package diskcache

import "syscall"

func AtimeNano(s *syscall.Stat_t) int64 {
	return s.Atimespec.Nano()
}
