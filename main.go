package main

import (
	"os"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("cspsolver")

func main() {

	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.7s} %{color:reset} %{message}`,
	)

	logBackend := logging.AddModuleLevel(logging.NewLogBackend(os.Stdout, "", 0))
	logBackend.SetLevel(logging.NOTICE, "")

	logging.SetBackend(logging.NewBackendFormatter(logBackend, format))

	// simple implementation of the map coloring problem for Australia
	colors := [3]string{"red", "green", "blue"}

	// set a variable for each of the provinces
	vars := IntVariables{
		// each has:   <name>,      <domain>
		NewIntVariable("WA", IntRange(1, len(colors)+1)),
		NewIntVariable("NT", IntRange(1, len(colors)+1)),
		NewIntVariable("Q", IntRange(1, len(colors)+1)),
		NewIntVariable("NSW", IntRange(1, len(colors)+1)),
		NewIntVariable("V", IntRange(1, len(colors)+1)),
		NewIntVariable("SA", IntRange(1, len(colors)+1)),
		NewIntVariable("T", IntRange(1, len(colors)+1)),
	}

	// bordering provinces cannot be equal
	constraints := IntConstraints{
		NotEqualsInt("WA", "NT"),
		NotEqualsInt("WA", "SA"),
		NotEqualsInt("NT", "SA"),
		NotEqualsInt("NT", "Q"),
		NotEqualsInt("Q", "SA"),
		NotEqualsInt("Q", "NSW"),
		NotEqualsInt("NSW", "V"),
		NotEqualsInt("NSW", "SA"),
		NotEqualsInt("V", "SA"),
	}

	solver := NewIntCSPSolver(vars, constraints, 500)
	begin := time.Now()
	success := solver.Solve()
	elapsed := time.Since(begin)

	if success {
		// log.Noticef("Found solution in %s, and variables are: %v\n", success, elapsed, solver.State.Vars)
		log.Noticef("Found solution in %s\n", elapsed)

		for _, variable := range solver.State.Vars {
			log.Noticef("Variable %v = %v", variable.Name, colors[variable.Value-1])
		}
	} else {
		log.Errorf("Could not find solution in %s\n", elapsed)
	}
}
