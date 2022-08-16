INCLUDES=-I includes/
CXXFLAGS=-std=c++20 $(INCLUDES)
# CXXFLAGS=-std=c++20 
CXX=clang++

clean:
	rm -f bin/*

exec: src/main.cc 
	$(CXX) $(CXXFLAGS) $^ -o bin/exec && ./bin/exec > ./bin/image.ppm && cd ..

tests: tests/tests.cc src/board.cc
	$(CXX) $(CXXFLAGS) $^ -o bin/tests

src/%.cc: includes/%.hpp
	touch $@

.DEFAULT_GOAL := tests
.PHONY: tests clean exec