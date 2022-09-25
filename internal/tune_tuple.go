package internal


/**
 * 声调组合。
 */
type TuneTuple3 struct {
	First  Tune
	Second Tune
	Third  Tune
}

func NewTuneTuple3(first Tune, second Tune, third Tune) *TuneTuple3 {
	return &TuneTuple3{
		First:  first,
		Second: second,
		Third:  third,
	}
}

func NewTuneTuple3FromNum(first int, second int, third int) (*TuneTuple3, error) {
	firstTune, err := Int2Tune(first)
	if err != nil {
		return nil, err
	}
	secondTune, err := Int2Tune(second)
	if err != nil {
		return nil, err
	}
	thirdTune, err := Int2Tune(third)
	if err != nil {
		return nil, err
	}
	return &TuneTuple3{
		First:  firstTune,
		Second: secondTune,
		Third: thirdTune,
	}, nil
}

var PreferredTuneTupleSet = NewHashSet()

func init() {
	// 平平仄
	AddPreferredTuneTuple3(1, 1, 3)
	AddPreferredTuneTuple3(1, 1, 4)
	AddPreferredTuneTuple3(1, 2, 3)
	AddPreferredTuneTuple3(1, 2, 4)
	AddPreferredTuneTuple3(2, 1, 3)
	AddPreferredTuneTuple3(2, 1, 4)
	AddPreferredTuneTuple3(2, 2, 3)
	AddPreferredTuneTuple3(2, 2, 4)

	// 平仄仄
	AddPreferredTuneTuple3(1, 3, 3)
	AddPreferredTuneTuple3(1, 3, 4)
	AddPreferredTuneTuple3(1, 4, 3)
	AddPreferredTuneTuple3(1, 4, 4)
	AddPreferredTuneTuple3(2, 3, 3)
	AddPreferredTuneTuple3(2, 3, 4)
	AddPreferredTuneTuple3(2, 4, 3)
	AddPreferredTuneTuple3(2, 4, 4)

	// 平仄平
	AddPreferredTuneTuple3(1, 3, 1)
	AddPreferredTuneTuple3(1, 3, 2)
	AddPreferredTuneTuple3(1, 4, 1)
	AddPreferredTuneTuple3(1, 4, 2)
	AddPreferredTuneTuple3(2, 3, 1)
	AddPreferredTuneTuple3(2, 3, 2)
	AddPreferredTuneTuple3(2, 4, 1)
	AddPreferredTuneTuple3(2, 4, 2)

	// 仄仄平
	AddPreferredTuneTuple3(3, 3, 1)
	AddPreferredTuneTuple3(3, 3, 2)
	AddPreferredTuneTuple3(3, 4, 1)
	AddPreferredTuneTuple3(3, 4, 2)
	AddPreferredTuneTuple3(4, 3, 1)
	AddPreferredTuneTuple3(4, 3, 2)
	AddPreferredTuneTuple3(4, 4, 1)
	AddPreferredTuneTuple3(4, 4, 2)
}

func AddPreferredTuneTuple3(first int, second int, third int) {
	tuple, _ := NewTuneTuple3FromNum(first, second, third)
	PreferredTuneTupleSet.Put(*tuple)
}
