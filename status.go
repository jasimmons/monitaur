package monitaur

const (
	OK      Status = 0
	Warning Status = 1
	Error   Status = 2
)

type Status int

func (s Status) IsOK() bool {
	return s == 0
}

func (s Status) IsWarning() bool {
	return s == 1
}

func (s Status) IsCritical() bool {
	return s == 2
}

func (s Status) IsUnknown() bool {
	return !(s.IsOK() || s.IsWarning() || s.IsCritical())
}
