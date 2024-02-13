package gocask

type CaskOpts struct {
	Role        string
	SyncOnWrite bool
}

func Open(dir string)
