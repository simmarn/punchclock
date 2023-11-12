package punchclock

import "fyne.io/fyne/v2"

type PreferencesWrapper interface {
	SetBool(string, bool)
	GetBool(string) bool
	SetString(string, string)
	GetString(string) string
	SetFloat(string, float64)
	GetFloatWithFallback(string, float64) float64
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

func (p *PrefWrapperImpl) GetFloatWithFallback(key string, fallback float64) float64 {
	return p.pref.FloatWithFallback(key, fallback)
}

func (p *PrefWrapperImpl) SetFloat(key string, value float64) {
	p.pref.SetFloat(key, value)
}
