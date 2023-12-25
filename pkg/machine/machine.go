package machine

type Machine interface {
	Lockout()
	Unlockout()
	Cmd(command string)
}

func NewMachine() Machine {
	return &MachineImpl{}
}

type MachineImpl struct {
}

func (m *MachineImpl) Lockout() {

}

func (m *MachineImpl) Unlockout() {

}

func (m *MachineImpl) Cmd(command string) {

}
