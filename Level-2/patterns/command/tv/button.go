package tv

type Button struct {
	Command Command
}

func (b *Button) press() {
	b.Command.Execute()
}
