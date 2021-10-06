# Stuff you should know

If you found this and you wanna use it, good luck lol proceed at your own risk. Here are some tips:

You need firefox installed, you probably need to install Go and build the binary yourself and you might also need to install mingw. I wrote this on windows for my windows machine so it works for windows. If you want it to work on any other platform you will need to get the gecko driver binary for that platform and modify the code to use that instead. I might do this one day if the mood strikes me but I'm not sure how well the code will port to other systems since it depends on robotgo.

Since I used a hacky heuristic for the initial text entry (wait 5 seconds and then simulate keyboard strokes, relying on the browser to auto focus) it might fail sometimes (or always if your computer is too slow), just restart it if it does.

Here is what you do: 

```
go build
./proton-gen.exe --email=myRecoveryEmail@email.com
```

Once it finishes (takes about 15 seconds), grab the username and password from your console, complete the captcha and you're done.

# Dependencies

here are some helpful links in case you want to use this:

* [golang](https://golang.org/doc/install)
* [firefox](https://www.mozilla.org/en-US/firefox/new/)
* [mingw ONLY INSTALL THIS IF YOU BUILT FROM SOURCE AND IT STILL FAILED TO RUN](https://www.mingw-w64.org/)