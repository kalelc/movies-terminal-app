package domain

type Movie struct {
	Id         int64
	Name       string
	Overview   string
	Popularity float32
}

func (m Movie) Title() string {
	return m.Name
}

func (m Movie) Description() string {
	return m.Overview
}

func (m Movie) FilterValue() string {
	return m.Name
}
