# Stuff you should know

If you found this and you wanna use it, good luck lol proceed at your own risk. Here are some tips:

You need firefox installed, you probably need to install Go and build the binary yourself and you might also need to install mingw. I wrote this on windows for my windows machine so it works for windows. If you want it to work on any other platform you will need to get the gecko driver binary for that platform and modify the code to use that instead. I might do this one day if the mood strikes me but I'm not sure how well the code will port to other systems since it depends on robotgo.

Since I used a hacky heuristic for the initial text entry (wait 5 seconds and then simulate keyboard strokes, relying on the browser to auto focus) it might fail sometimes (or always if your computer is too slow), just restart it if it does.

Once it finishes (takes about 15 seconds), grab the username and password from your console, complete the captcha and you're done.

# Build Dependencies

here are some helpful links in case you want to use this, all of these things are required to build from source:

* [golang](https://golang.org/doc/install)
* [mingw](https://www.mingw-w64.org/downloads/)

# Runtime Dependencies

I have no idea if the provided exe will actually work on other systems, but you are free to try it

* [firefox](https://www.mozilla.org/en-US/firefox/new/)
* [java](https://java.com/en/download/manual.jsp)

# Usage

Download the provided binary or build from source

Here is what you do: 

```
./proton-gen.exe --email=myRecoveryEmail@email.com
```

A firefox window will open, just wait 15-20 seconds for it to move through the prompts and enter the randomly generated info. Then you must manually complete the captcha. Once that is complete, grab the username and password from the terminal window that you run proton-gen in and you can use those details to login to protonmail.