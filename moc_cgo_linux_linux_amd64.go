//go:build !ubports
// +build !ubports

package main

/*
#cgo CFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -O2 -Wall -Wextra -D_REENTRANT -fPIC -DQT_NO_DEBUG -DQT_WIDGETS_LIB -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../shang -I. -I../Qt/5.15.2/gcc_64/include -I../Qt/5.15.2/gcc_64/include/QtWidgets -I../Qt/5.15.2/gcc_64/include/QtGui -I../Qt/5.15.2/gcc_64/include/QtCore -I. -I/usr/include/libdrm -I../Qt/5.15.2/gcc_64/mkspecs/linux-g++
#cgo LDFLAGS: -O1 -Wl,-rpath,/home/shang/Qt/5.15.2/gcc_64/lib
#cgo LDFLAGS:  /home/shang/Qt/5.15.2/gcc_64/lib/libQt5Widgets.so /home/shang/Qt/5.15.2/gcc_64/lib/libQt5Gui.so /home/shang/Qt/5.15.2/gcc_64/lib/libQt5Core.so -lGL -lpthread
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"
