package nife

import "slices"

var cmds = make(map[string]*Cmd)

// GetCmd returns a command by name or a ErrorCommandNotFound error if the command is not found.
func GetCmd(name string) (*Cmd, error) {
	cmd, ok := cmds[name]
	if !ok {
		return nil, ErrorCommandNotFound
	}
	return cmd, nil
}

// RegisterCmd registers a command with the nife package
// and panics if the command is invalid or already registered.
func RegisterCmd(c *Cmd) {

	if c == nil {
		panic("RegisterCommand: Command is nil")
	}
	if c.Name == "" {
		panic("RegisterCommand: Command has no Name")
	}
	if c.Title == "" {
		panic("RegisterCommand: Command has no Title")
	}
	if c.Usage == "" {
		panic("RegisterCommand: Command has no Usage")
	}

	if c.CmdHandler == nil {
		panic("RegisterCommand: Command has no CmdHandler")
	}

	_, ok := cmds[c.Name]
	if ok {
		panic("RegisterCommand: Command already registered.  " + c.Name)
	}

	cmds[c.Name] = c

}

// GetCmds returns a slice of all registered commands.
func GetCmds() []*Cmd {
	//ordere by name
	var cmdList []*Cmd
	for _, cmd := range cmds {
		cmdList = append(cmdList, cmd)
	}
	slices.SortFunc(cmdList, func(a, b *Cmd) int {
		if a.Name == b.Name {
			return 0
		}
		if a.Name > b.Name {
			return 1
		}
		return -1

	})

	return cmdList

}
