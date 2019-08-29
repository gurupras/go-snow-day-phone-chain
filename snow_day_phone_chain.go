package snowdayphonechain

import "sync"

type person struct {
	callsRemaining uint32
}

func newPerson(callsPerPerson uint32) *person {
	return &person{callsPerPerson}
}

type tickEvent int

const tickActionCall tickEvent = 1
const tickActionDone tickEvent = 2

type tickAction struct {
	action tickEvent
	person *person
}

// For each person in the list, reduce the number of calls available by 1
// If the person is done making callsPerPerson calls, then remove them from the persons list
// We do this by using different events
func (p *person) tick(eventChan chan *tickAction) {
	var action *tickAction
	if p.callsRemaining > 0 {
		action = &tickAction{
			tickActionCall,
			p,
		}
		p.callsRemaining--
	} else {
		action = &tickAction{
			tickActionDone,
			p,
		}
	}
	eventChan <- action
}

// CalculateNumPhoneCalls calculates the number of phone calls made within
// numMinutes time at callsPerPerson number of calls per person
func CalculateNumPhoneCalls(numMinutes uint32, callsPerPerson uint32) uint64 {
	// Short-circuit in case values are 0
	if numMinutes == 0 || callsPerPerson == 0 {
		return 0
	}
	// Because we're treating the number of minutes as a variable
	// we're going to use an iterative solution rather than a recursive one to ensure
	// that we're not blowing up the stack
	// We start with one person

	persons := make([]*person, 0)
	p := newPerson(callsPerPerson)
	persons = append(persons, p)

	removePerson := func(p *person) {
		targetIdx := -1
		for idx, _p := range persons {
			if _p == p {
				targetIdx = idx
				break
			}
		}
		if targetIdx != -1 {
			persons = append(persons[:targetIdx], persons[targetIdx+1:]...)
		}
	}

	var numCalls uint64
	for i := uint32(0); i < numMinutes; i++ {
		// The number of calls made this iteration
		// Each person can only make *one* call a minute
		tickChan := make(chan *tickAction)
		wg := sync.WaitGroup{}
		wg.Add(1)

		// Calls made this iteration
		var callsMade uint64

		go func() {
			defer wg.Done()
			for event := range tickChan {
				switch event.action {
				case tickActionCall:
					callsMade++
				case tickActionDone:
					removePerson(event.person)
				}
			}
		}()

		// Iterate over all the available persons and either place a call or remove them from the list if they're done
		tmp := make([]*person, len(persons))
		copy(tmp, persons)
		for _, p := range tmp {
			p.tick(tickChan)
		}
		close(tickChan)
		wg.Wait()

		// For each call that was made, a new person joins the chain
		for idx := uint64(0); idx < callsMade; idx++ {
			p := newPerson(callsPerPerson)
			persons = append(persons, p)
		}
		// Sum up the total number of calls made overall
		numCalls += callsMade
	}
	return numCalls
}
