package main

// Field injected by goreleaser
var (
	version    = "v0.0.8"
	commitDate = "date unknown"
	commit     = ""
	target     = ""
)

func Version() string {
	return version
}

func CommitDate() string {
	return commitDate
}

func Commit() string {
	return commit
}

func Target() string {
	return target
}
