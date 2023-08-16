package people

import "strings"

type People struct {
	Id       string
	Nickname string
	Name     string
	Birthday string
	Stacks   string
}

func (p *People) SetStacksFromArray(stacks []string) {
	p.Stacks = strings.Join(stacks, ",")
}
