package branch

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Delete = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Try to delete the checked out branch first (to no avail), and then delete another branch.",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			EmptyCommit("blah").
			NewBranch("branch-one").
			NewBranch("branch-two")
	},
	Run: func(shell *Shell, input *Input, keys config.KeybindingConfig) {
		input.Views().Branches().
			Focus().
			Lines(
				MatchesRegexp(`\*.*branch-two`).IsSelected(),
				MatchesRegexp(`branch-one`),
				MatchesRegexp(`master`),
			).
			Press(keys.Universal.Remove).
			Tap(func() {
				input.ExpectAlert().Title(Equals("Error")).Content(Contains("You cannot delete the checked out branch!")).Confirm()
			}).
			SelectNextItem().
			Press(keys.Universal.Remove).
			Tap(func() {
				input.ExpectConfirmation().
					Title(Equals("Delete Branch")).
					Content(Contains("Are you sure you want to delete the branch 'branch-one'?")).
					Confirm()
			}).
			Lines(
				MatchesRegexp(`\*.*branch-two`),
				MatchesRegexp(`master`).IsSelected(),
			)
	},
})
