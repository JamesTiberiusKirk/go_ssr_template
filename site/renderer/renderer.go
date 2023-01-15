package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

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

// Render render template with http.ResponseWriter
func (e *ViewEngine) Render(w http.ResponseWriter, statusCode int, name string, data interface{}, useMaster bool) error {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = HTMLContentType
	}
	w.WriteHeader(statusCode)
	return e.executeRender(w, name, data, useMaster)
}

// RenderWriter render template with io.Writer
func (e *ViewEngine) RenderWriter(w io.Writer, name string, data interface{}, useMaster bool) error {
	return e.executeRender(w, name, data, useMaster)
}

func (e *ViewEngine) executeRender(out io.Writer, name string, data interface{}, useMaster bool) error {
	// useMaster := true
	return e.executeTemplate(out, name, data, useMaster)
}

func (e *ViewEngine) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	var tpl *template.Template
	var err error
	var ok bool

	allFuncs := make(template.FuncMap, 0)

	allFuncs["include"] = func(tmpl string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, tmpl, data, false)
		return template.HTML(buf.String()), err
	}

	allFuncs["includeJs"] = func(tmpl string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := e.executeTemplate(buf, tmpl, data, false)
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
	if useMaster && e.config.Master != "" {
		exeName = e.config.Master
	}

	if !ok || e.config.DisableCache {
		tplList := make([]string, 0)
		if useMaster {
			//render()
			if e.config.Master != "" {
				tplList = append(tplList, e.config.Master)
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
