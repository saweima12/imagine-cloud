package modules

import "syscall"

type DiskStatus struct {
	All   uint64 `json:"all"`
	Usage uint64 `json:"usage"`
	Free  uint64 `json:"free"`
}

func ReadDiskInfo(path string) *DiskStatus {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)

	if err != nil {
		return nil
	}

	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)

	return &DiskStatus{
		All:   all,
		Usage: all - free,
		Free:  free,
	}

}
