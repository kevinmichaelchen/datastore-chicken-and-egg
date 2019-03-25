## How to run

### In one tab, run
```bash
# Clear DB
rm ~/.config/gcloud/emulators/datastore/WEB-INF/appengine-generated/local_db.bin

# Start DB
gcloud beta emulators datastore start
```

### In another tab, run
```bash
# Set environment so we know what Project ID to use
eval "$(gcloud beta emulators datastore env-init)"

go run *.go
``` 

## What does this show
If we use `goodKey`, we see this printed:
```
GET RESULT = Folder[ID=2, ParentID=1]
```

If we use `badKey`, we get the error `datastore: no such entity`.

Basically, there's a chicken and egg problem.

Imagine a file system.
Say we're just dealing with folders.
Say we want to lookup a folder by its key.
Its "key" comprises both the folder ID
and the folder's parent ID...

But how do we know the folder's parent ID
before looking up the folder itself?