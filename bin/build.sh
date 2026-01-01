pushd ../src/$1 
tinygo build -tags debug -target=pico -o ../../bin/$1.uf2
popd
