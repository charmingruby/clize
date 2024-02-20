package terminal

import (
	"fmt"
	"os"
	"os/exec"
)

const FetchCmd = "ftc"
const GetCmd = "get"
const CreateCmd = "crt"
const AddCmd = "add"
const ModifyCmd = "mod"
const DelCmd = "del"
const RemoveCmd = "rm"
const SubmitCmd = "sub"

func CommandWrapper(actor, action string) string {
	wp := fmt.Sprintf("%s-%s", actor, action)
	return wp
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
