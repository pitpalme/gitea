	"code.gitea.io/gitea/models/db"
	diff := `diff --git a/newfile2 b/newfile2
	diff2 := `diff --git "a/A \\ B" "b/A \\ B"
	diff2a := `diff --git "a/A \\ B" b/A/B
	diff3 := `diff --git a/README.md b/README.md

	assert.NoError(t, diff.LoadComments(db.DefaultContext, issue, user))
	gitRepo, err := git.OpenRepository(git.DefaultContext, "./testdata/academic-module")