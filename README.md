# wowzers
a minimal dynamic window manager written in Golang for X11

## why?
I was searching for the perfect window manager but because I couldnt
find one, I decided to write my own. Here are my complaints about every
window manager I tried and how they led me to making this one

* i3 is nice but its manual
* bspwm does its configuration/scripting in a very innovative way, but its manual
* dwm is a brilliant window manager, I like how it is dynamic and how it does tags,
* but the thing is tho its configured in C, which is not easy to understand for most users.
* awesome is bloated, why should I have a dmenu, bar (okay this doesnt really annoy me
but its bar is kinda bloated) notification daemon and A COMPOSITOR built in!? yes you can
just not use them, but they are still loaded into memory in some capacity with how lua works,
also there is a lua runtime taking up memory -yes lua runtime is pretty light, but something
like IPC in bspwm is so much lighter, dwm is the lightest tho-
* Xmonad is written in Haskell, which aside being written in a functional language, when they are designed
for just doing math and data science and not practical things like wms, Haskell needs you to bloat your system up
with very much not lightweight dev tools
* Qtile is written in python.

I want to strike a happy balance between minimalism, configurability and actually being a dynamic window manager

* Its still configurable with IPC, but you are encouraged to modify the build in a suckless fashion, (but with Go,
instead of C so you dont go insane) in order to add more advanced features. IPC is not that flexible, only allowing you to do basic configuration, theming, event listening and custom layouts. Im even thinking of just making
the IPC a patch for the suckless purists

> in the future i also want it to work with multi monitor setups
but first i have to finish a single monitor setup, although i would like to make
it work like dwm but with XRandR

### inspiration
* [BurntSushi/wingo](https://github.com/BurntSushi/wingo)
* [suckless dwm](https://dwm.suckless.org)
* [baskerville/bspwm](https://github.com/baskerville/bspwm)

## WHAT ABOUT MUH WAYLAND X11 IZ INSEKUR AND SLOW AND WEENDO MANAGRS R OVURHED
![wayland user](https://i.imgflip.com/5mwo5q.png)

_Exhibit A: average wayland user_

first off all you might want to consider first switching to a better init system than soystemd and then
you can lecture me about how much better wayland is 

wlroots bindings for Go absolutley suck and are incomplete, i could use cgo and just interact with wlroots
that way, but i could just write it in C then? why i am insisting that i write in go? because i want people to
dive in on a structured AND DOCUMENTED codebase and make changes without much work.
