package list

import (
	"fmt"

	"gitoa.ru/go-4devs/console/param"
)

const tpl = `
The <info>%[2]s</info> command lists all commands:
  <info>%[1]s %[2]s</info>
You can also display the commands for a specific namespace:
  <info>%[1]s %[2]s test</info>
You can also output the information in other formats by using the <comment>--format</comment> option:
  <info>%[1]s %[2]s --format=xml</info>
`

func Help(data param.HData) (string, error) {
	return fmt.Sprintf(tpl, data.Bin, data.Name), nil
}
