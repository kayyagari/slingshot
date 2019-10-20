package jnlp
// all the structs in this file are generated using https://github.com/csimplestring/xsd-2-go
// the tool was written in Java but it was helpful in getting what I needed
import (
	"encoding/xml"
)

type InstallerDesc struct {
	XMLName   xml.Name `xml:"installer-desc"`
	MainClass string   `xml:"main-class,attr,omitempty"`
}

type Description struct {
	XMLName xml.Name `xml:"description"`
	Kind    string   `xml:"kind,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type J2EeApplicationClientPermissions struct {
	XMLName xml.Name `xml:"j2ee-application-client-permissions"`
}

type Security struct {
	XMLName                          xml.Name                          `xml:"security"`
	AllPermissions                   *AllPermissions                   `xml:"all-permissions,omitempty"`
	J2EeApplicationClientPermissions *J2EeApplicationClientPermissions `xml:"j2ee-application-client-permissions,omitempty"`
}

type Icon struct {
	XMLName xml.Name `xml:"icon"`
	Size    string   `xml:"size,attr,omitempty"`
	Kind    string   `xml:"kind,attr,omitempty"`
	Depth   string   `xml:"depth,attr,omitempty"`
	Version string   `xml:"version,attr,omitempty"`
	Height  string   `xml:"height,attr,omitempty"`
	Width   string   `xml:"width,attr,omitempty"`
	Href    string   `xml:"href,attr"`
}

type AllPermissions struct {
	XMLName xml.Name `xml:"all-permissions"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`
	Value   string   `xml:"value,attr"`
	Name    string   `xml:"name,attr"`
}

type J2Se struct {
	XMLName         xml.Name    `xml:"j2se"`
	Version         string      `xml:"version,attr"`
	InitialHeapSize string      `xml:"initial-heap-size,attr,omitempty"`
	MaxHeapSize     string      `xml:"max-heap-size,attr,omitempty"`
	JavaVmArgs      string      `xml:"java-vm-args,attr,omitempty"`
	Href            string      `xml:"href,attr,omitempty"`
	Resources       []Resources `xml:"resources,omitempty"`
}

type Update struct {
	XMLName xml.Name `xml:"update"`
	Check   string   `xml:"check,attr,omitempty"`
	Policy  string   `xml:"policy,attr,omitempty"`
}

type Shortcut struct {
	XMLName xml.Name `xml:"shortcut"`
	Online  string   `xml:"online,attr,omitempty"`
	Desktop *Desktop `xml:"desktop,omitempty"`
	Menu    *Menu    `xml:"menu,omitempty"`
}

type Jar struct {
	XMLName  xml.Name `xml:"jar"`
	Part     string   `xml:"part,attr,omitempty"`
	Size     string   `xml:"size,attr,omitempty"`
	Download string   `xml:"download,attr,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
	Href     string   `xml:"href,attr"`
	Main     string   `xml:"main,attr,omitempty"`
}

type Nativelib struct {
	XMLName  xml.Name `xml:"nativelib"`
	Size     string   `xml:"size,attr,omitempty"`
	Part     string   `xml:"part,attr,omitempty"`
	Download string   `xml:"download,attr,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
	Href     string   `xml:"href,attr"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Value   string   `xml:"value,attr"`
	Name    string   `xml:"name,attr"`
}

type RelatedContent struct {
	XMLName     xml.Name     `xml:"related-content"`
	Href        string       `xml:"href,attr"`
	Title       string       `xml:"title,omitempty"`
	Description *Description `xml:"description,omitempty"`
	Icon        *Icon        `xml:"icon,omitempty"`
}

type ComponentDesc struct {
	XMLName xml.Name `xml:"component-desc"`
}

type Association struct {
	XMLName     xml.Name     `xml:"association"`
	MimeType    string       `xml:"mime-type,attr"`
	Extensions  string       `xml:"extensions,attr"`
	Description *Description `xml:"description,omitempty"`
	Icon        *Icon        `xml:"icon,omitempty"`
}

type ApplicationDesc struct {
	XMLName   xml.Name `xml:"application-desc"`
	MainClass string   `xml:"main-class,attr,omitempty"`
	Argument  []string `xml:"argument,omitempty"`
}

type ExtDownload struct {
	XMLName  xml.Name `xml:"ext-download"`
	Download string   `xml:"download,attr,omitempty"`
	ExtPart  string   `xml:"ext-part,attr"`
	Part     string   `xml:"part,attr,omitempty"`
}

type Desktop struct {
	XMLName xml.Name `xml:"desktop"`
}

type Jnlp struct {
	XMLName         xml.Name         `xml:"jnlp"`
	Href            string           `xml:"href,attr,omitempty"`
	Spec            string           `xml:"spec,attr,omitempty"`
	Version         string           `xml:"version,attr,omitempty"`
	Codebase        string           `xml:"codebase,attr,omitempty"`
	Information     []Information    `xml:"information"`
	Update          *Update          `xml:"update,omitempty"`
	Resources       []Resources      `xml:"resources,omitempty"`
	AppletDesc      *AppletDesc      `xml:"applet-desc,omitempty"`
	InstallerDesc   *InstallerDesc   `xml:"installer-desc,omitempty"`
	Security        *Security        `xml:"security,omitempty"`
	ApplicationDesc *ApplicationDesc `xml:"application-desc,omitempty"`
	ComponentDesc   *ComponentDesc   `xml:"component-desc,omitempty"`
}

type Homepage struct {
	XMLName xml.Name `xml:"homepage"`
	Href    string   `xml:"href,attr"`
}

type Extension struct {
	XMLName     xml.Name      `xml:"extension"`
	Name        string        `xml:"name,attr,omitempty"`
	Href        string        `xml:"href,attr"`
	Version     string        `xml:"version,attr,omitempty"`
	ExtDownload []ExtDownload `xml:"ext-download,omitempty"`
}

type Package struct {
	XMLName   xml.Name `xml:"package"`
	Name      string   `xml:"name,attr"`
	Recursive string   `xml:"recursive,attr,omitempty"`
	Part      string   `xml:"part,attr"`
}

type AppletDesc struct {
	XMLName      xml.Name `xml:"applet-desc"`
	Height       string   `xml:"height,attr"`
	Name         string   `xml:"name,attr"`
	Width        string   `xml:"width,attr"`
	Documentbase string   `xml:"documentbase,attr,omitempty"`
	MainClass    string   `xml:"main-class,attr"`
	Param        []Param  `xml:"param,omitempty"`
}

type Menu struct {
	XMLName xml.Name `xml:"menu"`
	Submenu string   `xml:"submenu,attr,omitempty"`
}

type Java struct {
	XMLName         xml.Name    `xml:"java"`
	Version         string      `xml:"version,attr"`
	InitialHeapSize string      `xml:"initial-heap-size,attr,omitempty"`
	MaxHeapSize     string      `xml:"max-heap-size,attr,omitempty"`
	JavaVmArgs      string      `xml:"java-vm-args,attr,omitempty"`
	Href            string      `xml:"href,attr,omitempty"`
	Resources       []Resources `xml:"resources,omitempty"`
}

type Resources struct {
	XMLName   xml.Name    `xml:"resources"`
	Os        string      `xml:"os,attr,omitempty"`
	Locale    string      `xml:"locale,attr,omitempty"`
	Arch      string      `xml:"arch,attr,omitempty"`
	Java      []Java      `xml:"java,omitempty"`
	J2se      []J2Se      `xml:"j2se,omitempty"`
	Jar       []Jar       `xml:"jar,omitempty"`
	Nativelib []Nativelib `xml:"nativelib,omitempty"`
	Extension []Extension `xml:"extension,omitempty"`
	Property  []Property  `xml:"property,omitempty"`
	Package   []Package   `xml:"package,omitempty"`
}

type Information struct {
	XMLName        xml.Name         `xml:"information"`
	Os             string           `xml:"os,attr,omitempty"`
	Locale         string           `xml:"locale,attr,omitempty"`
	Platform       string           `xml:"platform,attr,omitempty"`
	Arch           string           `xml:"arch,attr,omitempty"`
	RelatedContent []RelatedContent `xml:"related-content,omitempty"`
	Vendor         string           `xml:"vendor"`
	OfflineAllowed *OfflineAllowed  `xml:"offline-allowed,omitempty"`
	Shortcut       *Shortcut        `xml:"shortcut,omitempty"`
	Title          string           `xml:"title"`
	Homepage       *Homepage        `xml:"homepage,omitempty"`
	Icon           []Icon           `xml:"icon,omitempty"`
	Description    []Description    `xml:"description,omitempty"`
	Association    *Association     `xml:"association,omitempty"`
}

type OfflineAllowed struct {
	XMLName xml.Name `xml:"offline-allowed"`
}
