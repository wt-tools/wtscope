# Makefile include for generation of PlantUML docs.
# Project doc dir
DOCDIR:=doc
# Your local copy of plantuml, could be set as env var
# It requires Java installed!
PLANTJAR?=~/plantuml/plantuml.jar
PLANTURL = https://sourceforge.net/projects/plantuml/files/plantuml.jar/download
UMLSRC := $(wildcard $(DOCDIR)/*.plantuml)
PNGOUT := $(addsuffix .png, $(basename $(UMLSRC)))
SVGOUT := $(addsuffix .svg, $(basename $(UMLSRC)))
CLEANUP+= $(SVGOUT)

# Default target first; build PNGs, probably what we want most of the time
png: $(PLANTJAR) $(PNGOUT)

# SVG are nice-to-have but don't need to build by default
svg: $(PLANTJAR) $(SVGOUT)

# If the JAR file isn't already present, download it
$(PLANTJAR):
ifeq (,$(wildcard $(PLANTJAR)))
	mkdir -p ~/plantuml
	curl -sSfL -o ~/plantuml/plantuml.jar $(PLANTURL)
endif

doc/%.png: $(DOCDIR)/%.plantuml
	java -jar $(PLANTJAR) -nbthread auto -tpng $^

doc/%.svg: $(DOCDIR)/%.plantuml
	java -jar $(PLANTJAR) -nbthread auto -tsvg $^

doc: png svg

.PHONY: doc png svg
