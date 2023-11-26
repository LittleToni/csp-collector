package report

// Writer is an interface defining a method to write a report. This interface
// can be implemented by different types that handle the writing of report data
// to various destinations such as a file, a database, or over a network.
//
// The purpose of defining this interface is to allow for flexibility in how reports
// are written, enabling the implementation of multiple writing strategies without
// altering the core logic of report processing.
type Writer interface {
	Write(report string) error
}
