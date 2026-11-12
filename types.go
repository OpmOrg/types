package types

// Package represents an application package stored in the container.
// It contains the package name and its unique identifier.
//
// The XML tags define how the fields are mapped when encoding
// or decoding XML data.
type Package struct {
	Name string `xml:"name"` // Name of the package.
	ID   string `xml:"id"`   // Unique package identifier, e.g. com.example.myapp.
}

// Status represents the current state of a container.
//
// Using a custom integer type allows the application to define
// its own set of status values.
type Status int

// Name represents an XML element name.
//
// This structure was copied from the standard encoding/xml library.
// Space stores the XML namespace, while Local stores the element name.
type Name struct {
	Space, Local string
}

// PackageList represents a collection of packages.
//
// It is used to check whether a package exists and to make
// the available packages publicly accessible.
//
// The XMLName field identifies the root XML element,
// while Packages stores the individual package entries.
type PackageList struct {
	XMLName  Name      `xml:"packages"` // Root XML element containing the package list.
	Packages []Package `xml:"pkg"`      // Collection of packages available in the container.
	img      *Image    // Reference to the Image associated with this package list.
}

// Config represents the configuration of the server.
//
// Its fields are serialized and deserialized using JSON.
// The JSON tags define the corresponding property names
// in the configuration file.
type Config struct {
	Port  string `json:"port"`  // Port used by the server.
	Limit int64  `json:"limit"` // Maximum allowed package storage size, in MB.
	Name  string `json:"name"`  // Name used to identify and separate containers.
	ATS   bool   `json:"ats"`   // Indicates whether packages signed with a test signature are accepted.
}

// Image represents the data and resources associated with a container image.
//
// It groups together the storage volume, package information,
// server configuration and operation hash.
type Image struct {
	Volume      string      // Path to the volume where packages are stored.
	PackagesXML PackageList // XML representation of the packages available in this image.
	ConfigJSON  Config      // Server configuration loaded from or represented as JSON.
	OpHash      string      // Operation hash, stored using Argon2id password hashing.
}

// Container represents an individual container managed by the application.
//
// It holds the container's internal name, associated image
// and current execution status.
type Container struct {
	name   string // Internal name used to identify the container.
	img    Image  // Image associated with this container.
	status Status // Current status of the container.
}
