package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type PlainTextEdit_ITF interface {
	std_widgets.QPlainTextEdit_ITF
	PlainTextEdit_PTR() *PlainTextEdit
}

func (ptr *PlainTextEdit) PlainTextEdit_PTR() *PlainTextEdit {
	return ptr
}

func (ptr *PlainTextEdit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlainTextEdit_PTR().Pointer()
	}
	return nil
}

func (ptr *PlainTextEdit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPlainTextEdit_PTR().SetPointer(p)
	}
}

func PointerFromPlainTextEdit(ptr PlainTextEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.PlainTextEdit_PTR().Pointer()
	}
	return nil
}

func NewPlainTextEditFromPointer(ptr unsafe.Pointer) (n *PlainTextEdit) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(PlainTextEdit)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *PlainTextEdit:
			n = deduced

		case *std_widgets.QPlainTextEdit:
			n = &PlainTextEdit{QPlainTextEdit: *deduced}

		default:
			n = new(PlainTextEdit)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *PlainTextEdit) Init() { this.init() }

//export callbackPlainTextEdit209179_Constructor
func callbackPlainTextEdit209179_Constructor(ptr unsafe.Pointer) {
	this := NewPlainTextEditFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackPlainTextEdit209179_AddLine
func callbackPlainTextEdit209179_AddLine(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "addLine"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(v0))
	}

}

func (ptr *PlainTextEdit) ConnectAddLine(f func(v0 string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "addLine"); signal != nil {
			f := func(v0 string) {
				(*(*func(string))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "addLine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "addLine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PlainTextEdit) DisconnectAddLine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "addLine")
	}
}

func (ptr *PlainTextEdit) AddLine(v0 string) {
	if ptr.Pointer() != nil {
		var v0C *C.char
		if v0 != "" {
			v0C = C.CString(v0)
			defer C.free(unsafe.Pointer(v0C))
		}
		C.PlainTextEdit209179_AddLine(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(len(v0))})
	}
}

func PlainTextEdit_QRegisterMetaType() int {
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType()))
}

func (ptr *PlainTextEdit) QRegisterMetaType() int {
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType()))
}

func PlainTextEdit_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType2(typeNameC)))
}

func (ptr *PlainTextEdit) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType2(typeNameC)))
}

func PlainTextEdit_QmlRegisterType() int {
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType()))
}

func (ptr *PlainTextEdit) QmlRegisterType() int {
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType()))
}

func PlainTextEdit_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *PlainTextEdit) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func PlainTextEdit_QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func (ptr *PlainTextEdit) QmlRegisterUncreatableType(uri string, versionMajor int, versionMinor int, qmlName string, message string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	var messageC *C.char
	if message != "" {
		messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterUncreatableType(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC, C.struct_Moc_PackedString{data: messageC, len: C.longlong(len(message))})))
}

func PlainTextEdit_QmlRegisterAnonymousType(uri string, versionMajor int) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterAnonymousType(uriC, C.int(int32(versionMajor)))))
}

func (ptr *PlainTextEdit) QmlRegisterAnonymousType(uri string, versionMajor int) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	return int(int32(C.PlainTextEdit209179_PlainTextEdit209179_QmlRegisterAnonymousType(uriC, C.int(int32(versionMajor)))))
}

func (ptr *PlainTextEdit) __scrollBarWidgets_atList(i int) *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQWidgetFromPointer(C.PlainTextEdit209179___scrollBarWidgets_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __scrollBarWidgets_setList(i std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___scrollBarWidgets_setList(ptr.Pointer(), std_widgets.PointerFromQWidget(i))
	}
}

func (ptr *PlainTextEdit) __scrollBarWidgets_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___scrollBarWidgets_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.PlainTextEdit209179___actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *PlainTextEdit) __actions_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___actions_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __addActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.PlainTextEdit209179___addActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __addActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___addActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *PlainTextEdit) __addActions_actions_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___addActions_actions_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __insertActions_actions_atList(i int) *std_widgets.QAction {
	if ptr.Pointer() != nil {
		tmpValue := std_widgets.NewQActionFromPointer(C.PlainTextEdit209179___insertActions_actions_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __insertActions_actions_setList(i std_widgets.QAction_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___insertActions_actions_setList(ptr.Pointer(), std_widgets.PointerFromQAction(i))
	}
}

func (ptr *PlainTextEdit) __insertActions_actions_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___insertActions_actions_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PlainTextEdit209179___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PlainTextEdit) __children_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___children_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.PlainTextEdit209179___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *PlainTextEdit) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PlainTextEdit209179___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PlainTextEdit) __findChildren_newList() unsafe.Pointer {
	return C.PlainTextEdit209179___findChildren_newList(ptr.Pointer())
}

func (ptr *PlainTextEdit) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.PlainTextEdit209179___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *PlainTextEdit) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *PlainTextEdit) __findChildren_newList3() unsafe.Pointer {
	return C.PlainTextEdit209179___findChildren_newList3(ptr.Pointer())
}

func NewPlainTextEdit(parent std_widgets.QWidget_ITF) *PlainTextEdit {
	PlainTextEdit_QRegisterMetaType()
	tmpValue := NewPlainTextEditFromPointer(C.PlainTextEdit209179_NewPlainTextEdit(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewPlainTextEdit2(text string, parent std_widgets.QWidget_ITF) *PlainTextEdit {
	var textC *C.char
	if text != "" {
		textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
	}
	PlainTextEdit_QRegisterMetaType()
	tmpValue := NewPlainTextEditFromPointer(C.PlainTextEdit209179_NewPlainTextEdit2(C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))}, std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackPlainTextEdit209179_DestroyPlainTextEdit
func callbackPlainTextEdit209179_DestroyPlainTextEdit(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~PlainTextEdit"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).DestroyPlainTextEditDefault()
	}
}

func (ptr *PlainTextEdit) ConnectDestroyPlainTextEdit(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~PlainTextEdit"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~PlainTextEdit", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~PlainTextEdit", unsafe.Pointer(&f))
		}
	}
}

func (ptr *PlainTextEdit) DisconnectDestroyPlainTextEdit() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~PlainTextEdit")
	}
}

func (ptr *PlainTextEdit) DestroyPlainTextEdit() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PlainTextEdit209179_DestroyPlainTextEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *PlainTextEdit) DestroyPlainTextEditDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PlainTextEdit209179_DestroyPlainTextEditDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackPlainTextEdit209179_AppendHtml
func callbackPlainTextEdit209179_AppendHtml(ptr unsafe.Pointer, html C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "appendHtml"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(html))
	} else {
		NewPlainTextEditFromPointer(ptr).AppendHtmlDefault(cGoUnpackString(html))
	}
}

func (ptr *PlainTextEdit) AppendHtmlDefault(html string) {
	if ptr.Pointer() != nil {
		var htmlC *C.char
		if html != "" {
			htmlC = C.CString(html)
			defer C.free(unsafe.Pointer(htmlC))
		}
		C.PlainTextEdit209179_AppendHtmlDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: htmlC, len: C.longlong(len(html))})
	}
}

//export callbackPlainTextEdit209179_AppendPlainText
func callbackPlainTextEdit209179_AppendPlainText(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "appendPlainText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewPlainTextEditFromPointer(ptr).AppendPlainTextDefault(cGoUnpackString(text))
	}
}

func (ptr *PlainTextEdit) AppendPlainTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.PlainTextEdit209179_AppendPlainTextDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackPlainTextEdit209179_BlockCountChanged
func callbackPlainTextEdit209179_BlockCountChanged(ptr unsafe.Pointer, newBlockCount C.int) {
	if signal := qt.GetSignal(ptr, "blockCountChanged"); signal != nil {
		(*(*func(int))(signal))(int(int32(newBlockCount)))
	}

}

//export callbackPlainTextEdit209179_CanInsertFromMimeData
func callbackPlainTextEdit209179_CanInsertFromMimeData(ptr unsafe.Pointer, source unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canInsertFromMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QMimeData) bool)(signal))(std_core.NewQMimeDataFromPointer(source)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).CanInsertFromMimeDataDefault(std_core.NewQMimeDataFromPointer(source)))))
}

func (ptr *PlainTextEdit) CanInsertFromMimeDataDefault(source std_core.QMimeData_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_CanInsertFromMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(source))) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_CenterCursor
func callbackPlainTextEdit209179_CenterCursor(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "centerCursor"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).CenterCursorDefault()
	}
}

func (ptr *PlainTextEdit) CenterCursorDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_CenterCursorDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ChangeEvent
func callbackPlainTextEdit209179_ChangeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "changeEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).ChangeEventDefault(std_core.NewQEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) ChangeEventDefault(e std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ChangeEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))
	}
}

//export callbackPlainTextEdit209179_Clear
func callbackPlainTextEdit209179_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ClearDefault()
	}
}

func (ptr *PlainTextEdit) ClearDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ClearDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ContextMenuEvent
func callbackPlainTextEdit209179_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "contextMenuEvent"); signal != nil {
		(*(*func(*std_gui.QContextMenuEvent))(signal))(std_gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).ContextMenuEventDefault(std_gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) ContextMenuEventDefault(event std_gui.QContextMenuEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ContextMenuEventDefault(ptr.Pointer(), std_gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackPlainTextEdit209179_Copy
func callbackPlainTextEdit209179_Copy(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "copy"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).CopyDefault()
	}
}

func (ptr *PlainTextEdit) CopyDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_CopyDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_CopyAvailable
func callbackPlainTextEdit209179_CopyAvailable(ptr unsafe.Pointer, yes C.char) {
	if signal := qt.GetSignal(ptr, "copyAvailable"); signal != nil {
		(*(*func(bool))(signal))(int8(yes) != 0)
	}

}

//export callbackPlainTextEdit209179_CreateMimeDataFromSelection
func callbackPlainTextEdit209179_CreateMimeDataFromSelection(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createMimeDataFromSelection"); signal != nil {
		return std_core.PointerFromQMimeData((*(*func() *std_core.QMimeData)(signal))())
	}

	return std_core.PointerFromQMimeData(NewPlainTextEditFromPointer(ptr).CreateMimeDataFromSelectionDefault())
}

func (ptr *PlainTextEdit) CreateMimeDataFromSelectionDefault() *std_core.QMimeData {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQMimeDataFromPointer(C.PlainTextEdit209179_CreateMimeDataFromSelectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_CursorPositionChanged
func callbackPlainTextEdit209179_CursorPositionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cursorPositionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackPlainTextEdit209179_Cut
func callbackPlainTextEdit209179_Cut(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "cut"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).CutDefault()
	}
}

func (ptr *PlainTextEdit) CutDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_CutDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_DragEnterEvent
func callbackPlainTextEdit209179_DragEnterEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) DragEnterEventDefault(e std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(e))
	}
}

//export callbackPlainTextEdit209179_DragLeaveEvent
func callbackPlainTextEdit209179_DragLeaveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) DragLeaveEventDefault(e std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(e))
	}
}

//export callbackPlainTextEdit209179_DragMoveEvent
func callbackPlainTextEdit209179_DragMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) DragMoveEventDefault(e std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(e))
	}
}

//export callbackPlainTextEdit209179_DropEvent
func callbackPlainTextEdit209179_DropEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) DropEventDefault(e std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(e))
	}
}

//export callbackPlainTextEdit209179_FocusInEvent
func callbackPlainTextEdit209179_FocusInEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) FocusInEventDefault(e std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(e))
	}
}

//export callbackPlainTextEdit209179_FocusNextPrevChild
func callbackPlainTextEdit209179_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	if signal := qt.GetSignal(ptr, "focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(bool) bool)(signal))(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *PlainTextEdit) FocusNextPrevChildDefault(next bool) bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next))))) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_FocusOutEvent
func callbackPlainTextEdit209179_FocusOutEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) FocusOutEventDefault(e std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(e))
	}
}

//export callbackPlainTextEdit209179_InputMethodEvent
func callbackPlainTextEdit209179_InputMethodEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) InputMethodEventDefault(e std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(e))
	}
}

//export callbackPlainTextEdit209179_InputMethodQuery
func callbackPlainTextEdit209179_InputMethodQuery(ptr unsafe.Pointer, property C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(property)))
	}

	return std_core.PointerFromQVariant(NewPlainTextEditFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(property)))
}

func (ptr *PlainTextEdit) InputMethodQueryDefault(property std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PlainTextEdit209179_InputMethodQueryDefault(ptr.Pointer(), C.longlong(property)))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_InsertFromMimeData
func callbackPlainTextEdit209179_InsertFromMimeData(ptr unsafe.Pointer, source unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "insertFromMimeData"); signal != nil {
		(*(*func(*std_core.QMimeData))(signal))(std_core.NewQMimeDataFromPointer(source))
	} else {
		NewPlainTextEditFromPointer(ptr).InsertFromMimeDataDefault(std_core.NewQMimeDataFromPointer(source))
	}
}

func (ptr *PlainTextEdit) InsertFromMimeDataDefault(source std_core.QMimeData_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_InsertFromMimeDataDefault(ptr.Pointer(), std_core.PointerFromQMimeData(source))
	}
}

//export callbackPlainTextEdit209179_InsertPlainText
func callbackPlainTextEdit209179_InsertPlainText(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "insertPlainText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewPlainTextEditFromPointer(ptr).InsertPlainTextDefault(cGoUnpackString(text))
	}
}

func (ptr *PlainTextEdit) InsertPlainTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.PlainTextEdit209179_InsertPlainTextDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackPlainTextEdit209179_KeyPressEvent
func callbackPlainTextEdit209179_KeyPressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) KeyPressEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackPlainTextEdit209179_KeyReleaseEvent
func callbackPlainTextEdit209179_KeyReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) KeyReleaseEventDefault(e std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(e))
	}
}

//export callbackPlainTextEdit209179_LoadResource
func callbackPlainTextEdit209179_LoadResource(ptr unsafe.Pointer, ty C.int, name unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "loadResource"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(int, *std_core.QUrl) *std_core.QVariant)(signal))(int(int32(ty)), std_core.NewQUrlFromPointer(name)))
	}

	return std_core.PointerFromQVariant(NewPlainTextEditFromPointer(ptr).LoadResourceDefault(int(int32(ty)), std_core.NewQUrlFromPointer(name)))
}

func (ptr *PlainTextEdit) LoadResourceDefault(ty int, name std_core.QUrl_ITF) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.PlainTextEdit209179_LoadResourceDefault(ptr.Pointer(), C.int(int32(ty)), std_core.PointerFromQUrl(name)))
		qt.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_ModificationChanged
func callbackPlainTextEdit209179_ModificationChanged(ptr unsafe.Pointer, changed C.char) {
	if signal := qt.GetSignal(ptr, "modificationChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(changed) != 0)
	}

}

//export callbackPlainTextEdit209179_MouseDoubleClickEvent
func callbackPlainTextEdit209179_MouseDoubleClickEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) MouseDoubleClickEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackPlainTextEdit209179_MouseMoveEvent
func callbackPlainTextEdit209179_MouseMoveEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) MouseMoveEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackPlainTextEdit209179_MousePressEvent
func callbackPlainTextEdit209179_MousePressEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) MousePressEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackPlainTextEdit209179_MouseReleaseEvent
func callbackPlainTextEdit209179_MouseReleaseEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) MouseReleaseEventDefault(e std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(e))
	}
}

//export callbackPlainTextEdit209179_PaintEvent
func callbackPlainTextEdit209179_PaintEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paintEvent"); signal != nil {
		(*(*func(*std_gui.QPaintEvent))(signal))(std_gui.NewQPaintEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).PaintEventDefault(std_gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) PaintEventDefault(e std_gui.QPaintEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_PaintEventDefault(ptr.Pointer(), std_gui.PointerFromQPaintEvent(e))
	}
}

//export callbackPlainTextEdit209179_Paste
func callbackPlainTextEdit209179_Paste(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "paste"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).PasteDefault()
	}
}

func (ptr *PlainTextEdit) PasteDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_PasteDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_Redo
func callbackPlainTextEdit209179_Redo(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "redo"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).RedoDefault()
	}
}

func (ptr *PlainTextEdit) RedoDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_RedoDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_RedoAvailable
func callbackPlainTextEdit209179_RedoAvailable(ptr unsafe.Pointer, available C.char) {
	if signal := qt.GetSignal(ptr, "redoAvailable"); signal != nil {
		(*(*func(bool))(signal))(int8(available) != 0)
	}

}

//export callbackPlainTextEdit209179_ResizeEvent
func callbackPlainTextEdit209179_ResizeEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resizeEvent"); signal != nil {
		(*(*func(*std_gui.QResizeEvent))(signal))(std_gui.NewQResizeEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).ResizeEventDefault(std_gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) ResizeEventDefault(e std_gui.QResizeEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ResizeEventDefault(ptr.Pointer(), std_gui.PointerFromQResizeEvent(e))
	}
}

//export callbackPlainTextEdit209179_ScrollContentsBy
func callbackPlainTextEdit209179_ScrollContentsBy(ptr unsafe.Pointer, dx C.int, dy C.int) {
	if signal := qt.GetSignal(ptr, "scrollContentsBy"); signal != nil {
		(*(*func(int, int))(signal))(int(int32(dx)), int(int32(dy)))
	} else {
		NewPlainTextEditFromPointer(ptr).ScrollContentsByDefault(int(int32(dx)), int(int32(dy)))
	}
}

func (ptr *PlainTextEdit) ScrollContentsByDefault(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ScrollContentsByDefault(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

//export callbackPlainTextEdit209179_SelectAll
func callbackPlainTextEdit209179_SelectAll(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectAll"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).SelectAllDefault()
	}
}

func (ptr *PlainTextEdit) SelectAllDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SelectAllDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_SelectionChanged
func callbackPlainTextEdit209179_SelectionChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "selectionChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackPlainTextEdit209179_SetPlainText
func callbackPlainTextEdit209179_SetPlainText(ptr unsafe.Pointer, text C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setPlainText"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(text))
	} else {
		NewPlainTextEditFromPointer(ptr).SetPlainTextDefault(cGoUnpackString(text))
	}
}

func (ptr *PlainTextEdit) SetPlainTextDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.PlainTextEdit209179_SetPlainTextDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackPlainTextEdit209179_ShowEvent
func callbackPlainTextEdit209179_ShowEvent(ptr unsafe.Pointer, vqs unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showEvent"); signal != nil {
		(*(*func(*std_gui.QShowEvent))(signal))(std_gui.NewQShowEventFromPointer(vqs))
	} else {
		NewPlainTextEditFromPointer(ptr).ShowEventDefault(std_gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *PlainTextEdit) ShowEventDefault(vqs std_gui.QShowEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowEventDefault(ptr.Pointer(), std_gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackPlainTextEdit209179_TextChanged
func callbackPlainTextEdit209179_TextChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "textChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackPlainTextEdit209179_Undo
func callbackPlainTextEdit209179_Undo(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "undo"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).UndoDefault()
	}
}

func (ptr *PlainTextEdit) UndoDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_UndoDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_UndoAvailable
func callbackPlainTextEdit209179_UndoAvailable(ptr unsafe.Pointer, available C.char) {
	if signal := qt.GetSignal(ptr, "undoAvailable"); signal != nil {
		(*(*func(bool))(signal))(int8(available) != 0)
	}

}

//export callbackPlainTextEdit209179_UpdateRequest
func callbackPlainTextEdit209179_UpdateRequest(ptr unsafe.Pointer, rect unsafe.Pointer, dy C.int) {
	if signal := qt.GetSignal(ptr, "updateRequest"); signal != nil {
		(*(*func(*std_core.QRect, int))(signal))(std_core.NewQRectFromPointer(rect), int(int32(dy)))
	}

}

//export callbackPlainTextEdit209179_WheelEvent
func callbackPlainTextEdit209179_WheelEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(e))
	} else {
		NewPlainTextEditFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *PlainTextEdit) WheelEventDefault(e std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(e))
	}
}

//export callbackPlainTextEdit209179_ZoomIn
func callbackPlainTextEdit209179_ZoomIn(ptr unsafe.Pointer, ran C.int) {
	if signal := qt.GetSignal(ptr, "zoomIn"); signal != nil {
		(*(*func(int))(signal))(int(int32(ran)))
	} else {
		NewPlainTextEditFromPointer(ptr).ZoomInDefault(int(int32(ran)))
	}
}

func (ptr *PlainTextEdit) ZoomInDefault(ran int) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ZoomInDefault(ptr.Pointer(), C.int(int32(ran)))
	}
}

//export callbackPlainTextEdit209179_ZoomOut
func callbackPlainTextEdit209179_ZoomOut(ptr unsafe.Pointer, ran C.int) {
	if signal := qt.GetSignal(ptr, "zoomOut"); signal != nil {
		(*(*func(int))(signal))(int(int32(ran)))
	} else {
		NewPlainTextEditFromPointer(ptr).ZoomOutDefault(int(int32(ran)))
	}
}

func (ptr *PlainTextEdit) ZoomOutDefault(ran int) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ZoomOutDefault(ptr.Pointer(), C.int(int32(ran)))
	}
}

//export callbackPlainTextEdit209179_Event
func callbackPlainTextEdit209179_Event(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *PlainTextEdit) EventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_MinimumSizeHint
func callbackPlainTextEdit209179_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewPlainTextEditFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *PlainTextEdit) MinimumSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.PlainTextEdit209179_MinimumSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_SetupViewport
func callbackPlainTextEdit209179_SetupViewport(ptr unsafe.Pointer, viewport unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setupViewport"); signal != nil {
		(*(*func(*std_widgets.QWidget))(signal))(std_widgets.NewQWidgetFromPointer(viewport))
	} else {
		NewPlainTextEditFromPointer(ptr).SetupViewportDefault(std_widgets.NewQWidgetFromPointer(viewport))
	}
}

func (ptr *PlainTextEdit) SetupViewportDefault(viewport std_widgets.QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetupViewportDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(viewport))
	}
}

//export callbackPlainTextEdit209179_SizeHint
func callbackPlainTextEdit209179_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewPlainTextEditFromPointer(ptr).SizeHintDefault())
}

func (ptr *PlainTextEdit) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.PlainTextEdit209179_SizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_ViewportEvent
func callbackPlainTextEdit209179_ViewportEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "viewportEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).ViewportEventDefault(std_core.NewQEventFromPointer(event)))))
}

func (ptr *PlainTextEdit) ViewportEventDefault(event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_ViewportEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_ViewportSizeHint
func callbackPlainTextEdit209179_ViewportSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "viewportSizeHint"); signal != nil {
		return std_core.PointerFromQSize((*(*func() *std_core.QSize)(signal))())
	}

	return std_core.PointerFromQSize(NewPlainTextEditFromPointer(ptr).ViewportSizeHintDefault())
}

func (ptr *PlainTextEdit) ViewportSizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQSizeFromPointer(C.PlainTextEdit209179_ViewportSizeHintDefault(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackPlainTextEdit209179_ActionEvent
func callbackPlainTextEdit209179_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "actionEvent"); signal != nil {
		(*(*func(*std_gui.QActionEvent))(signal))(std_gui.NewQActionEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).ActionEventDefault(std_gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) ActionEventDefault(event std_gui.QActionEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ActionEventDefault(ptr.Pointer(), std_gui.PointerFromQActionEvent(event))
	}
}

//export callbackPlainTextEdit209179_Close
func callbackPlainTextEdit209179_Close(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).CloseDefault())))
}

func (ptr *PlainTextEdit) CloseDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_CloseDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_CloseEvent
func callbackPlainTextEdit209179_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "closeEvent"); signal != nil {
		(*(*func(*std_gui.QCloseEvent))(signal))(std_gui.NewQCloseEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).CloseEventDefault(std_gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) CloseEventDefault(event std_gui.QCloseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_CloseEventDefault(ptr.Pointer(), std_gui.PointerFromQCloseEvent(event))
	}
}

//export callbackPlainTextEdit209179_CustomContextMenuRequested
func callbackPlainTextEdit209179_CustomContextMenuRequested(ptr unsafe.Pointer, pos unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customContextMenuRequested"); signal != nil {
		(*(*func(*std_core.QPoint))(signal))(std_core.NewQPointFromPointer(pos))
	}

}

//export callbackPlainTextEdit209179_EnterEvent
func callbackPlainTextEdit209179_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enterEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).EnterEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) EnterEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_EnterEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPlainTextEdit209179_HasHeightForWidth
func callbackPlainTextEdit209179_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *PlainTextEdit) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_HasHeightForWidthDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_HeightForWidth
func callbackPlainTextEdit209179_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32((*(*func(int) int)(signal))(int(int32(w)))))
	}

	return C.int(int32(NewPlainTextEditFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *PlainTextEdit) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PlainTextEdit209179_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackPlainTextEdit209179_Hide
func callbackPlainTextEdit209179_Hide(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hide"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).HideDefault()
	}
}

func (ptr *PlainTextEdit) HideDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_HideDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_HideEvent
func callbackPlainTextEdit209179_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hideEvent"); signal != nil {
		(*(*func(*std_gui.QHideEvent))(signal))(std_gui.NewQHideEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).HideEventDefault(std_gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) HideEventDefault(event std_gui.QHideEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_HideEventDefault(ptr.Pointer(), std_gui.PointerFromQHideEvent(event))
	}
}

//export callbackPlainTextEdit209179_InitPainter
func callbackPlainTextEdit209179_InitPainter(ptr unsafe.Pointer, painter unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "initPainter"); signal != nil {
		(*(*func(*std_gui.QPainter))(signal))(std_gui.NewQPainterFromPointer(painter))
	} else {
		NewPlainTextEditFromPointer(ptr).InitPainterDefault(std_gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *PlainTextEdit) InitPainterDefault(painter std_gui.QPainter_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_InitPainterDefault(ptr.Pointer(), std_gui.PointerFromQPainter(painter))
	}
}

//export callbackPlainTextEdit209179_LeaveEvent
func callbackPlainTextEdit209179_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "leaveEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).LeaveEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) LeaveEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_LeaveEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPlainTextEdit209179_Lower
func callbackPlainTextEdit209179_Lower(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "lower"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).LowerDefault()
	}
}

func (ptr *PlainTextEdit) LowerDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_LowerDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_Metric
func callbackPlainTextEdit209179_Metric(ptr unsafe.Pointer, m C.longlong) C.int {
	if signal := qt.GetSignal(ptr, "metric"); signal != nil {
		return C.int(int32((*(*func(std_gui.QPaintDevice__PaintDeviceMetric) int)(signal))(std_gui.QPaintDevice__PaintDeviceMetric(m))))
	}

	return C.int(int32(NewPlainTextEditFromPointer(ptr).MetricDefault(std_gui.QPaintDevice__PaintDeviceMetric(m))))
}

func (ptr *PlainTextEdit) MetricDefault(m std_gui.QPaintDevice__PaintDeviceMetric) int {
	if ptr.Pointer() != nil {
		return int(int32(C.PlainTextEdit209179_MetricDefault(ptr.Pointer(), C.longlong(m))))
	}
	return 0
}

//export callbackPlainTextEdit209179_MoveEvent
func callbackPlainTextEdit209179_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "moveEvent"); signal != nil {
		(*(*func(*std_gui.QMoveEvent))(signal))(std_gui.NewQMoveEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).MoveEventDefault(std_gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) MoveEventDefault(event std_gui.QMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_MoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMoveEvent(event))
	}
}

//export callbackPlainTextEdit209179_NativeEvent
func callbackPlainTextEdit209179_NativeEvent(ptr unsafe.Pointer, eventType unsafe.Pointer, message unsafe.Pointer, result *C.long) C.char {
	var resultR int
	if result != nil {
		resultR = int(int32(*result))
		defer func() { *result = C.long(int32(resultR)) }()
	}
	if signal := qt.GetSignal(ptr, "nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QByteArray, unsafe.Pointer, *int) bool)(signal))(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).NativeEventDefault(std_core.NewQByteArrayFromPointer(eventType), message, &resultR))))
}

func (ptr *PlainTextEdit) NativeEventDefault(eventType std_core.QByteArray_ITF, message unsafe.Pointer, result *int) bool {
	if ptr.Pointer() != nil {
		var resultC C.long
		if result != nil {
			resultC = C.long(int32(*result))
			defer func() { *result = int(int32(resultC)) }()
		}
		return int8(C.PlainTextEdit209179_NativeEventDefault(ptr.Pointer(), std_core.PointerFromQByteArray(eventType), message, &resultC)) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_PaintEngine
func callbackPlainTextEdit209179_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return std_gui.PointerFromQPaintEngine((*(*func() *std_gui.QPaintEngine)(signal))())
	}

	return std_gui.PointerFromQPaintEngine(NewPlainTextEditFromPointer(ptr).PaintEngineDefault())
}

func (ptr *PlainTextEdit) PaintEngineDefault() *std_gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return std_gui.NewQPaintEngineFromPointer(C.PlainTextEdit209179_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackPlainTextEdit209179_Raise
func callbackPlainTextEdit209179_Raise(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "raise"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *PlainTextEdit) RaiseDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_RaiseDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_Repaint
func callbackPlainTextEdit209179_Repaint(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "repaint"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *PlainTextEdit) RepaintDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_RepaintDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_SetDisabled
func callbackPlainTextEdit209179_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	if signal := qt.GetSignal(ptr, "setDisabled"); signal != nil {
		(*(*func(bool))(signal))(int8(disable) != 0)
	} else {
		NewPlainTextEditFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *PlainTextEdit) SetDisabledDefault(disable bool) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackPlainTextEdit209179_SetEnabled
func callbackPlainTextEdit209179_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setEnabled"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewPlainTextEditFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *PlainTextEdit) SetEnabledDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackPlainTextEdit209179_SetFocus2
func callbackPlainTextEdit209179_SetFocus2(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setFocus2"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *PlainTextEdit) SetFocus2Default() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_SetHidden
func callbackPlainTextEdit209179_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	if signal := qt.GetSignal(ptr, "setHidden"); signal != nil {
		(*(*func(bool))(signal))(int8(hidden) != 0)
	} else {
		NewPlainTextEditFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *PlainTextEdit) SetHiddenDefault(hidden bool) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackPlainTextEdit209179_SetStyleSheet
func callbackPlainTextEdit209179_SetStyleSheet(ptr unsafe.Pointer, styleSheet C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStyleSheet"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(styleSheet))
	} else {
		NewPlainTextEditFromPointer(ptr).SetStyleSheetDefault(cGoUnpackString(styleSheet))
	}
}

func (ptr *PlainTextEdit) SetStyleSheetDefault(styleSheet string) {
	if ptr.Pointer() != nil {
		var styleSheetC *C.char
		if styleSheet != "" {
			styleSheetC = C.CString(styleSheet)
			defer C.free(unsafe.Pointer(styleSheetC))
		}
		C.PlainTextEdit209179_SetStyleSheetDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: styleSheetC, len: C.longlong(len(styleSheet))})
	}
}

//export callbackPlainTextEdit209179_SetVisible
func callbackPlainTextEdit209179_SetVisible(ptr unsafe.Pointer, visible C.char) {
	if signal := qt.GetSignal(ptr, "setVisible"); signal != nil {
		(*(*func(bool))(signal))(int8(visible) != 0)
	} else {
		NewPlainTextEditFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *PlainTextEdit) SetVisibleDefault(visible bool) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackPlainTextEdit209179_SetWindowModified
func callbackPlainTextEdit209179_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "setWindowModified"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	} else {
		NewPlainTextEditFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *PlainTextEdit) SetWindowModifiedDefault(vbo bool) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackPlainTextEdit209179_SetWindowTitle
func callbackPlainTextEdit209179_SetWindowTitle(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setWindowTitle"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	} else {
		NewPlainTextEditFromPointer(ptr).SetWindowTitleDefault(cGoUnpackString(vqs))
	}
}

func (ptr *PlainTextEdit) SetWindowTitleDefault(vqs string) {
	if ptr.Pointer() != nil {
		var vqsC *C.char
		if vqs != "" {
			vqsC = C.CString(vqs)
			defer C.free(unsafe.Pointer(vqsC))
		}
		C.PlainTextEdit209179_SetWindowTitleDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: vqsC, len: C.longlong(len(vqs))})
	}
}

//export callbackPlainTextEdit209179_Show
func callbackPlainTextEdit209179_Show(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "show"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ShowDefault()
	}
}

func (ptr *PlainTextEdit) ShowDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ShowFullScreen
func callbackPlainTextEdit209179_ShowFullScreen(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showFullScreen"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *PlainTextEdit) ShowFullScreenDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ShowMaximized
func callbackPlainTextEdit209179_ShowMaximized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMaximized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *PlainTextEdit) ShowMaximizedDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ShowMinimized
func callbackPlainTextEdit209179_ShowMinimized(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showMinimized"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *PlainTextEdit) ShowMinimizedDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_ShowNormal
func callbackPlainTextEdit209179_ShowNormal(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "showNormal"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *PlainTextEdit) ShowNormalDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_TabletEvent
func callbackPlainTextEdit209179_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "tabletEvent"); signal != nil {
		(*(*func(*std_gui.QTabletEvent))(signal))(std_gui.NewQTabletEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).TabletEventDefault(std_gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) TabletEventDefault(event std_gui.QTabletEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_TabletEventDefault(ptr.Pointer(), std_gui.PointerFromQTabletEvent(event))
	}
}

//export callbackPlainTextEdit209179_Update
func callbackPlainTextEdit209179_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *PlainTextEdit) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_UpdateDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_UpdateMicroFocus
func callbackPlainTextEdit209179_UpdateMicroFocus(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateMicroFocus"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *PlainTextEdit) UpdateMicroFocusDefault() {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_WindowIconChanged
func callbackPlainTextEdit209179_WindowIconChanged(ptr unsafe.Pointer, icon unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowIconChanged"); signal != nil {
		(*(*func(*std_gui.QIcon))(signal))(std_gui.NewQIconFromPointer(icon))
	}

}

//export callbackPlainTextEdit209179_WindowTitleChanged
func callbackPlainTextEdit209179_WindowTitleChanged(ptr unsafe.Pointer, title C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "windowTitleChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(title))
	}

}

//export callbackPlainTextEdit209179_ChildEvent
func callbackPlainTextEdit209179_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackPlainTextEdit209179_ConnectNotify
func callbackPlainTextEdit209179_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPlainTextEditFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PlainTextEdit) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPlainTextEdit209179_CustomEvent
func callbackPlainTextEdit209179_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackPlainTextEdit209179_DeleteLater
func callbackPlainTextEdit209179_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewPlainTextEditFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *PlainTextEdit) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.PlainTextEdit209179_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackPlainTextEdit209179_Destroyed
func callbackPlainTextEdit209179_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackPlainTextEdit209179_DisconnectNotify
func callbackPlainTextEdit209179_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewPlainTextEditFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *PlainTextEdit) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackPlainTextEdit209179_EventFilter
func callbackPlainTextEdit209179_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewPlainTextEditFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *PlainTextEdit) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.PlainTextEdit209179_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackPlainTextEdit209179_ObjectNameChanged
func callbackPlainTextEdit209179_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackPlainTextEdit209179_TimerEvent
func callbackPlainTextEdit209179_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewPlainTextEditFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *PlainTextEdit) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.PlainTextEdit209179_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
}
