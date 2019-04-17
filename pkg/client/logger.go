/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains the definition of the logger interface that is used by the client.

package client

// Logger is the interface that must be implemented by objects that are used for logging by the
// client. By default the client uses a logger based on the `glog` package, but that can be changed
// using the `Logger` method of the builder.
type Logger interface {
	// DebugEnabled returns true iff the debug level is enabled.
	DebugEnabled() bool

	// InfoEnabled returns true iff the information level is enabled.
	InfoEnabled() bool

	// WarnEnabled returns true iff the warning level is enabled.
	WarnEnabled() bool

	// ErrorEnabled returns true iff the error level is enabled.
	ErrorEnabled() bool

	// Debug sends to the log a debug message formatted using the fmt.Sprintf function and the
	// given format and arguments.
	Debug(format string, args ...interface{})

	// Info sends to the log an information message formatted using the fmt.Sprintf function and
	// the given format and arguments.
	Info(format string, args ...interface{})

	// Warn sends to the log a warning message formatted using the fmt.Sprintf function and the
	// given format and arguments.
	Warn(format string, args ...interface{})

	// Error sends to the log an error message formatted using the fmt.Sprintf function and the
	// given format and arguments.
	Error(format string, args ...interface{})
}
