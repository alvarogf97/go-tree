package scan

// Scanner options struct
type ScannerOptions struct {
	showHidden bool
}

// Creates a new scanner option struct
func NewScannerOptions(showHidden bool) ScannerOptions {
	return ScannerOptions{showHidden: showHidden}
}
