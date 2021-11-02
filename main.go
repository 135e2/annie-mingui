package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"

	"golang.org/x/text/language"

	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/golang/glog"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"

	"github.com/iawia002/annie/downloader"
	"github.com/iawia002/annie/extractors"
	"github.com/iawia002/annie/extractors/types"

	"github.com/fanaticscripter/annie-mingui/util"
)

const appName = "annie-mingui"
const appAuthor = "Zhiming Wang"
const appAuthorDomain = "zhimingwang.org"
const about = `<p><strong>annie-mingui</strong> v0.11.0</p>
<p>Copyright (c) 2020 Zhiming Wang</p>
<p>Copyright (c) 2021 135e2 <135e2@135e2.tk></p>
<p>annie-mingui is a Qt wrapper for <a href="https://github.com/iawia002/annie">iawia002/annie</a> the video downloader. Credits:</p>
<ul>
  <li>iawia002/annie, MIT license;</li>
  <li><a href="https://qt.io">Qt</a>, LGPLv3 license;</li>
  <li><a href="https://github.com/therecipe/qt">therecipe/qt</a>, Qt bindings for golang, LGPLv3 license.</li>
</ul>
<p>Project URL(old): <a href="https://github.com/fanaticscripter/annie-mingui">github.com/fanaticscripter/annie-mingui</a>.</p>
<p>Project URL(revived): <a href="https://github.com/135e2/annie-mingui">github.com/135e2/annie-mingui</a>.</p>
`

// GUIConfig carries all application configurations.
type GUIConfig struct {
	DestinationFolder string
	PlaylistEnabled   bool
}

// GetDestinationFolder attempts to retrieve (if not yet) and returns the
// configured destination folder.
func (c *GUIConfig) GetDestinationFolder() string {
	fallback := util.DefaultDownloadsFolder()
	val := settings.Value("destinationFolder", core.NewQVariant12(fallback)).ToString()
	c.DestinationFolder = val
	return val
}

// GetPlaylistEnabled attempts to retrieve (if not yet) and returns the
// configured value of whether playlist download is enabled.
func (c *GUIConfig) GetPlaylistEnabled() bool {
	fallback := true
	val := settings.Value("playlistEnabled", core.NewQVariant9(fallback)).ToBool()
	c.PlaylistEnabled = val
	return val
}

// SetDestinationFolder sets and persists the configured destination folder.
func (c *GUIConfig) SetDestinationFolder(val string) {
	c.DestinationFolder = val
	go func() { settings.SetValue("destinationFolder", core.NewQVariant12(val)) }()
}

// SetPlaylistEnabled sets and persists the configured value of whether playlist
// download is enabled.
func (c *GUIConfig) SetPlaylistEnabled(val bool) {
	c.PlaylistEnabled = val
	go func() { settings.SetValue("playlistEnabled", core.NewQVariant9(val)) }()
}

// LoadGUIConfig loads config values from persisted settings if possible, or
// loads defaults otherwise.
func LoadGUIConfig() *GUIConfig {
	c := &GUIConfig{}
	c.GetDestinationFolder()
	c.GetPlaylistEnabled()
	return c
}

var settings *core.QSettings
var guiConfig *GUIConfig

var locales = map[string]map[string]string{
	"zh-Hans": {
		"About":               "关于",
		"Application":         "应用",
		"Awaiting user input": "等待用户输入",
		"Destination folder":  "目标文件夹",
		"Download":            "下载",
		"Download playlists":  "下载完整播单",
		"Download started":    "下载已开始",
		"Download completed":  "下载已完成",
		"On network errors, e.g. HTTP 403, please retry a few times.": "如遇HTTP 403等网络错误，请重试几次。",
		"Pick another folder": "选择文件夹",
		"Video URL":           "视频链接",
	},
}
var locale map[string]string

// Poor man's gettext.
func tr(s string) string {
	t, ok := locale[s]
	if ok {
		return t
	}
	return s
}

// ----- START ANNIE CODE -----
//
// The following code is slightly modified from
//
//   https://github.com/iawia002/annie/blob/master/main.go

func download(c *cli.Context, videoURL string) error {
	data, err := extractors.Extract(videoURL, types.Options{
		Playlist:         guiConfig.PlaylistEnabled,
		Items:            c.String("items"),
		ItemStart:        int(c.Uint("start")),
		ItemEnd:          int(c.Uint("end")),
		ThreadNumber:     int(c.Uint("thread")),
		EpisodeTitleOnly: c.Bool("episode-title-only"),
		Cookie:           c.String("cookie"),
		YoukuCcode:       c.String("youku-ccode"),
		YoukuCkey:        c.String("youku-ckey"),
		YoukuPassword:    c.String("youku-password"),
	})
	if err != nil {
		// if this error occurs, it means that an error occurred before actually starting to extract data
		// (there is an error in the preparation step), and the data list is empty.
		return err
	}

	if c.Bool("json") {
		jsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", jsonData)
		return nil
	}

	defaultDownloader := downloader.New(downloader.Options{
		InfoOnly:   c.Bool("info"),
		Stream:     c.String("stream-format"),
		Refer:      c.String("refer"),
		OutputPath: guiConfig.DestinationFolder,
		// TODO: Avoid the dirty approach (looks so stupid now)
		OutputName:     c.String("output-name"),
		FileNameLength: int(c.Uint("file-name-length")),
		Caption:        c.Bool("caption"),
		MultiThread:    c.Bool("multi-thread"),
		ThreadNumber:   int(c.Uint("thread")),
		RetryTimes:     int(c.Uint("retry")),
		ChunkSizeMB:    int(c.Uint("chunk-size")),
		UseAria2RPC:    c.Bool("aria2"),
		Aria2Token:     c.String("aria2-token"),
		Aria2Method:    c.String("aria2-method"),
		Aria2Addr:      c.String("aria2-addr"),
	})
	errors := make([]error, 0)
	for _, item := range data {
		if item.Err != nil {
			// if this error occurs, the preparation step is normal, but the data extraction is wrong.
			// the data is an empty struct.
			errors = append(errors, item.Err)
			continue
		}
		if err = defaultDownloader.Download(item); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) != 0 {
		return errors[0]
	}
	return nil
}

// ----- END ANNIE CODE -----

// PlainTextEdit extends QPlainTextEdit.
type PlainTextEdit struct {
	widgets.QPlainTextEdit

	_ func() `constructor:"init"`

	_ func(string) `slot:"addLine"`

	carriageReturnInAction bool
}

func (p *PlainTextEdit) init() {
	p.carriageReturnInAction = false
	p.ConnectAddLine(p.addLine)
}

// addLines adds a line that might end in LF, CRLF, CR, or none of the above (in
// which case an LF is appended). Aware of CR and scroll position.
func (p *PlainTextEdit) addLine(line string) {
	scrollBar := p.VerticalScrollBar()
	currentScroll := scrollBar.Value()
	userScrolledBack := currentScroll != scrollBar.Maximum()

	cursor := p.TextCursor()
	cursor.MovePosition(gui.QTextCursor__End, gui.QTextCursor__MoveAnchor, 1)
	if p.carriageReturnInAction {
		// Remove last line.
		cursor.Select(gui.QTextCursor__LineUnderCursor)
		cursor.RemoveSelectedText()
	}
	p.carriageReturnInAction = false
	if len(line) > 0 {
		switch lastCh := line[len(line)-1]; lastCh {
		case '\n':
			cursor.InsertText(line)
		case '\r':
			cursor.InsertText(line[:len(line)-1])
			p.carriageReturnInAction = true
		default:
			cursor.InsertText(line + "\n")
		}
	} else {
		cursor.InsertText("\n")
	}

	if userScrolledBack {
		scrollBar.SetValue(currentScroll)
	} else {
		scrollBar.SetValue(scrollBar.Maximum())
	}
}

type outputBuffer struct {
	textEdit *PlainTextEdit
	reader   *bufio.Reader
	scanner  *bufio.Scanner
}

func newOutputBuffer(textEdit *PlainTextEdit) *outputBuffer {
	return &outputBuffer{
		textEdit: textEdit,
		reader:   nil,
		scanner:  nil,
	}
}

func (b *outputBuffer) attachReader(r io.Reader) {
	b.reader = bufio.NewReaderSize(r, 64*1024)
	b.scanner = bufio.NewScanner(b.reader)
	re := regexp.MustCompile(`^[^\r\n]*(\r\n|\r|\n)`)
	b.scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		token = re.Find(data)
		if token != nil {
			return len(token), token, nil
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	})
}

func (b *outputBuffer) addLine(line string) {
	b.textEdit.AddLine(line)
}

func (b *outputBuffer) readLineAndUpdate() (fullLine string, err error) {
	if !b.scanner.Scan() {
		err = b.scanner.Err()
		if err != nil {
			return "", err
		}
		err = io.EOF
	}
	fullLine = b.scanner.Text()
	if len(fullLine) > 0 {
		b.addLine(fullLine)
	}
	return
}

func init() {
	flag.Parse()

	userLang, err := jibber_jabber.DetectIETF()
	if err != nil {
		userLang = "en"
	}
	availableLangs := []string{"en"}
	availableTags := []language.Tag{language.English}
	for l := range locales {
		availableLangs = append(availableLangs, l)
		availableTags = append(availableTags, language.Make(l))
	}
	matcher := language.NewMatcher(availableTags)
	_, index, _ := matcher.Match(language.Make(userLang))
	locale = locales[availableLangs[index]]

	core.QCoreApplication_SetOrganizationName(appAuthor)
	core.QCoreApplication_SetOrganizationDomain(appAuthorDomain)

	settings = core.NewQSettings5(nil)
	guiConfig = LoadGUIConfig()
}

func initAboutWindow() *widgets.QDialog {
	w := widgets.NewQDialog(nil, 0)
	label := widgets.NewQLabel2(about, nil, 0)
	label.SetOpenExternalLinks(true)
	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(label, 0, 0)
	w.SetLayout(layout)
	return w
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(appName)
	window.ConnectCloseEvent(func(*gui.QCloseEvent) {
		settings.SetValue("_geometry", core.NewQVariant13(window.SaveGeometry()))
		settings.SetValue("_windowState", core.NewQVariant13(window.SaveState(0)))
	})

	window.RestoreGeometry(settings.Value("_geometry", core.NewQVariant()).ToByteArray())
	window.RestoreState(settings.Value("_windowState", core.NewQVariant()).ToByteArray(), 0)

	centralWidget := widgets.NewQWidget(window, 0)
	window.SetCentralWidget(centralWidget)

	menuBar := window.MenuBar()
	applicationMenu := menuBar.AddMenu2(tr("Application"))
	aboutWindow := initAboutWindow()
	aboutAction := applicationMenu.AddAction(tr("About"))
	aboutAction.SetMenuRole(widgets.QAction__AboutRole)
	aboutAction.ConnectTriggered(func(bool) {
		aboutWindow.Show()
		aboutWindow.Raise()
	})

	urlLineEdit := widgets.NewQLineEdit(nil)

	folderLineEdit := widgets.NewQLineEdit2(guiConfig.DestinationFolder, nil)
	folderLineEdit.SetReadOnly(true)
	folderLineEdit.SetMinimumWidth(250)
	folderButton := widgets.NewQPushButton2(tr("Pick another folder"), nil)
	folderDialog := widgets.NewQFileDialog2(nil, tr("Destination folder"), guiConfig.DestinationFolder, "")
	folderDialog.SetFileMode(widgets.QFileDialog__DirectoryOnly)
	folderButton.ConnectClicked(func(bool) {
		if folderDialog.Exec() != int(widgets.QDialog__Accepted) {
			return
		}
		destinationFolder := core.QDir_ToNativeSeparators(folderDialog.SelectedFiles()[0])
		folderLineEdit.SetText(destinationFolder)
		guiConfig.SetDestinationFolder(destinationFolder)
	})
	folderHBoxLayout := widgets.NewQHBoxLayout()
	folderHBoxLayout.AddWidget(folderLineEdit, 1, 0)
	folderHBoxLayout.AddWidget(folderButton, 0, 0)

	playlistCheckBox := widgets.NewQCheckBox(nil)
	playlistCheckBox.SetChecked(guiConfig.PlaylistEnabled)
	playlistCheckBox.ConnectToggled(func(checked bool) {
		guiConfig.SetPlaylistEnabled(checked)
	})

	inputFormLayout := widgets.NewQFormLayout(nil)
	inputFormLayout.SetFieldGrowthPolicy(widgets.QFormLayout__AllNonFixedFieldsGrow)
	inputFormLayout.AddRow3(tr("Video URL"), urlLineEdit)
	inputFormLayout.AddRow4(tr("Destination folder"), folderHBoxLayout)
	inputFormLayout.AddRow3(tr("Download playlists"), playlistCheckBox)

	outputTextEdit := NewPlainTextEdit(nil)
	outputTextEdit.SetReadOnly(true)
	outputTextEdit.SetMinimumHeight(400)
	outputTextEdit.SetMinimumWidth(1000)
	outputTextEdit.SetLineWrapMode(widgets.QPlainTextEdit__NoWrap)
	monospaceFont := gui.NewQFont2("Courier", -1, -1, false)
	monospaceFont.SetStyleHint(gui.QFont__Monospace, 0)
	outputTextEdit.SetFont(monospaceFont)
	output := newOutputBuffer(outputTextEdit)
	output.addLine(tr("Awaiting user input"))

	downloadButton := widgets.NewQPushButton2(tr("Download"), nil)
	downloadButton.ConnectClicked(func(bool) {
		url := strings.TrimSpace(urlLineEdit.Text())
		if len(url) > 0 {
			output.addLine(time.Now().Format("15:04:05 ") + tr("Download started"))

			// config.OutputPath = guiConfig.DestinationFolder
			// config.Playlist = guiConfig.PlaylistEnabled
			glog.Infof("downloading %s => %s", url, guiConfig.DestinationFolder)

			savedStdout := os.Stdout
			r, w, _ := os.Pipe()
			output.attachReader(r)
			os.Stdout = w

			go func() {
				for {
					_, err := output.readLineAndUpdate()
					if err != nil {
						if err == io.EOF {
							break
						}
						glog.Fatal(err)
					}
					// fmt.Fprint(savedStdout, line)
				}
				output.addLine(tr("Download completed"))
			}()

			go func() {
				if download(nil, url) != nil {
					fmt.Println(tr("On network errors, e.g. HTTP 403, please retry a few times."))
				}
				w.Close()
				os.Stdout = savedStdout
			}()
		}
	})

	layout := widgets.NewQVBoxLayout()
	layout.AddLayout(inputFormLayout, 0)
	layout.AddWidget(downloadButton, 0, 0)
	layout.AddWidget(outputTextEdit, 1, 0)
	centralWidget.SetLayout(layout)

	window.Show()

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGSEGV, syscall.SIGABRT)
	go func() {
		for {
			sig := <-sigs
			glog.Error(sig)
		}
	}()

	widgets.QApplication_Exec()
}
