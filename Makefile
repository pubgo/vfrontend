
build:
	gopherjs build github.com/pubgo/vfrontend -o static/frontend.js -m

s:
	gopherjs serve github.com/pubgo/vfrontend -m -v --http=:3000

jsgo:
	echo https://compile.jsgo.io/pubgo/vfrontend