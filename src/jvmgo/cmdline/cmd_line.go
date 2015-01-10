package cmdline 

import "strings"

// java [ options ] class [ arguments ]
type Command struct {
    options     []*Option
    class       string
    args        []string
}

type Option struct {
    name    string
    value   string
}

type CmdLineArgs struct {
    args []string
}

func (self *CmdLineArgs) empty() (bool) {
    return len(self.args) > 0
}

func (self *CmdLineArgs) first() (string) {
    return self.args[0]
}

func (self *CmdLineArgs) removeFirst() (first string) {
    first = self.args[0]
    self.args = self.args[1:]
    return
}

func (self *Command) parseOptions(args *CmdLineArgs) {
    if !args.empty() {
        if strings.HasPrefix(args.first(), "-") {
            option := &Option{}
            option.name = args.removeFirst()
            if !args.empty() {
                if true {
                    option.value = args.removeFirst()
                }
            }
        }
    }
}

func (self *Command) parseClass(args *CmdLineArgs) {

}

func (self *Command) parseArgs(args *CmdLineArgs) {

}

func ParseCommand(osArgs []string) {
    cmdLineArgs := &CmdLineArgs{osArgs[1:]}
    cmd := &Command{}
    cmd.options = []*Option{} // len == 0
    cmd.parseOptions(cmdLineArgs)
    cmd.parseClass(cmdLineArgs)
    cmd.parseArgs(cmdLineArgs)
}
