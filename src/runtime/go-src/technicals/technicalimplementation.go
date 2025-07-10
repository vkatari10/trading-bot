package technicals

// contains all implementations for all declared/supported
// technical indicators for the "Technical" interface in 
// technicaltypes.go

// Tag Methods
func (Delta) 	Tag() {}
func (Diff) 	Tag() {}

// Type Methods
func (Delta)	Type() (string) {return "delta"}
func (Diff) 	Type() (string) {return "diff"}

// All LastValue() Methods should be implemented in their own files