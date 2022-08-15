package main

import (
	"github.com/thpiron/webex-helper/cmd"
	_ "github.com/thpiron/webex-helper/cmd/memberships"
	_ "github.com/thpiron/webex-helper/cmd/messages"
	_ "github.com/thpiron/webex-helper/cmd/people"
	_ "github.com/thpiron/webex-helper/cmd/rooms"
	_ "github.com/thpiron/webex-helper/cmd/teamMemberships"
	_ "github.com/thpiron/webex-helper/cmd/teams"
)

func main() {
	cmd.InitConfig()
	cmd.Execute()
}
