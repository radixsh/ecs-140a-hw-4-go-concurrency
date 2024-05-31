package nfa

// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, sym rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
    if len(input) == 0 {
        return start == final
    }

    state_ch := make(chan state, 10)
    input_ch := make(chan rune, 10)
   
    // Populate with the first stuff
    state_ch <- transitions(start, input[0])
    input_ch <- input[0]

    count := 0
    var accessible_states []state
    for next_input := range input_ch {
        for next_state := range state_ch {
            accessible_states := <- state_ch
            state_ch <- transitions(next_state, next_input)
            count++
            input_ch <- input[count]
        }
    }

	// return accessible_states.contains(final)
	for _, st := range accessible_states {
		if st == final {
			return true
		}
	}
    return false

    /*
    // Asynchronously get all the states reachable from the first state via the
    // first input rune. Save them to a channel.
    ch := make(chan []state)
    ch <- transitions(st, char)
    // Then, for each state in that array, keep following down asynchronously,
    // popping runes every time.
    for state := range ch {


    }
    // Once all the runes are done, see if the provided "final state" is an
    // element of the set.
    // locking it before writing to ensure no race conditions.
    // append them to a growing list

	old_states := []state{start}
    ch := make(chan []state, 1)
	for _, char := range input {
		new_states := []state{}
		for _, st := range old_states {
            // Send return value of transitions() fn call into channel
            ch <- transitions(st, char)
			new_states = append(new_states, <- ch)
		}
		old_states = new_states
	}

	for _, st := range old_states {
		if st == final {
			return true
		}
	}
	return false

    // divide here

	old_states := []state{start}
    ch := make(chan []state, 1)
	for _, char := range input {
		new_states := []state{}
		for _, st := range old_states {
            // Send return value of transitions() fn call into channel
            ch <- transitions(st, char)
			new_states = append(new_states, <- ch)
		}
		old_states = new_states
	}

	for _, st := range old_states {
		if st == final {
			return true
		}
	}
	return false
    */
}
