package pkg

// CommitPipeline is a function used to make git add/commit/push
func CommitPipeline(commitMessage string, isNew bool, branchName string) {
	// git add command
	ShellRun("git", "add", ".")

	// make a git commit command
	ShellRun("git", "commit", "-m", commitMessage)

	if isNew {
		// make a git push command
		ShellRun("git", "push", "origin", branchName)
	} else {
		ShellRun("git", "push")
	}
}
