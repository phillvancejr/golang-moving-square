default: desktop

desktop: main.go minifb.go
	go build

.PHONY: web
web: main.go web_wasm.go
	cp web_wasm.go main.go web && \
	tinygo build -o web/move_square.wasm -target wasm --no-debug ./web

server: server_wasm.go
	go build -o $@ $<

platform=
ifndef platform
	platform = mac
endif

cmake_flags=
ifndef cmake_flags
	cmake_flags = -DUSE_METAL_API=ON 
endif

minifb:
	if [ ! -f deps/$(platform)/*minifb* ]; then \
	if [ ! -d deps/$(platform) ]; then mkdir deps/$(platform); fi;\
	mkdir deps/minifb/build;\
	cd deps/minifb/build;\
	cmake .. $(cmake_flags) -DMINIFB_BUILD_EXAMPLES=OFF;\
	cmake --build . ;\
	cp *minifb.* ../../$(platform);\
	fi

.SILENT: clean
clean:
	if [ -f move_square ]; then rm move_square; fi && \
	if [ -f web/move_square.wasm ]; then rm web/move_square.wasm; fi && \
	if [ -f web/main.go ]; then rm web/main.go; fi && \
	if [ -f web/web_wasm.go ]; then rm web/web_wasm.go; fi && \
	if [ -d deps/$(platform) ]; then rm -rf deps/$(platform); fi && \
	if [ -d deps/minifb/build ]; then rm -rf deps/minifb/build; fi && \
	if [ -f server ]; then rm server; fi
