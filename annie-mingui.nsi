; Script generated by the HM NIS Edit Script Wizard.

; HM NIS Edit Wizard helper defines
!define PRODUCT_NAME "annie-mingui"
!define PRODUCT_VERSION "v0.11.0"
!define PRODUCT_PUBLISHER "Zhiming Wang & 135e2"
!define PRODUCT_WEB_SITE "https://github.com/135e2/annie-mingui"
!define PRODUCT_DIR_REGKEY "Software\Microsoft\Windows\CurrentVersion\App Paths\annie-mingui.exe"
!define PRODUCT_UNINST_KEY "Software\Microsoft\Windows\CurrentVersion\Uninstall\${PRODUCT_NAME}"
!define PRODUCT_UNINST_ROOT_KEY "HKLM"

; MUI 1.67 compatible ------
!include "MUI.nsh"

; MUI Settings
!define MUI_ABORTWARNING
!define MUI_ICON "${NSISDIR}\Contrib\Graphics\Icons\modern-install.ico"
!define MUI_UNICON "${NSISDIR}\Contrib\Graphics\Icons\modern-uninstall.ico"

; Language Selection Dialog Settings
!define MUI_LANGDLL_REGISTRY_ROOT "${PRODUCT_UNINST_ROOT_KEY}"
!define MUI_LANGDLL_REGISTRY_KEY "${PRODUCT_UNINST_KEY}"
!define MUI_LANGDLL_REGISTRY_VALUENAME "NSIS:Language"

; Welcome page
!insertmacro MUI_PAGE_WELCOME
; License page
!insertmacro MUI_PAGE_LICENSE "C:\Users\Administrator\annie-mingui\COPYING"
; Directory page
!insertmacro MUI_PAGE_DIRECTORY
; Instfiles page
!insertmacro MUI_PAGE_INSTFILES
; Finish page
!define MUI_FINISHPAGE_RUN "$INSTDIR\annie-mingui.exe"
!insertmacro MUI_PAGE_FINISH

; Uninstaller pages
!insertmacro MUI_UNPAGE_INSTFILES

; Language files
!insertmacro MUI_LANGUAGE "English"
!insertmacro MUI_LANGUAGE "SimpChinese"
!insertmacro MUI_LANGUAGE "TradChinese"

; MUI end ------

Name "${PRODUCT_NAME} ${PRODUCT_VERSION}"
OutFile "Setup.exe"
InstallDir "$PROGRAMFILES\annie-mingui"
InstallDirRegKey HKLM "${PRODUCT_DIR_REGKEY}" ""
ShowInstDetails show
ShowUnInstDetails show

Function .onInit
  !insertmacro MUI_LANGDLL_DISPLAY
FunctionEnd

Section "Main" SEC01
  SetOutPath "$INSTDIR"
  SetOverwrite try
  File "C:\Users\Administrator\annie-mingui\deploy\windows\annie-mingui.exe"
  CreateDirectory "$SMPROGRAMS\annie-mingui"
  CreateShortCut "$SMPROGRAMS\annie-mingui\annie-mingui.lnk" "$INSTDIR\annie-mingui.exe"
  CreateShortCut "$DESKTOP\annie-mingui.lnk" "$INSTDIR\annie-mingui.exe"
  SetOutPath "$INSTDIR\bearer"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\bearer\qgenericbearer.dll"
  SetOutPath "$INSTDIR"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\D3Dcompiler_47.dll"
  SetOutPath "$INSTDIR\iconengines"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\iconengines\qsvgicon.dll"
  SetOutPath "$INSTDIR\imageformats"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qgif.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qicns.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qico.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qjpeg.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qsvg.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qtga.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qtiff.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qwbmp.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\imageformats\qwebp.dll"
  SetOutPath "$INSTDIR"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\libEGL.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\libgcc_s_seh-1.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\libGLESV2.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\libstdc++-6.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\libwinpthread-1.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\opengl32sw.dll"
  SetOutPath "$INSTDIR\platforminputcontexts"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\platforminputcontexts\qtvirtualkeyboardplugin.dll"
  SetOutPath "$INSTDIR\platforms"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\platforms\qwindows.dll"
  SetOutPath "$INSTDIR\qmltooling"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_debugger.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_inspector.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_local.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_messages.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_native.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_nativedebugger.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_preview.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_profiler.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_quickprofiler.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_server.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\qmltooling\qmldbg_tcp.dll"
  SetOutPath "$INSTDIR"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Core.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Gui.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Network.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Qml.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Quick.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Svg.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\Qt5Widgets.dll"
  SetOutPath "$INSTDIR\styles"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\styles\qwindowsvistastyle.dll"
  SetOutPath "$INSTDIR\translations"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_ar.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_bg.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_ca.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_cs.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_da.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_de.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_en.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_es.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_fi.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_fr.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_gd.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_he.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_hu.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_it.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_ja.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_ko.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_lv.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_pl.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_ru.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_sk.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_uk.qm"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\translations\qt_zh_TW.qm"
  SetOutPath "$INSTDIR\virtualkeyboard"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\virtualkeyboard\qtvirtualkeyboard_hangul.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\virtualkeyboard\qtvirtualkeyboard_openwnn.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\virtualkeyboard\qtvirtualkeyboard_pinyin.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\virtualkeyboard\qtvirtualkeyboard_tcime.dll"
  File "C:\Users\Administrator\annie-mingui\deploy\windows\virtualkeyboard\qtvirtualkeyboard_thai.dll"
SectionEnd

Section -AdditionalIcons
  SetOutPath $INSTDIR
  WriteIniStr "$INSTDIR\${PRODUCT_NAME}.url" "InternetShortcut" "URL" "${PRODUCT_WEB_SITE}"
  CreateShortCut "$SMPROGRAMS\annie-mingui\Website.lnk" "$INSTDIR\${PRODUCT_NAME}.url"
  CreateShortCut "$SMPROGRAMS\annie-mingui\Uninstall.lnk" "$INSTDIR\uninst.exe"
SectionEnd

Section -Post
  WriteUninstaller "$INSTDIR\uninst.exe"
  WriteRegStr HKLM "${PRODUCT_DIR_REGKEY}" "" "$INSTDIR\annie-mingui.exe"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "DisplayName" "$(^Name)"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "UninstallString" "$INSTDIR\uninst.exe"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "DisplayIcon" "$INSTDIR\annie-mingui.exe"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "DisplayVersion" "${PRODUCT_VERSION}"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "URLInfoAbout" "${PRODUCT_WEB_SITE}"
  WriteRegStr ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}" "Publisher" "${PRODUCT_PUBLISHER}"
SectionEnd


Function un.onUninstSuccess
  HideWindow
  MessageBox MB_ICONINFORMATION|MB_OK "$(^Name) 已成功地从你的计算机移除。"
FunctionEnd

Function un.onInit
!insertmacro MUI_UNGETLANGUAGE
  MessageBox MB_ICONQUESTION|MB_YESNO|MB_DEFBUTTON2 "你确实要完全移除 $(^Name) ，其及所有的组件？" IDYES +2
  Abort
FunctionEnd

Section Uninstall
  Delete "$INSTDIR\${PRODUCT_NAME}.url"
  Delete "$INSTDIR\uninst.exe"
  Delete "$INSTDIR\virtualkeyboard\qtvirtualkeyboard_thai.dll"
  Delete "$INSTDIR\virtualkeyboard\qtvirtualkeyboard_tcime.dll"
  Delete "$INSTDIR\virtualkeyboard\qtvirtualkeyboard_pinyin.dll"
  Delete "$INSTDIR\virtualkeyboard\qtvirtualkeyboard_openwnn.dll"
  Delete "$INSTDIR\virtualkeyboard\qtvirtualkeyboard_hangul.dll"
  Delete "$INSTDIR\translations\qt_zh_TW.qm"
  Delete "$INSTDIR\translations\qt_uk.qm"
  Delete "$INSTDIR\translations\qt_sk.qm"
  Delete "$INSTDIR\translations\qt_ru.qm"
  Delete "$INSTDIR\translations\qt_pl.qm"
  Delete "$INSTDIR\translations\qt_lv.qm"
  Delete "$INSTDIR\translations\qt_ko.qm"
  Delete "$INSTDIR\translations\qt_ja.qm"
  Delete "$INSTDIR\translations\qt_it.qm"
  Delete "$INSTDIR\translations\qt_hu.qm"
  Delete "$INSTDIR\translations\qt_he.qm"
  Delete "$INSTDIR\translations\qt_gd.qm"
  Delete "$INSTDIR\translations\qt_fr.qm"
  Delete "$INSTDIR\translations\qt_fi.qm"
  Delete "$INSTDIR\translations\qt_es.qm"
  Delete "$INSTDIR\translations\qt_en.qm"
  Delete "$INSTDIR\translations\qt_de.qm"
  Delete "$INSTDIR\translations\qt_da.qm"
  Delete "$INSTDIR\translations\qt_cs.qm"
  Delete "$INSTDIR\translations\qt_ca.qm"
  Delete "$INSTDIR\translations\qt_bg.qm"
  Delete "$INSTDIR\translations\qt_ar.qm"
  Delete "$INSTDIR\styles\qwindowsvistastyle.dll"
  Delete "$INSTDIR\Qt5Widgets.dll"
  Delete "$INSTDIR\Qt5Svg.dll"
  Delete "$INSTDIR\Qt5Quick.dll"
  Delete "$INSTDIR\Qt5Qml.dll"
  Delete "$INSTDIR\Qt5Network.dll"
  Delete "$INSTDIR\Qt5Gui.dll"
  Delete "$INSTDIR\Qt5Core.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_tcp.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_server.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_quickprofiler.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_profiler.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_preview.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_nativedebugger.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_native.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_messages.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_local.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_inspector.dll"
  Delete "$INSTDIR\qmltooling\qmldbg_debugger.dll"
  Delete "$INSTDIR\platforms\qwindows.dll"
  Delete "$INSTDIR\platforminputcontexts\qtvirtualkeyboardplugin.dll"
  Delete "$INSTDIR\opengl32sw.dll"
  Delete "$INSTDIR\libwinpthread-1.dll"
  Delete "$INSTDIR\libstdc++-6.dll"
  Delete "$INSTDIR\libGLESV2.dll"
  Delete "$INSTDIR\libgcc_s_seh-1.dll"
  Delete "$INSTDIR\libEGL.dll"
  Delete "$INSTDIR\imageformats\qwebp.dll"
  Delete "$INSTDIR\imageformats\qwbmp.dll"
  Delete "$INSTDIR\imageformats\qtiff.dll"
  Delete "$INSTDIR\imageformats\qtga.dll"
  Delete "$INSTDIR\imageformats\qsvg.dll"
  Delete "$INSTDIR\imageformats\qjpeg.dll"
  Delete "$INSTDIR\imageformats\qico.dll"
  Delete "$INSTDIR\imageformats\qicns.dll"
  Delete "$INSTDIR\imageformats\qgif.dll"
  Delete "$INSTDIR\iconengines\qsvgicon.dll"
  Delete "$INSTDIR\D3Dcompiler_47.dll"
  Delete "$INSTDIR\bearer\qgenericbearer.dll"
  Delete "$INSTDIR\annie-mingui.exe"

  Delete "$SMPROGRAMS\annie-mingui\Uninstall.lnk"
  Delete "$SMPROGRAMS\annie-mingui\Website.lnk"
  Delete "$DESKTOP\annie-mingui.lnk"
  Delete "$SMPROGRAMS\annie-mingui\annie-mingui.lnk"

  RMDir "$SMPROGRAMS\annie-mingui"
  RMDir "$INSTDIR\virtualkeyboard"
  RMDir "$INSTDIR\translations"
  RMDir "$INSTDIR\styles"
  RMDir "$INSTDIR\qmltooling"
  RMDir "$INSTDIR\platforms"
  RMDir "$INSTDIR\platforminputcontexts"
  RMDir "$INSTDIR\imageformats"
  RMDir "$INSTDIR\iconengines"
  RMDir "$INSTDIR\bearer"
  RMDir "$INSTDIR"

  DeleteRegKey ${PRODUCT_UNINST_ROOT_KEY} "${PRODUCT_UNINST_KEY}"
  DeleteRegKey HKLM "${PRODUCT_DIR_REGKEY}"
  SetAutoClose true
SectionEnd