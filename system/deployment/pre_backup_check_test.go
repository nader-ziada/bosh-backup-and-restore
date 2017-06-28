package deployment

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	. "github.com/pivotal-cf/bosh-backup-and-restore/system"
)

var _ = Describe("pre-backup-check", func() {
	It("backs up, and cleans up the backup on the remote", func() {
		preBackupCheckCommand := RunCommandOnRemoteAsVcap(
			JumpBoxSSHCommand(),
			fmt.Sprintf(`cd %s; \
			    BOSH_CLIENT_SECRET=%s ./bbr deployment \
			       --ca-cert bosh.crt \
			       --username %s \
			       --target %s \
			       --deployment %s \
			       pre-backup-check`,
				workspaceDir,
				MustHaveEnv("BOSH_CLIENT_SECRET"),
				MustHaveEnv("BOSH_CLIENT"),
				MustHaveEnv("BOSH_URL"),
				RedisDeployment()),
		)

		Eventually(preBackupCheckCommand).Should(gexec.Exit(0))
		Expect(preBackupCheckCommand.Out.Contents()).To(ContainSubstring("Deployment 'redis-dev-1' can be backed up"))
	})
})