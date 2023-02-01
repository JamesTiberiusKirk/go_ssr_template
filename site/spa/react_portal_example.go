package spa

func NewReactPortal() *Site {
	return &Site{
		MenuID:  "portal",
		Path:    "/portal",
		Dist:    "site/spa/react_portal/build/",
		Index:   "index.html",
		Routing: true,
	}
}
