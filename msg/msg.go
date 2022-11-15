package lab1

import (
	"finis/rot"
	"finis/vigenere"
)

type MSG struct {
	Msg string `json:"msg"`
	key string
}

func (m *MSG) SetMSG(msg string) {
	m.Msg = msg
}

func (m *MSG) SetKey(key string) {
	m.key = key
}

func (m *MSG) Encrypted() {
	m.Msg = rot.Rot(m.Msg, 7)
	m.Msg = vigenere.Encipher(m.Msg, m.key)
	m.Msg = rot.Rot(m.Msg, 12)
}
func (m *MSG) Decrypted() {
	m.Msg = rot.Rot(m.Msg, -12)
	m.Msg = vigenere.Decipher(m.Msg, m.key)
	m.Msg = rot.Rot(m.Msg, -7)
}
