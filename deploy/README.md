# ./deploy.sh

Deploy script.

## Usage

Simply run `./deploy.sh` without argument to see usage.

```
./deploy.sh [host] [db] [port]
```

## Example

To compile a go main package, switch to the directory of the package and do:

```
./deploy.sh freebsd-do test.db 80
```

where `freebsd-do` is a user@host entry in your `~/.ssh/config`


After a successful deploy, your server will be publicly accessible
at `<public-ip>:<port>`.
