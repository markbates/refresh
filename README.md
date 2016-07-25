# Refresh

## Rebuild and re-run your Go applications when files change.

This project was inspired by [https:#github.com/pilu/fresh](https:#github.com/pilu/fresh). Lack of updates and response from the maintainer, but a non-idiomatic codebase, numerous bugs, and lack of detailed reporting made the project a dead end for me to use. Enter `refresh`.

This simple command line application will watch your files, trigger a build of your Go binary and restart the application for you.

## Installation

```
$ go get github.com/markbates/fresh
```

## Getting Started

First you'll want to create a `refresh.yml` configuration file:

```
$ refresh init
```

If you want the config file in a different directory:

```
$ refresh init -c path/to/config.yml
```

Set it up the way you want, but I believe the defaults really speak for themselves, and will probably work for 90% of the use cases out there.

## Usage

Once you have your configuration all set up, all you need to do is run it:

```
$ refresh run
```

That's it! Now, as you change your code the binary will be re-built and re-started for you.

## HTTP Handler

Refresh is nice enough to ship with an `http.Handler` that you can wrap around your requests. Why would you want to do that?
Well, if there is an error doing a build, the built in `http.Handler` will print the error in your browser in giant text so you'll know that
there was a problem, and where to fix it (hopefully).

```go
...
m := http.NewServeMux()
err = http.ListenAndServe(":3000", web.ErrorChecker(m))
...
```

## Configuration Settings

```yml
# this is the root of your application relavite to your configuration file:
app_root: .
# a list of folders you don't want to watch. the folders you ignore, the faster things will be:
ignored_folders:
  - vendor
  - log
  - tmp
# a list of file extensions you want to watch for changes:
included_extensions:
  - .go
# the directory you want to build your binary in:
build_path: /tmp
# `fsnotify` can trigger many events at once when you change a file.
# in order to help cut down on the amount of builds that occur, a delay
# is used to let the extra events fly away.
build_delay: 200ms
# what would you like to call the built binary:
binary_name: refresh-build
# any extra commands you want to send to the built binary when it is run:
command_flags: []
# do you want to use colors when printing out log messages:
enable_colors: true
```
