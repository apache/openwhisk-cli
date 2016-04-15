package commands

///////////
// Flags //
///////////

var flags struct {
	global struct {
		verbose    bool
		auth       string
		namespace  string
		apihost    string
		apiversion string
	}

	common struct {
		blocking bool
		// annotation string
		annotation []string
		// param      string
		param   []string
		shared  bool // AKA "public" or "publish"
		skip    int  // skip first N records
		limit   int  // return max N records
		full    bool // return full records (docs=true for client request)
		summary bool
	}

	property struct {
		auth       bool
		apihost    bool
		apiversion bool
		namespace  bool
		cliversion bool
		apibuild   bool
		all        bool
	}

	action struct {
		docker   bool
		copy     bool
		pipe     bool
		shared   bool
		sequence bool
		lib      string
		timeout  int
		memory   int
		result   bool
		xPackage string
	}

	activation struct {
		action       string // retrieve results for this action
		upto         int64  // retrieve results up to certain time
		since        int64  // retrieve results after certain time
		seconds      int    // stop polling for activation upda
		sinceSeconds int
		sinceMinutes int
		sinceHours   int
		sinceDays    int
		exit         int
	}

	xPackage struct {
		serviceGUID string
	}

	// rule
	rule struct {
		enable  bool
		disable bool
	}
}
