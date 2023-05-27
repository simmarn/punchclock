BINARY_NAME=punchclock
NEXTCLOUD=/run/user/1000/gvfs/dav\:host\=simmarn.mydns.se\,ssl\=true\,user\=martin\,prefix\=%2Fnextcloud%2Fremote.php%2Fwebdav/myApps/punchclock/


linux: 
	~/go/bin/fyne package -os linux -icon Icon.png

windows: 
	~/go/bin/fyne-cross windows -arch=amd64

export: 
	cp fyne-cross/bin/windows-amd64/Punchclock.exe $(NEXTCLOUD)
	ls -lh $(NEXTCLOUD)
