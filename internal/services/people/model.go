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
func (p *People) GetArrayFromStringStack() []string {
	if len(p.Stacks) == 0 {
		return []string{}
	}
	stacks := strings.Split(p.Stacks, ",")
	return stacks
}
