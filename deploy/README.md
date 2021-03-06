# ./deploy.sh

Deploy script.

## Usage

Simply run `./deploy.sh` without arguments to see usage.

```
./deploy.sh [host] [db] [port]
```

For the port argument to be used, the script expects
your main program accepts a `-http` flag for the HTTP address
(similar to godoc).

## Example

To deploy a main go package, switch to the directory of the package and do:

```
./deploy.sh freebsd-do test.db 80
```

where `freebsd-do` is a user@host entry in your `~/.ssh/config`


After a successful deploy, your server will be publicly accessible
at `<public-ip>:<port>`.
