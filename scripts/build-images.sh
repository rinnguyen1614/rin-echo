echo "================================"
echo "Build System Service..."
echo "================================"
./src/services/system/build-image.sh


echo "================================"
echo "Build web ..."
echo "================================"
./src/web/build-image.sh
