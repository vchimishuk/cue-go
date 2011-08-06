include $(GOROOT)/src/Make.inc

TARG=cue

GOFILES=\
	cue.go\
	sheet.go\
	parser.go\
	utils.go\

include $(GOROOT)/src/Make.pkg

format:
	find . -type f -name '*.go' -exec gofmt -w {} \;

arch-install:
	mkdir -p "$(DESTDIR)$(pkgdir)"
	cp _obj/$(TARG).a "$(DESTDIR)$(pkgdir)"
