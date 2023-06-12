package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	//panic("TODO: implement this!")
	if len(input) == 0 {
		return start == final
	}
	n := len(input)
	var state1 state = start
	var state2 state = 10
	var new_state1 state
	var new_state2 state
	for i := 0; i < n; i++ {
		new_state1 = 10
		new_state2 = 10
		if state1 == 10 {
			return false
		} else {
			var arr1 = transitions(state1, input[i])
			for j := range arr1 {
				if new_state1 == 10 {
					new_state1 = arr1[j]
				} else {
					new_state2 = arr1[j]
				}
			}
			if state2 != 10 {
				var arr2 = transitions(state2, input[i])
				for j := range arr2 {
					if new_state2 == 10 {
						new_state2 = arr2[j]
					}
				}
			}
		}
		state1 = new_state1
		state2 = new_state2
	}

	return state1 == final || state2 == final
}
