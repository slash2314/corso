package flags

import (
	"github.com/spf13/cobra"
)

const (
	BackupFN             = "backup"
	BackupIDsFN          = "backups"
	AWSAccessKeyFN       = "aws-access-key"
	AWSSecretAccessKeyFN = "aws-secret-access-key"
	AWSSessionTokenFN    = "aws-session-token"

	// Corso Flags
	CorsoPassphraseFN = "passphrase"
	SucceedIfExistsFN = "succeed-if-exists"
)

var (
	BackupIDFV           string
	BackupIDsFV          []string
	AWSAccessKeyFV       string
	AWSSecretAccessKeyFV string
	AWSSessionTokenFV    string
	CorsoPassphraseFV    string
	SucceedIfExistsFV    bool
)

// AddMultipleBackupIDsFlag adds the --backups flag.
func AddMultipleBackupIDsFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringSliceVar(
		&BackupIDsFV,
		BackupIDsFN, nil,
		"',' separated IDs of the backup to retrieve")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(BackupIDsFN))
	}
}

// AddBackupIDFlag adds the --backup flag.
func AddBackupIDFlag(cmd *cobra.Command, require bool) {
	cmd.Flags().StringVar(&BackupIDFV, BackupFN, "", "ID of the backup to retrieve.")

	if require {
		cobra.CheckErr(cmd.MarkFlagRequired(BackupFN))
	}
}

// ---------------------------------------------------------------------------
// storage
// ---------------------------------------------------------------------------

func AddAllStorageFlags(cmd *cobra.Command) {
	AddCorsoPassphaseFlags(cmd)
	// AddAzureCredsFlags is added by ProviderFlags
	AddAWSCredsFlags(cmd)
}

func AddAWSCredsFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&AWSAccessKeyFV, AWSAccessKeyFN, "", "S3 access key")
	fs.StringVar(&AWSSecretAccessKeyFV, AWSSecretAccessKeyFN, "", "S3 access secret")
	fs.StringVar(&AWSSessionTokenFV, AWSSessionTokenFN, "", "S3 session token")
}

// M365 flags
func AddCorsoPassphaseFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&CorsoPassphraseFV,
		CorsoPassphraseFN,
		"",
		"Passphrase to protect encrypted repository contents")
}

// ---------------------------------------------------------------------------
// Provider
// ---------------------------------------------------------------------------

func AddAllProviderFlags(cmd *cobra.Command) {
	AddAzureCredsFlags(cmd)
}
