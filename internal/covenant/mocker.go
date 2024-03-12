package covenant

import "errors"

type Mocker struct {
	forceErrors bool
}

func NewMocker() Database {
	return &Mocker{
		forceErrors: false,
	}
}

func NewMockerError() Database {
	return &Mocker{
		forceErrors: true,
	}
}

func (m *Mocker) Create(destiny interface{}) error {

	if m.forceErrors {
		return errors.New("Mock create error")

	}

	return nil
}

func (m *Mocker) Read(destiny interface{}, conditionals ...interface{}) error {

	if m.forceErrors {
		return errors.New("Mock read error")

	}

	return nil
}

func (m *Mocker) Update(destiny interface{}) error {

	if m.forceErrors {
		return errors.New("Mock update error")
	}

	return nil
}

func (m *Mocker) Delete(destiny interface{}, conditionals ...interface{}) error {

	if m.forceErrors {
		return errors.New("Mock delete error")
	}

	return nil
}
