package commands

///////////
// Flags //
///////////

var flags struct {
	global struct {
		verbose    bool
		edge       bool
		auth       string
		namespace  string
		apihost    string
		apiversion string
	}

	common struct {
		blocking   bool
		annotation string
		param      string
		shared     bool // AKA "public" or "publish"
		skip       int  // skip first N records
		limit      int  // return max N records
		full       bool // return full records (docs=true for client request)
	}

	// // shared persistent
	// verbose    bool
	// edge       bool
	// auth       string
	// namespace  string
	// apihost    string
	// apiversion string
	//
	// // shared
	// blocking   bool
	// annotation string
	// param      string
	//
	// // shared list
	// skip  int  // skip first N records
	// limit int  // return max N records
	// full  bool // return full records (docs=true for client request)

	// TODO :: consider using a sub-struct to namespace these.. like property struct { auth bool } etc.  then &flags.property.auth
	property struct {
		auth       bool
		apihost    bool
		apiversion bool
		namespace  bool
		cliversion bool
		apibuild   bool
		all        bool
	}

	// property
	// bAuth       bool
	// bApihost    bool
	// bApiversion bool
	// bNamespace  bool
	// cliversion  bool
	// apibuild    bool
	// all         bool

	action struct {
		docker   bool
		copy     bool
		pipe     bool
		shared   bool
		lib      string
		timeout  int
		memory   int
		xPackage string
	}

	// // action
	// docker   bool
	// copy     bool
	// pipe     bool
	// shared   bool
	// lib      string
	// timeout  int
	// memory   int
	// xPackage string

	activation struct {
		action  string // retrieve results for this action
		upto    int    // retrieve results up to certain time
		since   int    // retrieve results after certain time
		seconds int    // stop polling for activation upda
	}

	// // activation
	// action  string // retrieve results for this action
	// upto    int    // retrieve results up to certain time
	// since   int    // retrieve results after certain time
	// seconds int    // stop polling for activation updates after certain time

	xPackage struct {
		serviceGUID string
	}

	// // package
	// serviceGUID string

	// rule
	rule struct {
		auto bool
	}

	// auto bool
}
