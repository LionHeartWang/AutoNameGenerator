package internal

// 拼音表
type Tone int

const (
	// 声母
	b  = 1
	p  = 2
	m  = 3
	f  = 4
	d  = 5
	t  = 6
	n  = 7
	l  = 8
	g  = 9
	k  = 10
	h  = 11
	j  = 12
	q  = 13
	x  = 14
	zh = 15
	ch = 16
	sh = 17
	r  = 18
	z  = 19
	c  = 20
	s  = 21
	y  = 22
	w  = 23

	// 单韵母
	a = 24
	o = 25
	e = 26
	i = 27
	u = 28
	v = 29

	// 复韵母
	ai = 30
	ei = 31
	ui = 32
	ao = 33
	ou = 34
	iu = 35
	ie = 36
	ve = 37
	er = 38

	// 前鼻韵母
	an = 39
	en = 40
	in = 41
	un = 42
	vn = 43

	// 后鼻韵母
	ang = 44
	eng = 45
	ing = 46
	ong = 47

	unknown = 48
)

func Str2Tone(str string) Tone {
	switch str {
	case "b":
		return b
	case "p":
		return p
	case "m":
		return m
	case "f":
		return f
	case "d":
		return d
	case "t":
		return t
	case "n":
		return n
	case "l":
		return l
	case "g":
		return g
	case "k":
		return k
	case "h":
		return h
	case "j":
		return j
	case "q":
		return q
	case "x":
		return x
	case "zh":
		return zh
	case "ch":
		return ch
	case "sh":
		return sh
	case "r":
		return r
	case "z":
		return z
	case "c":
		return c
	case "s":
		return s
	case "y":
		return y
	case "w":
		return w
	case "a":
		return a
	case "o":
		return o
	case "e":
		return e
	case "i":
		return i
	case "u":
		return u
	case "v":
		return v
	case "ai":
		return ai
	case "ei":
		return ei
	case "ui":
		return ui
	case "ao":
		return ao
	case "ou":
		return ou
	case "iu":
		return iu
	case "ie":
		return ie
	case "ve":
		return ve
	case "er":
		return er
	case "an":
		return an
	case "en":
		return en
	case "in":
		return in
	case "un":
		return un
	case "vn":
		return vn
	case "ang":
		return ang
	case "eng":
		return eng
	case "ing":
		return ing
	case "ong":
		return ong
	default:
		return unknown
	}
}

func IsInitialTone(tone Tone) bool {
	return tone >= b && tone <= w
}

func IsFinalTone(tone Tone) bool {
	return tone >= a && tone <= ong
}
