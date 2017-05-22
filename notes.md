
## Tasks
- [X] Environment variables
  + what are they?
  + unclear ...
- [x] sdk
- [x] apibuild + cliversion
- [x] props variable, revisited
- [x] Edge ?
-

---

+ parse toml file into properties and constants
+

## Notes

Now testing...  first, read through the tests and make sure that they look like they should pass...
- [ ] make a list of potentially failing tests -->
  + [x] activation poll
    + add method
      + in python app, `poll` wraps `console`
        + console
          + if no `since-`flags have been passed, then return fetch most recent (activation.list({limit:1,}))
          + else, once a second do:
            +

    + add flags
      + sinceSeconds [int ?]
      + sinceMinutes
      + sinceHours
      + sinceDays
      + exit int

  + [ ] namespace list --> type ?
  + [ ] add package refresh command / client action
    + `refresh` command
      + just takes an optional namespace argument
    + service
      +
  + [x] action create  --> add flags [ and maybe add to client as well ]
    + [ ] memory
    + [ ] sequence
    + [ ] param
    + [ ] annotation
    + [ ] timeout

- [ ] fix or approve tests in list

---
- [x] fails gracelessly when given invalid url as apiHOST

- [x] edge broken
  - needs protocol / https


- [x] set property broken

- [x] (unable to assign apihost with flag)


- [x] apibuild does not print out
--> client.Info.Get()
    --> info is blank --> print out response body

---



[x] Props -> namespace --> need to make a request to "/namespaces" first to get a list of legal namespaces, then confirm that requested is legal.


---

Need to consider how props is being stored ... --> need to have a single global props, with defaults


top-level properties struct

on init, load properties from .wskprops, environment, flags.

initialize client config from properties.



getProperties --> print out Properties --> according to flags set.

setProperties -->

  readProps

  according to flags, update props
  write props


unsetProperties -->

  readProps
  delete relevant ones.
  writeProps


readProps -> map[string]string, ok
writeProps(map[string]string) -> ok


---


Check that it is up to date...
Anything known to be missing?

SDK --> simple enough...  just do it.

What's the deal with apibuild, and cliversion ?

Do a side-by-side comparision of wsk versions.  should be the same, except formatting

Anything else I'm missing?

Environment variables.
WHISK_APIVERSION
WHISK_AUTH
WHISK_etc..
WHISK_




---

Updates to the client / command line api


client changes:
[X] add namespace get / list --> modify to be current (list is list triggers etc for current, get is list of namespaces)

Get namespace contents:

wsk namespace get --> GET /v1/namespaces/_
wsk namespace get wilsonth@us.ibm.com --> GET /v1/namespaces/_/wilsonth--

Get list of namespaces available

wsk namespace list -> GET /v1/namespaces

So ... --> update namespace client.


new top-level flags:
- `--apihost`: whisk api host
- `--apiversion`: whisk api version

[X] add `PropertyCmd`


Need to remove persistent global flags and re-add. (auth etc.)

- use: "work with whisk properties"
- set
  + -u, --auth
  + --apihost
  + --apiversion
  + --namespace


- get
  + -u, --auth
  + --apihost
  + --apiversion
  + --cliversion
  + --apibuild
  + --namespace
  + --all

Namespace --> api is `:443/api/{apivesion}/namespaces/{namespace}`

[x] remove top-level commands:
- auth
- list
- whoami
- health
- clean
- namespace
- version

remove top-level flags:

- --auth (added to local level ?)
-



Implemenmt sdk command
- install
  - component {docker, swift, iOS}

---

Parsing params, annotations, and action#invoke payload --> as json data

params and annotations --> attempt to parse as json into map[string]interface{}.  if it fails, then throw error

payload is the same except for that if it is not valid json then obj is created "{payload: arg}".

What about response object ??  Will also be a map[string]interface{} ?? 

To start: --> change action invoke :payload to a map[string]interface{} and see if it breaks.

Ok mostly working ...

--> need to add it back in to
- [x] trigger
- [x] package
- [x] rule ?


What does trigger#fire return?  {id: "id"}



---
Order of variables: flags -> env -> .wsk

1. load flags, env, and .wsk (props)
2. for each value, check and assign in order
3. initialize client + other variables that depend on values

config:
- `auth`
- `namespace`
- `edge`

- [X] load props and env. variables in main init.  Write a top-level persistent pre-run function to read the command line variables.

- [X] add setter functions
  + [X] `auth`
  + [X] `namespace`

---

package.bind ...

```python

def bind(self, args, props):
        url = 'https://%(url)s/api/v1/%(namespace)s/packages/%(name)s' % {
            'url': props['api'],
            'namespace': urllib.quote(args.namespace),
            'name': self.getSafeName(args.name)
        }
        split = args.package.split(':')
        binding = {}
        if (len(split) == 1):
            binding = { 'name': split[0], 'namespace': args.namespace}
        elif (len(split) == 2):
            binding = { 'name': split[1], 'namespace': split[0]}
        else:
            print 'package name malformed. name or namespace/name allowed'
            sys.exit(1)

        payload = {
            'name': args.name,
            'binding': binding,
            'annotations': getAnnotations(args),
            'parameters': getParams(args)
        }
        args.shared = False
        self.addPublish(payload, args)
        headers= {
            'Content-Type': 'application/json'
        }
        res = request('PUT', url, json.dumps(payload), headers, auth=args.auth, verbose=args.verbose)

        resBody = res.read()
        result = json.loads(resBody)

        if res.status == httplib.OK:
            print 'ok: created binding %(name)s ' % {'name': args.name }
            return 0
        else:
            print 'error: ' + result['error']
            return res.status
```

---


```python

def create(self, args, props, update):
        exe = self.getExec(args, props)
        if args.pipe:
            if args.param is None:
                args.param = []
            args.param.append([ '_actions', json.dumps(self.csvToList(args.artifact))])

        validExe = exe is not None and ('image' in exe or 'code' in exe)
        if update or validExe: # if create action, then exe must be valid
            payload = {
               'name': args.name,
               'annotations': getAnnotations(args),
               'parameters': getParams(args),
               'limits' : self.getLimits(args)
            }
            if validExe:
                payload['exec'] = exe
            self.addPublish(payload, args)
            return self.put(args, props, update, json.dumps(payload))
        else:
            print 'the artifact "%s" is not a valid file. If this is a docker image, use --docker.' % args.artifact
            return 2


# creates { code: "js code", image: "docker image", initializer: "base64 encoded string" }
# where code and image are mutually exclusive and initializer is optional
def getExec(self, args, props):
    exe = {}
    if args.docker:
        exe['image'] = args.artifact
    elif args.copy:
        existingAction = args.artifact
        exe = self.getActionExec(args, props, existingAction)
    elif args.pipe:
        args2 = copy.copy(args) # shallow copy of args object
        args2.namespace = 'client.system'
        pipeAction = 'common/pipe'
        exe = self.getActionExec(args2, props, pipeAction)
    elif args.artifact is not None and os.path.isfile(args.artifact):
        exe['code'] = open(args.artifact, 'rb').read()
    if args.lib:
        exe['initializer'] = base64.b64encode(args.lib.read())
    return exe

def getActionExec(self, args, props, name):
    res = self.Get(args, props, name)
    resBody = res.read()
    if res.status == httplib.OK:
        execField = json.loads(resBody)['exec']
    else:
        execField = None
    return execField

```


Action Create:

```golang

  if flags.docker
    exec.image = artifact // what artifact ?

  else if flags.copy
    -> actions.Get(actionName), copy exec

  else if flags.pipe
    -> (copy args)
    -> client.Config.Namespace = "client.system"
    -> actionName = "common/pipe"
    -> actions.Get(actionName), copy exec

  else if artifact != "" && os.FileExists(artifact)
    -> exec.code = os.ReadFile(artifact)

  if flags.lib
    -> exec.init = base64.Encode(flag.lib.read())  // lib is gzipped or tar file.




```


---

Thinking about how to persist data in between wsk calls.  The way that the python version does it is to write to a file on disk.  What other ways are there to do this?
- How does github cli do this ?

- Start working on command ...
  + fill out methods.
    + first need to create a reference to the whisk...  Top-level variable. --> parse flags, then assign

## To do's

- [X] actionInvokeCmd --> parse payload properly.

- [ ] better error responses
  + read resp.Body for message.

- [ ] Add support for environment variables
  - EDGE_HOST
  - CLI_API_HOST
  - WHISK_VERSION_DATE

- [x] create action
  + see above
- [ ] verbose
- [ ] Verbose (with a writer or something fancy)
  + how to avoid putting this on the client ?
- [X] params / annotations
- [X] positional arguments
- [ ] SDK


- [X] finish all simple methods
  + [x] package
  + [X] rule
  + [X] trigger

- [ ] Figure out how to properly define positional arguments with cobra / pflags.

- [ ] Implement verbose mode to help with debugging.

- [X] finish complex methods
  + [X] Action.Create with exec and all flags
  + [X] Action.Invoke with params

- Local install
  - vagrant.
  - test everything locally, include how in docs.
  - Debug locally.

- review how other cli packages store props (to disk)
  + hugo
  + github

- Cmd
  + implement loadConfig + updateConfig
  + add basic Client methods
    + auth
    + clean
    + version
  + add verbose
  + add arguments
  + add flags
    + top-level
    + sub-cmd-level

  + add messages
  + add functions (link up with stubbed out client + props)
- Client
  + [X] stub out methods for all services (with arguments)
  + [X] complete services
  + [X] complete request method for Client ()
  + [X] figure out what namespaces is about
  + [X] add auth to Client



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
  + Add Auth header in *whisk.Request
- verbose
  + include bool in Client struct
  + print out in *whisk.Do
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

        return
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


## Thoughts

---

Now... -> fill out command functions.

Perhaps start with a different function... with fewer flags ??


Need to figure out how flags will be done...

Can put in cmd.init() functions ??  yeah ...

would be really nice if i could test this ...


---

How to set namespace properly... ?

stored in .wskprops, initialized in whisk.
client offers namespaceService.List() only.

---

Review other python cli commands that are not listed in swagger doc (e.g. namespaces, sdk )

---
What does "clean" do ??

---

What am I doing with Config / props ??

What is the requirement?
> read .wsk config into map[string]string
> write map[string]string to file (configurable)


---


current issue:
Optional parameters... should not be listed in url params.  If not there, then don't print.
e.g. How to deal with activationsListOptions .since and .upto

possible solution: use pointers.  if pointer is nil, then ignore.


```go
func addRouteOptions(route string, options interface{}) (string, error) {
    v := reflect.ValueOf(options)
    if v.Kind() == reflect.Ptr && v.IsNil() {
        return route, nil
    }

    u, err := url.Parse(route)
    if err != nil {
        return route, err
    }

    qs, err := query.Values(options)
    if err != nil {
        return route, err
    }

    u.RawQuery = qs.Encode()
    return u.String(), nil
}
```
- How does go-querystring/query.Values() work ?? has options ??
  + want to take all non-nil values from options and write key=value in url.  Should ignore nil values (like for pointers and empty structs.)


Already does this!  using the tags... --> omit empty.  Anything to worry about then ... ?


- consider using pointers to structs.
-
- for now, just flag and skip anything difficult.
