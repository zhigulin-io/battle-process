package model

import "errors"

type Model struct {
	Position Position
}

func (m *Model) AdvanceTo(position Position) error {
	return m.moveTo(position, 6)
}

func (m *Model) RushTo(position Position) error {
	return m.moveTo(position, 12)
}

func (m *Model) ChargeTo(position Position) error {
	return m.moveTo(position, 12)
}

func (m *Model) moveTo(position Position, limit float64) error {
	distance := m.Position.DistanceTo(position)
	if distance > limit {
		return errors.New("required position is too far")
	}

	m.Position = position
	return nil
}
