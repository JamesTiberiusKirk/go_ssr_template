package cra

const (
	siteConfig = "siteConfig"
	apiConfig  = "apiConfig"
)

var (
	configs = map[string]string{
		siteConfig: "\ts := site.NewSite(config.Http.RootPath, db, sessionManager, e)\n\ts.Serve()\n\n",
		apiConfig:  "\ta := api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db, sessionManager)\n\ta.Serve()\n\n",
	}
)

const (
	Session AuthType = "session"
)

type AuthType string

type Selections struct {
	API  bool
	SSR  bool
	Auth AuthType
}

type Options struct {
	ProjectParentDir    string
	TemplateDir         string
	ProjectName         string
	GoProjectModuleName string
	Vendoring           bool
	Selections          Selections
}

const (
	file ItemType = "file"
	dir  ItemType = "dir"
)

type ItemType string

var (
	baseFiles = map[string]ItemType{
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
	}

	ssrFiles = map[string]ItemType{
		"site": dir,
	}

	apiFiles = map[string]ItemType{
		"api": dir,
	}

	mainGo     = "main.go"
	mainGoTmpl = "main.go.template"

	commands = map[string]string{
		"": "",
	}
)

const sedTemplateNameCommand = "s/go_ssr_template/%s/g"

type CMD struct {
	Key     string
	Command []string
}

var (
	installCmds = []CMD{
		{
			Key:     "git",
			Command: []string{"git", "init"},
		},
		{
			Key:     "init",
			Command: []string{"go", "mod", "init"},
		},
		{
			Key:     "itdy",
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
	}
)
