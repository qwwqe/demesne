##
# Demesne diagrams
#
# @file
# @version 0.1

SRC=$(wildcard *.puml)
OBJ=$(SRC:.puml=.svg)

.PHONY: all
all:
	$(MAKE) --no-print-directory $(OBJ)

%.svg: %.puml
	plantuml -tsvg -p < $< | xsltproc enrichsvg.xsl - > $@

.PHONY: all
clean:
	rm $(OBJ)
# end
