## Thoughts

What else now ???
--> 


## Notes

Thinking about how to persist data in between wsk calls.  The way that the python version does it is to write to a file on disk.  What other ways are there to do this?
- Can keep it all in memory, in the client.  Can have a single point of entry `wsk` and then have an interactive environment based on that...  Actually I think the `.wskprops` approach is standard... git writes to a config file in a similar way... look at the github cli (or hugo / something written with cobra ) for inspiration.

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
