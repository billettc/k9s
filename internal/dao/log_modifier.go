package dao

import zappretty "github.com/maoueh/zap-pretty"

type LogModifier interface {
	Modify(line []byte) []byte
}

func init() {
	RegisterLogModifier("zap-pretty", NewZapPrettyLogModifier())
}

type ZapPrettyLogModifier struct {
}

func NewZapPrettyLogModifier() *ZapPrettyLogModifier {
	return &ZapPrettyLogModifier{}
}

func (m *ZapPrettyLogModifier) Modify(line []byte) []byte {
	l := string(line)
	pretty, err := zappretty.PrettyLine(l, true)
	if err != nil {
		return line
	}
	return []byte(pretty)
}
