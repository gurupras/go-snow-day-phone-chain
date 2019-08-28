package snowdayphonechain

// CalculateNumPhoneCalls calculates the number of phone calls made within
// numMinutes time at callsPerPerson number of calls per person
func CalculateNumPhoneCalls(numMinutes uint32, callsPerPerson uint32) uint64 {
	// Because we're treating the number of minutes as a variable
	// we're going to use an iterative solution rather than a recursive one to ensure
	// that we're not blowing up the stack
	// We start with one person

	var numPersons uint64 = 1
	var numCalls uint64
	for i := uint32(0); i < numMinutes; i++ {
		callsMade := (uint64(numPersons) * uint64(callsPerPerson)) // The number of calls made this iteration
		// In the next iteration, each of these people are going to make callsPerPerson calls
		numPersons = callsMade
		// Sum up the total number of calls made overall
		numCalls += callsMade
	}
	return numCalls
}
