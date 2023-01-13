package cra

const (
	mainSiteConfig = "mainSiteConfig"
	mainApiConfig  = "mainApiConfig"
)

var (
	configs = map[string]string{
		mainSiteConfig: "\ts := site.NewSite(config.Http.RootPath, db, sessionManager, e)\n\ts.Serve()\n\n",
		mainApiConfig:  "\ta := api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db, sessionManager)\n\ta.Serve()\n\n",
	}
)

const (
	Session AuthType = "session"
)

type AuthType string

type CMD struct {
	Key     string
	Command []string
}

type Site struct {
	Templating bool
	SSR        bool
	Static     bool // TODO: Not yet implemented
	StaticSPA  bool // TODO: Not yet implemented
}

type Selections struct {
	API        bool
	Site       bool
	SiteConfig *Site
	Auth       AuthType
}

type Options struct {
	ProjectParentDir    string
	TemplateDir         string
	ProjectName         string
	GoProjectModuleName string
	Vendoring           bool
	Selections          Selections
}

type WebTemplate struct {
	options                Options
	baseFiles              map[string]ItemType
	siteFiles              map[string]ItemType
	apiFiles               map[string]ItemType
	mainGo                 string
	mainGoTemplate         string
	commands               []CMD
	sedTemplateNameCommand string
	configs                map[string]string
}

func NewWebTemplate(options Options) WebTemplate {
	return WebTemplate{
		options: options,
		baseFiles: map[string]ItemType{
			"dev_run.sh":         file,
			"init.go":            file,
			"config.go":          file,
			"README.md":          file,
			"docker-compose.yml": file,
			"Dockerfile":         file,
			"Makefile":           file,
			".gitignore":         file,
			".env":               file,
			"session":            dir,
			"models":             dir,
		},
		siteFiles: map[string]ItemType{
			"site": dir,
		},
		apiFiles: map[string]ItemType{
			"api": dir,
		},
		mainGo:                 "main.go",
		mainGoTemplate:         "main.go.template",
		sedTemplateNameCommand: "s/go_web_template/%s/g",
		commands: []CMD{
			{
				Key:     "git",
				Command: []string{"git", "init"},
			},
			{
				Key:     "init",
				Command: []string{"go", "mod", "init"},
			},
			{
				Key:     "tidy",
				Command: []string{"go", "mod", "tidy"},
			},
			{
				Key:     "imports",
				Command: []string{"goimports", "-w"},
			},
			{
				Key:     "get",
				Command: []string{"go", "get", "./..."},
			},
		},
		configs: map[string]string{
			mainSiteConfig: "\ts := site.NewSite(config.Http.RootPath, db, sessionManager, e)\n\ts.Serve()\n\n",
			mainApiConfig:  "\ta := api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db, sessionManager)\n\ta.Serve()\n\n",
		},
	}
}

const (
	file         ItemType = "file"
	templateFile ItemType = "templateFile"
	dir          ItemType = "dir"
)

type ItemType string
