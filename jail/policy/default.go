package policy

import (
	"strings"
)

type DefaultPolicy struct {
}

func (p *DefaultPolicy) CheckPathOp(path string, op PathOps, mode int) bool {
	var allow bool
	switch op {
	case OP_CHMOD:
		allow = strings.HasPrefix(path, "/home/work/")
	default:
		allow = true
	}
	return allow
}

func (p *DefaultPolicy) GetExecAllowance() int {
	return 0
}

func (p *DefaultPolicy) GetForkAllowance() int {
	return -1
}

func (p *DefaultPolicy) GetMaxChildProcs() uint {
	return 32
}

func (p *DefaultPolicy) CheckPathExecutable(path string) bool {
	return true
}

func (p *DefaultPolicy) GetExtraEnvs() []string {
	return []string{}
}

func (p *DefaultPolicy) GetPreservedEnvKeys() []string {
	return []string{"HOME", "PATH", "PYENV_ROOT", "PYTHONPATH"}
}

// vim: ts=4 sts=4 sw=4 noet
