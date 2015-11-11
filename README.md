# syslogger.go

A simple utility to use in place of the `logger` GNU utility. It reads from
STDIN and outputs each line to the local syslog socket. Once it receives an EOF
it will quit. You can use the `-tag` flag to specify a syslog "tag" (or
"applicaiton"). The default is "go".

## Usage

```
echo "some text" | syslogger -tag sometag
```
## TODO

* Better error handling
* Command line arguments for priority (facility, severity). Currently
  everything goes to `local4.notice`.
