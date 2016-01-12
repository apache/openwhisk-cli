## Thoughts


## Notes

Thinking about how to persist data in between wsk calls.  The way that the python version does it is to write to a file on disk.  What other ways are there to do this?
- How does github cli do this ?




## To do's

- review how other cli packages store props (to disk)
  + hugo
  + github



- Cmd
  + implement loadConfig + updateConfig
  + add basic Client methods
    + auth
    + clean
    + version
  + add arguments
  + add flags
    + top-level
    + sub-cmd-level

  + add messages
  + add functions (link up with stubbed out client)
- Client
  + stub out methods for all services (with arguments)
  + complete request method for Client ... include route specific stuff
  + add auth to Client
  + add verbose to Client



## Solutions




<!--

NONE OF THIS MATTERS!  stuff is reloaded every time the
 - watching props
  + have a loadPropsFromFile function
    + if file missing, use default
  + have an update prop(s) function
    + updates the file.
  + have a watch propsFile function -> updates client when it detects a change. -->

- auth
  + include token in Client struct (base64 encoded?)
  + Add Auth header in *Client.Request
- verbose
  + include bool in Client struct
  + print out in *Client.Do
- BUT ALSO!!!
  + need to store on disk so that it is the same in between invocations.  This is done in cmd --> initialized the client based on contents of .wskprops

## Code Samples From whisk *python

---

`wskitem.py`

```python
def put(self, args, props, update, payload):
  url = 'https://%(url)s/%(service)s/v1/%(namespace)s/%(collection)s/%(name)s%(update)s' % {
      'url': props[self.service],
      'service': self.service,
      'namespace': urllib.quote(args.namespace),
      'collection': self.collection,
      'name': self.getSafeName(args.name),
      'update': '?overwrite=true' if update else ''
  }

  headers= {
      'Content-Type': 'application/json'
  }

  res = request('PUT', url, payload, headers, auth=args.auth, verbose=args.verbose)
  resBody = res.read()

  if res.status == httplib.OK:
      print 'ok: %(mode)s %(item)s %(name)s' % { 'mode': 'updated' if update else 'created', 'item': self.name, 'name': args.name }
      return 0
  else:
      result = json.loads(resBody)
      print 'error: ' + result['error']
      return res.status
```

- This gives the url structure for resource requests
- `request` is defined in `wskutil.py`

---

`wskutil.py`

```python
def request(method, urlString, body = "", headers = {}, auth = None, verbose = False):
    url = urlparse(urlString)
    if url.scheme == 'http':
        conn = httplib.HTTPConnection(url.netloc)
    else:
        if hasattr(ssl, '_create_unverified_context'):
            conn = httplib.HTTPSConnection(url.netloc, context=ssl._create_unverified_context())
        else:
            conn = httplib.HTTPSConnection(url.netloc)

    if auth != None:
        auth = base64.encodestring(auth).replace('\n', '')
        headers['Authorization'] = 'Basic %s' % auth

    if verbose:
        print "========"
        print "REQUEST:"
        print "%s %s" % (method, urlString)
        print "Headers sent:"
        print json.dumps(headers, indent=4)
        if body != "":
            print "Body sent:"
            print body

    conn.request(method, urlString, body, headers)
    res = conn.getresponse()
    body = res.read()

    # patch the read to return just the body since the normal read
    # can only be done once
    res.read = lambda: body

    if verbose:
        print "--------"
        print "RESPONSE:"
        print "Got response with code %s" % res.status
        print "Body received:"
        print res.read()
        print "========"

    return res
```

- Shows auth scheme --> just base64 encode and use basic authorization for requests


## Code samples from hugo

---

from `commands/hugo.go`

```go

// Execute adds all child commands to the root command HugoCmd and sets flags appropriately.
func Execute() {
	HugoCmd.SetGlobalNormalizationFunc(helpers.NormalizeHugoFlags)

	HugoCmd.SilenceUsage = true

	AddCommands()

	if c, err := HugoCmd.ExecuteC(); err != nil {
		if isUserError(err) {
			c.Println("")
			c.Println(c.UsageString())
		}

		os.Exit(-1)
	}
}

// AddCommands adds child commands to the root command HugoCmd.
func AddCommands() {
	HugoCmd.AddCommand(serverCmd)
	HugoCmd.AddCommand(versionCmd)
	HugoCmd.AddCommand(configCmd)
	HugoCmd.AddCommand(checkCmd)
	HugoCmd.AddCommand(benchmarkCmd)
	HugoCmd.AddCommand(convertCmd)
	HugoCmd.AddCommand(newCmd)
	HugoCmd.AddCommand(listCmd)
	HugoCmd.AddCommand(undraftCmd)
	HugoCmd.AddCommand(importCmd)

	HugoCmd.AddCommand(genCmd)
	genCmd.AddCommand(genautocompleteCmd)
	genCmd.AddCommand(gendocCmd)
	genCmd.AddCommand(genmanCmd)
}

```

- use a top-level execute function like this to wrap Cmd.Execute().
- Add all the commands in one place in the main function --> makes it much easier to see what's going on vs. in lots of different init files --> also easier to update / edit etc.

```go




// Flags that are to be added to commands.
var BuildWatch, IgnoreCache, Draft, Future, UglyURLs, CanonifyURLs, Verbose, Logging, VerboseLog, DisableRSS, DisableSitemap, DisableRobotsTXT, PluralizeListTitles, PreserveTaxonomyNames, NoTimes, ForceSync bool
var Source, CacheDir, Destination, Theme, BaseURL, CfgFile, LogFile, Editor string

func initCoreCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&Draft, "buildDrafts", "D", false, "include content marked as draft")
	cmd.Flags().BoolVarP(&Future, "buildFuture", "F", false, "include content with publishdate in the future")
	cmd.Flags().BoolVar(&DisableRSS, "disableRSS", false, "Do not build RSS files")
	cmd.Flags().BoolVar(&DisableSitemap, "disableSitemap", false, "Do not build Sitemap file")
	cmd.Flags().BoolVar(&DisableRobotsTXT, "disableRobotsTXT", false, "Do not build Robots TXT file")
	cmd.Flags().StringVarP(&Source, "source", "s", "", "filesystem path to read files relative from")
	cmd.Flags().StringVarP(&CacheDir, "cacheDir", "", "", "filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/")
	cmd.Flags().BoolVarP(&IgnoreCache, "ignoreCache", "", false, "Ignores the cache directory for reading but still writes to it")
	cmd.Flags().StringVarP(&Destination, "destination", "d", "", "filesystem path to write files to")
	cmd.Flags().StringVarP(&Theme, "theme", "t", "", "theme to use (located in /themes/THEMENAME/)")
	cmd.Flags().BoolVar(&UglyURLs, "uglyURLs", false, "if true, use /filename.html instead of /filename/")
	cmd.Flags().BoolVar(&CanonifyURLs, "canonifyURLs", false, "if true, all relative URLs will be canonicalized using baseURL")
	cmd.Flags().StringVarP(&BaseURL, "baseURL", "b", "", "hostname (and path) to the root, e.g. http://spf13.com/")
	cmd.Flags().StringVar(&CfgFile, "config", "", "config file (default is path/config.yaml|json|toml)")
	cmd.Flags().StringVar(&Editor, "editor", "", "edit new content with this editor, if provided")
	cmd.Flags().BoolVar(&nitro.AnalysisOn, "stepAnalysis", false, "display memory and timing of different steps of the program")
	cmd.Flags().BoolVar(&PluralizeListTitles, "pluralizeListTitles", true, "Pluralize titles in lists using inflect")
	cmd.Flags().BoolVar(&PreserveTaxonomyNames, "preserveTaxonomyNames", false, `Preserve taxonomy names as written ("Gérard Depardieu" vs "gerard-depardieu")`)
	cmd.Flags().BoolVarP(&ForceSync, "forceSyncStatic", "", false, "Copy all files when static is changed.")
	// For bash-completion
	validConfigFilenames := []string{"json", "js", "yaml", "yml", "toml", "tml"}
	cmd.Flags().SetAnnotation("config", cobra.BashCompFilenameExt, validConfigFilenames)
	cmd.Flags().SetAnnotation("source", cobra.BashCompSubdirsInDir, []string{})
	cmd.Flags().SetAnnotation("cacheDir", cobra.BashCompSubdirsInDir, []string{})
	cmd.Flags().SetAnnotation("destination", cobra.BashCompSubdirsInDir, []string{})
	cmd.Flags().SetAnnotation("theme", cobra.BashCompSubdirsInDir, []string{"themes"})
}
```

- Parse variables like this.

---
