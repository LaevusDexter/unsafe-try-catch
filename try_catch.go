package utc

import (
	// unasfe - for linking gosave and gogo
	_ "unsafe"
)

//go:linkname gosave runtime.gosave
//go:noescape
func gosave(buf *byte)

//go:linkname gogo runtime.gogo
//go:noescape
func gogo(buf *byte)

// TryCatcher - 
type TryCatcher struct {
	/* just because I don't want to copy whole struct from runtime src */
	gobuf [80]byte
	err error
}

// Catch - must be called in same function as Try (above)
func (tc *TryCatcher) Catch() error {
	gosave(&tc.gobuf[0])
	if tc.err != nil {
		err := tc.err
		tc.err = nil
		
		return err
	}

	return nil
}

// Try - must be called in same function as Catch (below)
func (tc *TryCatcher) Try(err error) {
	if err != nil {
		tc.err = err

		gogo(&tc.gobuf[0])
	}
}
