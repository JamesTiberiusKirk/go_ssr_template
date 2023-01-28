package spa

func NewReactPortal() *Site {
	return &Site{
		Frame: true,
		Path:  "/portal",
		Dist:  "site/spa/react_portal/build/",
		Index: "index.html",
	}
}
