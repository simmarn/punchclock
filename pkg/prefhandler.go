package punchclock

import "fyne.io/fyne/v2"

type PrefHandler interface {
	SetBool(string, bool)
	GetBool(string) bool
	SetString(string, string)
	GetString(string) string
}

type PrefHandlerImpl struct {
	pref fyne.Preferences
}

func NewPrefHandler(p fyne.Preferences) *PrefHandlerImpl {
	handler := new(PrefHandlerImpl)
	handler.pref = p
	return handler
}

func (p *PrefHandlerImpl) SetBool(key string, value bool) {
	p.pref.SetBool(key, value)
}

func (p *PrefHandlerImpl) GetBool(key string) bool {
	return p.pref.BoolWithFallback(key, false)
}

func (p *PrefHandlerImpl) SetString(key string, value string) {
	p.pref.SetString(key, value)
}

func (p *PrefHandlerImpl) GetString(key string) string {
	return p.pref.StringWithFallback(key, "")
}
