package std

/*
Logger defines a common logger interface for packages that accept an
injected logger.

This interface declares that, because golang.org/pkg/log is insufficient
(or at least cumbersome) for good visibility into high-throughput
software in load-balanced, clustered, and/or multi-stage environments, a
logger that is more friendly to those situations is required.

Packages that implement this interface should support level codes and
expect usage as follows:

	- Panic, highest level of severity. Log and then call panic.
	- Fatal, log and then call `os.Exit(<code>)` with a non-zero value.
	- Error, used for errors that should definitely be noted and addressed.
	- Warn, non-critical information about undesirable behavior that needs
	  to be addressed in some way.
	- Info, general operational information about what's happening inside an
	  application.
	- Debug, usually only enabled when debugging. Add anything useful, but
	  still exclude PII or sensitive data...

Each level should include all the log levels above it. Compatible logger
packages include:

	- "github.com/sirupsen/logrus"
*/
type Logger interface {

	// Panic methods should end with a call to the panic() builtin to
	// allow recovery.
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	// Fatal methods should call os.Exit() with a non-zero exit status.
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	// Level 2
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	// Level 3
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	// Level 4
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	// Level 5
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
}
