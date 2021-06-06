package pkg

// CommitPipeline is a function used to make git add/commit/push
func CommitPipeline(commitMessage string) {
	// git add command
	errout, out, err := Shellout("git", "add", ".")
	Output(out, errout, err)

	// make a git commit command
	errout, out, err = Shellout("git", "commit", "-m", commitMessage)
	Output(out, errout, err)

	// make a git push command
	errout, out, err = Shellout("git", "push")
	Output(out, errout, err)
}
