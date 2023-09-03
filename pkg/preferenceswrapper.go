package punchclock

import "fyne.io/fyne/v2"

type PreferencesWrapper interface {
	SetBool(string, bool)
	GetBool(string) bool
	SetString(string, string)
	GetString(string) string
}

type PrefWrapperImpl struct {
	pref fyne.Preferences
}

func NewPreferencesWrapper(p fyne.Preferences) *PrefWrapperImpl {
	handler := new(PrefWrapperImpl)
	handler.pref = p
	return handler
}

func (p *PrefWrapperImpl) SetBool(key string, value bool) {
	p.pref.SetBool(key, value)
}

func (p *PrefWrapperImpl) GetBool(key string) bool {
	return p.pref.BoolWithFallback(key, false)
}

func (p *PrefWrapperImpl) SetString(key string, value string) {
	p.pref.SetString(key, value)
}

func (p *PrefWrapperImpl) GetString(key string) string {
	return p.pref.StringWithFallback(key, "")
}
