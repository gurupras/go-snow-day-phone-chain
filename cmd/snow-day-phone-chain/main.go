package main

import (
	"flag"

	snowdayphonechain "github.com/gurupras/go-snow-day-phone-chain"
	log "github.com/sirupsen/logrus"
)

var minutes = flag.Uint("minutes", 0, "Number of minutes to run the phone chain for")
var callsPerPerson = flag.Uint("calls", 2, "Number of calls that each person makes")

func main() {
	flag.Parse()
	log.Debugf("Running phone-chain for %v minutes", *minutes)

	numPhoneCalls := snowdayphonechain.CalculateNumPhoneCalls(uint32(*minutes), uint32(*callsPerPerson))
	log.Infof("Total calls made in %vminutes: %v", *minutes, numPhoneCalls)
}
