package commands

///////////
// Flags //
///////////

var flags struct {
	// shared persistent
	verbose   bool
	edge      bool
	auth      string
	namespace string

	// shared
	blocking   bool
	annotation []string
	param      []string

	// shared list
	skip  int  // skip first N records
	limit int  // return max N records
	full  bool // return full records (docs=true for client request)

	// auth
	xType string

	// action
	docker  bool
	copy    bool
	pipe    bool
	shared  bool
	lib     string
	timeout int
	memory  int

	// activation
	upto    int // retrieve results up to certain time
	since   int // retrieve results after certain time
	seconds int // stop polling for activation updates after certain time

	// namespace

	// package
	serviceGuid string

	// rule
	auto bool
}
