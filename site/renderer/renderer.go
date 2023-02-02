package renderer

import (
	"bytes"
	"fmt"
	"github.com/JamesTiberiusKirk/go_web_template/site/page"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

const (
	Master   TemplateType = "usesMaster"
	NoMaster TemplateType = "noMaster"
	Include  TemplateType = "include"
)

type TemplateType string

// HTMLContentType const templateEngineKey = "httpx_templateEngine"
var HTMLContentType = []string{"text/html; charset=utf-8"}

// DefaultConfig default config
var DefaultConfig = Config{
	Root:         "views",
	Master:       "layouts/master",
	Partials:     []string{},
	Funcs:        make(template.FuncMap),
	DisableCache: false,
	Delims:       Delims{Left: "{{", Right: "}}"},
}

// ViewEngine view template engine
type ViewEngine struct {
	config      Config
	tplMap      map[string]*template.Template
	tplMutex    sync.RWMutex
	fileHandler FileHandler
}

// Config configuration options
type Config struct {
	Root         string           //view root
	Master       string           //template master
	NoFrame      string           //template master
	Partials     []string         //template partial, such as head, foot
	Funcs        template.FuncMap //template functions
	DisableCache bool             //disable cache, debug mode
	Delims       Delims           //delimeters
}

// M map interface for data
type M map[string]interface{}

// Delims delims for template
type Delims struct {
	Left  string
	Right string
}

// New new template engine
func New(config Config) *ViewEngine {
	return &ViewEngine{
		config:      config,
		tplMap:      make(map[string]*template.Template),
		tplMutex:    sync.RWMutex{},
		fileHandler: defaultFileHandler(),
	}
}

// Default new default template engine
func Default() *ViewEngine {
	return New(DefaultConfig)
}

// Render render template for echo interface
func (e *ViewEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	frame := c.Get(page.UseFrameName).(bool)
	return e.RenderWriter(w, name, data, frame)
}

// RenderWriter render template with io.Writer
func (e *ViewEngine) RenderWriter(w io.Writer, name string, data interface{}, useMaster bool) error {
	frame := Master
	if useMaster {
		frame = Master
	} else {
		frame = NoMaster
	}

	return e.executeTemplate(w, name, data, frame)
}

func (e *ViewEngine) executeTemplate(out io.Writer, name string, data interface{}, frame TemplateType) error {
	var tpl *template.Template
	var err error
	var ok bool

	allFuncs := make(template.FuncMap, 0)

	allFuncs["include"] = func(tmpl string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, tmpl, data, Include)
		return template.HTML(buf.String()), err
	}

	allFuncs["includeJs"] = func(tmpl string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, tmpl, data, Include)
		js := template.JS(buf.String())
		return template.HTML("\n<script>\n" + js + "\n</script>\n"), err
	}

	allFuncs["includeTs"] = func(tmpl string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		path := "jsdist/" + strings.Replace(tmpl, ".ts", ".js", 1)
		err := e.executeTemplate(buf, path, data, Include)
		js := template.JS(buf.String())
		return template.HTML("\n<script>\n" + js + "\n</script>\n"), err
	}

	// Get the plugin collection
	for k, v := range e.config.Funcs {
		allFuncs[k] = v
	}

	e.tplMutex.RLock()
	tpl, ok = e.tplMap[name]
	e.tplMutex.RUnlock()

	exeName := name
	if frame == Master && e.config.Master != "" {
		exeName = e.config.Master
	} else if frame == NoMaster && e.config.NoFrame != "" {
		exeName = e.config.NoFrame
	}

	if !ok || e.config.DisableCache {
		tplList := make([]string, 0)
		if frame == Master {
			if e.config.Master != "" {
				tplList = append(tplList, e.config.Master)
			}
		} else if frame == NoMaster {
			if e.config.NoFrame != "" {
				tplList = append(tplList, e.config.NoFrame)
			}
		}

		tplList = append(tplList, name)
		tplList = append(tplList, e.config.Partials...)

		// Loop through each template and test the full path
		tpl = template.New(name).Funcs(allFuncs).Delims(e.config.Delims.Left, e.config.Delims.Right)
		for _, t := range tplList {
			var data string
			data, err = e.fileHandler(e.config, t)
			if err != nil {
				return err
			}
			var tmpl *template.Template
			if t == name {
				tmpl = tpl
			} else {
				tmpl = tpl.New(t)
			}
			_, err = tmpl.Parse(data)
			if err != nil {
				return fmt.Errorf("ViewEngine render parser name:%v, error: %v", t, err)
			}
		}

		e.tplMutex.Lock()
		e.tplMap[name] = tpl
		e.tplMutex.Unlock()
	}

	// Display the content to the screen
	err = tpl.Funcs(allFuncs).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("ViewEngine execute template error: %v", err)
	}

	return nil
}

// FileHandler file handler interface
type FileHandler func(config Config, tplFile string) (content string, err error)

// defaultFileHandler new default file handler
// This is just supposed to read the file and return a string as content
func defaultFileHandler() FileHandler {
	return func(config Config, tplFile string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(config.Root + string(os.PathSeparator) + tplFile)
		if err != nil {
			return "", fmt.Errorf("ViewEngine path:%v error: %v", path, err)
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("ViewEngine render read name:%v, path:%v, error: %v", tplFile, path, err)
		}
		return string(data), nil
	}
}
