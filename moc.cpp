

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPlainTextEdit>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class PlainTextEdit209179: public QPlainTextEdit
{
Q_OBJECT
public:
	PlainTextEdit209179(QWidget *parent = Q_NULLPTR) : QPlainTextEdit(parent) {qRegisterMetaType<quintptr>("quintptr");PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType();PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaTypes();callbackPlainTextEdit209179_Constructor(this);};
	PlainTextEdit209179(const QString &text, QWidget *parent = Q_NULLPTR) : QPlainTextEdit(text, parent) {qRegisterMetaType<quintptr>("quintptr");PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType();PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaTypes();callbackPlainTextEdit209179_Constructor(this);};
	 ~PlainTextEdit209179() { callbackPlainTextEdit209179_DestroyPlainTextEdit(this); };
	void appendHtml(const QString & html) { QByteArray* t950a39 = new QByteArray(html.toUtf8()); Moc_PackedString htmlPacked = { const_cast<char*>(t950a39->prepend("WHITESPACE").constData()+10), t950a39->size()-10, t950a39 };callbackPlainTextEdit209179_AppendHtml(this, htmlPacked); };
	void appendPlainText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); Moc_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackPlainTextEdit209179_AppendPlainText(this, textPacked); };
	void Signal_BlockCountChanged(int newBlockCount) { callbackPlainTextEdit209179_BlockCountChanged(this, newBlockCount); };
	bool canInsertFromMimeData(const QMimeData * source) const { return callbackPlainTextEdit209179_CanInsertFromMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(source)) != 0; };
	void centerCursor() { callbackPlainTextEdit209179_CenterCursor(this); };
	void changeEvent(QEvent * e) { callbackPlainTextEdit209179_ChangeEvent(this, e); };
	void clear() { callbackPlainTextEdit209179_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackPlainTextEdit209179_ContextMenuEvent(this, event); };
	void copy() { callbackPlainTextEdit209179_Copy(this); };
	void Signal_CopyAvailable(bool yes) { callbackPlainTextEdit209179_CopyAvailable(this, yes); };
	QMimeData * createMimeDataFromSelection() const { return static_cast<QMimeData*>(callbackPlainTextEdit209179_CreateMimeDataFromSelection(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_CursorPositionChanged() { callbackPlainTextEdit209179_CursorPositionChanged(this); };
	void cut() { callbackPlainTextEdit209179_Cut(this); };
	void dragEnterEvent(QDragEnterEvent * e) { callbackPlainTextEdit209179_DragEnterEvent(this, e); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackPlainTextEdit209179_DragLeaveEvent(this, e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackPlainTextEdit209179_DragMoveEvent(this, e); };
	void dropEvent(QDropEvent * e) { callbackPlainTextEdit209179_DropEvent(this, e); };
	void focusInEvent(QFocusEvent * e) { callbackPlainTextEdit209179_FocusInEvent(this, e); };
	bool focusNextPrevChild(bool next) { return callbackPlainTextEdit209179_FocusNextPrevChild(this, next) != 0; };
	void focusOutEvent(QFocusEvent * e) { callbackPlainTextEdit209179_FocusOutEvent(this, e); };
	void inputMethodEvent(QInputMethodEvent * e) { callbackPlainTextEdit209179_InputMethodEvent(this, e); };
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const { return *static_cast<QVariant*>(callbackPlainTextEdit209179_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), property)); };
	void insertFromMimeData(const QMimeData * source) { callbackPlainTextEdit209179_InsertFromMimeData(this, const_cast<QMimeData*>(source)); };
	void insertPlainText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); Moc_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackPlainTextEdit209179_InsertPlainText(this, textPacked); };
	void keyPressEvent(QKeyEvent * e) { callbackPlainTextEdit209179_KeyPressEvent(this, e); };
	void keyReleaseEvent(QKeyEvent * e) { callbackPlainTextEdit209179_KeyReleaseEvent(this, e); };
	QVariant loadResource(int ty, const QUrl & name) { return *static_cast<QVariant*>(callbackPlainTextEdit209179_LoadResource(this, ty, const_cast<QUrl*>(&name))); };
	void Signal_ModificationChanged(bool changed) { callbackPlainTextEdit209179_ModificationChanged(this, changed); };
	void mouseDoubleClickEvent(QMouseEvent * e) { callbackPlainTextEdit209179_MouseDoubleClickEvent(this, e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackPlainTextEdit209179_MouseMoveEvent(this, e); };
	void mousePressEvent(QMouseEvent * e) { callbackPlainTextEdit209179_MousePressEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackPlainTextEdit209179_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackPlainTextEdit209179_PaintEvent(this, e); };
	void paste() { callbackPlainTextEdit209179_Paste(this); };
	void redo() { callbackPlainTextEdit209179_Redo(this); };
	void Signal_RedoAvailable(bool available) { callbackPlainTextEdit209179_RedoAvailable(this, available); };
	void resizeEvent(QResizeEvent * e) { callbackPlainTextEdit209179_ResizeEvent(this, e); };
	void scrollContentsBy(int dx, int dy) { callbackPlainTextEdit209179_ScrollContentsBy(this, dx, dy); };
	void selectAll() { callbackPlainTextEdit209179_SelectAll(this); };
	void Signal_SelectionChanged() { callbackPlainTextEdit209179_SelectionChanged(this); };
	void setPlainText(const QString & text) { QByteArray* t372ea0 = new QByteArray(text.toUtf8()); Moc_PackedString textPacked = { const_cast<char*>(t372ea0->prepend("WHITESPACE").constData()+10), t372ea0->size()-10, t372ea0 };callbackPlainTextEdit209179_SetPlainText(this, textPacked); };
	void showEvent(QShowEvent * vqs) { callbackPlainTextEdit209179_ShowEvent(this, vqs); };
	void Signal_TextChanged() { callbackPlainTextEdit209179_TextChanged(this); };
	void undo() { callbackPlainTextEdit209179_Undo(this); };
	void Signal_UndoAvailable(bool available) { callbackPlainTextEdit209179_UndoAvailable(this, available); };
	void Signal_UpdateRequest(const QRect & rect, int dy) { callbackPlainTextEdit209179_UpdateRequest(this, const_cast<QRect*>(&rect), dy); };
	void wheelEvent(QWheelEvent * e) { callbackPlainTextEdit209179_WheelEvent(this, e); };
	void zoomIn(int ran) { callbackPlainTextEdit209179_ZoomIn(this, ran); };
	void zoomOut(int ran) { callbackPlainTextEdit209179_ZoomOut(this, ran); };
	bool event(QEvent * event) { return callbackPlainTextEdit209179_Event(this, event) != 0; };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackPlainTextEdit209179_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void setupViewport(QWidget * viewport) { callbackPlainTextEdit209179_SetupViewport(this, viewport); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackPlainTextEdit209179_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool viewportEvent(QEvent * event) { return callbackPlainTextEdit209179_ViewportEvent(this, event) != 0; };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackPlainTextEdit209179_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void actionEvent(QActionEvent * event) { callbackPlainTextEdit209179_ActionEvent(this, event); };
	bool close() { return callbackPlainTextEdit209179_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackPlainTextEdit209179_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackPlainTextEdit209179_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackPlainTextEdit209179_EnterEvent(this, event); };
	bool hasHeightForWidth() const { return callbackPlainTextEdit209179_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackPlainTextEdit209179_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	void hide() { callbackPlainTextEdit209179_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackPlainTextEdit209179_HideEvent(this, event); };
	void initPainter(QPainter * painter) const { callbackPlainTextEdit209179_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	void leaveEvent(QEvent * event) { callbackPlainTextEdit209179_LeaveEvent(this, event); };
	void lower() { callbackPlainTextEdit209179_Lower(this); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackPlainTextEdit209179_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void moveEvent(QMoveEvent * event) { callbackPlainTextEdit209179_MoveEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackPlainTextEdit209179_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, result) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackPlainTextEdit209179_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	void raise() { callbackPlainTextEdit209179_Raise(this); };
	void repaint() { callbackPlainTextEdit209179_Repaint(this); };
	void setDisabled(bool disable) { callbackPlainTextEdit209179_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackPlainTextEdit209179_SetEnabled(this, vbo); };
	void setFocus() { callbackPlainTextEdit209179_SetFocus2(this); };
	void setHidden(bool hidden) { callbackPlainTextEdit209179_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray* t728ae7 = new QByteArray(styleSheet.toUtf8()); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7->prepend("WHITESPACE").constData()+10), t728ae7->size()-10, t728ae7 };callbackPlainTextEdit209179_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackPlainTextEdit209179_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackPlainTextEdit209179_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray* tda39a3 = new QByteArray(vqs.toUtf8()); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3->prepend("WHITESPACE").constData()+10), tda39a3->size()-10, tda39a3 };callbackPlainTextEdit209179_SetWindowTitle(this, vqsPacked); };
	void show() { callbackPlainTextEdit209179_Show(this); };
	void showFullScreen() { callbackPlainTextEdit209179_ShowFullScreen(this); };
	void showMaximized() { callbackPlainTextEdit209179_ShowMaximized(this); };
	void showMinimized() { callbackPlainTextEdit209179_ShowMinimized(this); };
	void showNormal() { callbackPlainTextEdit209179_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackPlainTextEdit209179_TabletEvent(this, event); };
	void update() { callbackPlainTextEdit209179_Update(this); };
	void updateMicroFocus() { callbackPlainTextEdit209179_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackPlainTextEdit209179_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray* t3c6de1 = new QByteArray(title.toUtf8()); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1->prepend("WHITESPACE").constData()+10), t3c6de1->size()-10, t3c6de1 };callbackPlainTextEdit209179_WindowTitleChanged(this, titlePacked); };
	void childEvent(QChildEvent * event) { callbackPlainTextEdit209179_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackPlainTextEdit209179_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackPlainTextEdit209179_CustomEvent(this, event); };
	void deleteLater() { callbackPlainTextEdit209179_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackPlainTextEdit209179_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackPlainTextEdit209179_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackPlainTextEdit209179_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackPlainTextEdit209179_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackPlainTextEdit209179_TimerEvent(this, event); };
signals:
public slots:
	void addLine(QString v0) { QByteArray* tea1dd7 = new QByteArray(v0.toUtf8()); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7->prepend("WHITESPACE").constData()+10), tea1dd7->size()-10, tea1dd7 };callbackPlainTextEdit209179_AddLine(this, v0Packed); };
private:
};

Q_DECLARE_METATYPE(PlainTextEdit209179*)


void PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaTypes() {
}

void PlainTextEdit209179_AddLine(void* ptr, struct Moc_PackedString v0)
{
	QMetaObject::invokeMethod(static_cast<PlainTextEdit209179*>(ptr), "addLine", Q_ARG(QString, QString::fromUtf8(v0.data, v0.len)));
}

int PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType()
{
	return qRegisterMetaType<PlainTextEdit209179*>();
}

int PlainTextEdit209179_PlainTextEdit209179_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<PlainTextEdit209179*>(const_cast<const char*>(typeName));
}

int PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PlainTextEdit209179>();
#else
	return 0;
#endif
}

int PlainTextEdit209179_PlainTextEdit209179_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<PlainTextEdit209179>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int PlainTextEdit209179_PlainTextEdit209179_QmlRegisterUncreatableType(char* uri, int versionMajor, int versionMinor, char* qmlName, struct Moc_PackedString message)
{
#ifdef QT_QML_LIB
	return qmlRegisterUncreatableType<PlainTextEdit209179>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName), QString::fromUtf8(message.data, message.len));
#else
	return 0;
#endif
}

int PlainTextEdit209179_PlainTextEdit209179_QmlRegisterAnonymousType(char* uri, int versionMajor)
{
#ifdef QT_QML_LIB
	return qmlRegisterAnonymousType<PlainTextEdit209179>(const_cast<const char*>(uri), versionMajor);
#else
	return 0;
#endif
}

void* PlainTextEdit209179___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* PlainTextEdit209179___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* PlainTextEdit209179___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* PlainTextEdit209179___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* PlainTextEdit209179___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* PlainTextEdit209179___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* PlainTextEdit209179___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* PlainTextEdit209179___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* PlainTextEdit209179___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PlainTextEdit209179___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* PlainTextEdit209179___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void PlainTextEdit209179___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* PlainTextEdit209179___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* PlainTextEdit209179___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PlainTextEdit209179___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* PlainTextEdit209179___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void PlainTextEdit209179___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* PlainTextEdit209179___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* PlainTextEdit209179_NewPlainTextEdit(void* parent)
{
		return new PlainTextEdit209179(static_cast<QWidget*>(parent));
}

void* PlainTextEdit209179_NewPlainTextEdit2(struct Moc_PackedString text, void* parent)
{
		return new PlainTextEdit209179(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent));
}

void PlainTextEdit209179_DestroyPlainTextEdit(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->~PlainTextEdit209179();
}

void PlainTextEdit209179_DestroyPlainTextEditDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void PlainTextEdit209179_AppendHtmlDefault(void* ptr, struct Moc_PackedString html)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::appendHtml(QString::fromUtf8(html.data, html.len));
}

void PlainTextEdit209179_AppendPlainTextDefault(void* ptr, struct Moc_PackedString text)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::appendPlainText(QString::fromUtf8(text.data, text.len));
}

char PlainTextEdit209179_CanInsertFromMimeDataDefault(void* ptr, void* source)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::canInsertFromMimeData(static_cast<QMimeData*>(source));
}

void PlainTextEdit209179_CenterCursorDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::centerCursor();
}

void PlainTextEdit209179_ChangeEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::changeEvent(static_cast<QEvent*>(e));
}

void PlainTextEdit209179_ClearDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::clear();
}

void PlainTextEdit209179_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void PlainTextEdit209179_CopyDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::copy();
}

void* PlainTextEdit209179_CreateMimeDataFromSelectionDefault(void* ptr)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::createMimeDataFromSelection();
}

void PlainTextEdit209179_CutDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::cut();
}

void PlainTextEdit209179_DragEnterEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::dragEnterEvent(static_cast<QDragEnterEvent*>(e));
}

void PlainTextEdit209179_DragLeaveEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void PlainTextEdit209179_DragMoveEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void PlainTextEdit209179_DropEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::dropEvent(static_cast<QDropEvent*>(e));
}

void PlainTextEdit209179_FocusInEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::focusInEvent(static_cast<QFocusEvent*>(e));
}

char PlainTextEdit209179_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::focusNextPrevChild(next != 0);
}

void PlainTextEdit209179_FocusOutEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::focusOutEvent(static_cast<QFocusEvent*>(e));
}

void PlainTextEdit209179_InputMethodEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::inputMethodEvent(static_cast<QInputMethodEvent*>(e));
}

void* PlainTextEdit209179_InputMethodQueryDefault(void* ptr, long long property)
{
	return new QVariant(static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void PlainTextEdit209179_InsertFromMimeDataDefault(void* ptr, void* source)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::insertFromMimeData(static_cast<QMimeData*>(source));
}

void PlainTextEdit209179_InsertPlainTextDefault(void* ptr, struct Moc_PackedString text)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::insertPlainText(QString::fromUtf8(text.data, text.len));
}

void PlainTextEdit209179_KeyPressEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::keyPressEvent(static_cast<QKeyEvent*>(e));
}

void PlainTextEdit209179_KeyReleaseEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::keyReleaseEvent(static_cast<QKeyEvent*>(e));
}

void* PlainTextEdit209179_LoadResourceDefault(void* ptr, int ty, void* name)
{
	return new QVariant(static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::loadResource(ty, *static_cast<QUrl*>(name)));
}

void PlainTextEdit209179_MouseDoubleClickEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::mouseDoubleClickEvent(static_cast<QMouseEvent*>(e));
}

void PlainTextEdit209179_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void PlainTextEdit209179_MousePressEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::mousePressEvent(static_cast<QMouseEvent*>(e));
}

void PlainTextEdit209179_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void PlainTextEdit209179_PaintEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::paintEvent(static_cast<QPaintEvent*>(e));
}

void PlainTextEdit209179_PasteDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::paste();
}

void PlainTextEdit209179_RedoDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::redo();
}

void PlainTextEdit209179_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::resizeEvent(static_cast<QResizeEvent*>(e));
}

void PlainTextEdit209179_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::scrollContentsBy(dx, dy);
}

void PlainTextEdit209179_SelectAllDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::selectAll();
}

void PlainTextEdit209179_SetPlainTextDefault(void* ptr, struct Moc_PackedString text)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setPlainText(QString::fromUtf8(text.data, text.len));
}

void PlainTextEdit209179_ShowEventDefault(void* ptr, void* vqs)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::showEvent(static_cast<QShowEvent*>(vqs));
}

void PlainTextEdit209179_UndoDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::undo();
}

void PlainTextEdit209179_WheelEventDefault(void* ptr, void* e)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::wheelEvent(static_cast<QWheelEvent*>(e));
}

void PlainTextEdit209179_ZoomInDefault(void* ptr, int ran)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::zoomIn(ran);
}

void PlainTextEdit209179_ZoomOutDefault(void* ptr, int ran)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::zoomOut(ran);
}

char PlainTextEdit209179_EventDefault(void* ptr, void* event)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::event(static_cast<QEvent*>(event));
}

void* PlainTextEdit209179_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void PlainTextEdit209179_SetupViewportDefault(void* ptr, void* viewport)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setupViewport(static_cast<QWidget*>(viewport));
}

void* PlainTextEdit209179_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char PlainTextEdit209179_ViewportEventDefault(void* ptr, void* event)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::viewportEvent(static_cast<QEvent*>(event));
}

void* PlainTextEdit209179_ViewportSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void PlainTextEdit209179_ActionEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::actionEvent(static_cast<QActionEvent*>(event));
}

char PlainTextEdit209179_CloseDefault(void* ptr)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::close();
}

void PlainTextEdit209179_CloseEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::closeEvent(static_cast<QCloseEvent*>(event));
}

void PlainTextEdit209179_EnterEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::enterEvent(static_cast<QEvent*>(event));
}

char PlainTextEdit209179_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::hasHeightForWidth();
}

int PlainTextEdit209179_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::heightForWidth(w);
}

void PlainTextEdit209179_HideDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::hide();
}

void PlainTextEdit209179_HideEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::hideEvent(static_cast<QHideEvent*>(event));
}

void PlainTextEdit209179_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::initPainter(static_cast<QPainter*>(painter));
}

void PlainTextEdit209179_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::leaveEvent(static_cast<QEvent*>(event));
}

void PlainTextEdit209179_LowerDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::lower();
}

int PlainTextEdit209179_MetricDefault(void* ptr, long long m)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void PlainTextEdit209179_MoveEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::moveEvent(static_cast<QMoveEvent*>(event));
}

char PlainTextEdit209179_NativeEventDefault(void* ptr, void* eventType, void* message, long* result)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::nativeEvent(*static_cast<QByteArray*>(eventType), message, result);
}

void* PlainTextEdit209179_PaintEngineDefault(void* ptr)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::paintEngine();
}

void PlainTextEdit209179_RaiseDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::raise();
}

void PlainTextEdit209179_RepaintDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::repaint();
}

void PlainTextEdit209179_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setDisabled(disable != 0);
}

void PlainTextEdit209179_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setEnabled(vbo != 0);
}

void PlainTextEdit209179_SetFocus2Default(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setFocus();
}

void PlainTextEdit209179_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setHidden(hidden != 0);
}

void PlainTextEdit209179_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void PlainTextEdit209179_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setVisible(visible != 0);
}

void PlainTextEdit209179_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setWindowModified(vbo != 0);
}

void PlainTextEdit209179_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void PlainTextEdit209179_ShowDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::show();
}

void PlainTextEdit209179_ShowFullScreenDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::showFullScreen();
}

void PlainTextEdit209179_ShowMaximizedDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::showMaximized();
}

void PlainTextEdit209179_ShowMinimizedDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::showMinimized();
}

void PlainTextEdit209179_ShowNormalDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::showNormal();
}

void PlainTextEdit209179_TabletEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::tabletEvent(static_cast<QTabletEvent*>(event));
}

void PlainTextEdit209179_UpdateDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::update();
}

void PlainTextEdit209179_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::updateMicroFocus();
}

void PlainTextEdit209179_ChildEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::childEvent(static_cast<QChildEvent*>(event));
}

void PlainTextEdit209179_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void PlainTextEdit209179_CustomEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::customEvent(static_cast<QEvent*>(event));
}

void PlainTextEdit209179_DeleteLaterDefault(void* ptr)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::deleteLater();
}

void PlainTextEdit209179_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char PlainTextEdit209179_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void PlainTextEdit209179_TimerEventDefault(void* ptr, void* event)
{
	static_cast<PlainTextEdit209179*>(ptr)->QPlainTextEdit::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
