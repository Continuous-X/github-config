package output

import "fmt"

const (
	MessagePrefixCli                      = "\n>> ghc"
	MessageCommandCheck                   = "CHECK"
	MessagePrefixCommandCheck             = MessagePrefixCli + " " + MessageCommandCheck
	MessageCommandGhProtection            = "GHPROTECTION"
	MessagePrefixCommandCheckGhProtection = MessagePrefixCommandCheck + " " + MessageCommandGhProtection

	MessageStateSuccess = "SUCCESS"
	MessageStateError   = "ERROR"
)

func PrintCliInfo(msg string) {
	PrintYellow(fmt.Sprintf("%s [INFO]: ", MessagePrefixCli))
	PrintWhite(msg)
}

func PrintCliWarning(msg string) {
	PrintYellow(fmt.Sprintf("%s [WARNING]: ", MessagePrefixCli))
	PrintWhite(msg)
}

func PrintCliError(err error) {
	PrintRed(fmt.Sprintf("%s [ERROR]: ", MessagePrefixCli))
	PrintWhite(fmt.Sprintf("%s", err))
}

func PrintCheckGhProtectionSuccess() {
	PrintCommandSuccess(MessagePrefixCommandCheckGhProtection)
}

func PrintCheckGhProtectionError() {
	PrintCommandError(MessagePrefixCommandCheckGhProtection)
}

func PrintCommandSuccess(command string) {
	PrintGreen(fmt.Sprintf("%s: %s", command, MessageStateSuccess))
}

func PrintCommandError(command string) {
	PrintRed(fmt.Sprintf("%s: %s", command, MessageStateError))
}

func PrintRed(msg string) {
	getRed().Print(msg)
}

func PrintYellow(msg string) {
	getYellow().Print(msg)
}

func PrintGreen(msg string) {
	getGreen().Print(msg)
}

func PrintWhite(msg string) {
	getWhite().Print(msg)
}
