/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package input

// IfExistsAction determines what to do if the scaffold file already exists
type IfExistsAction int

const (
	// Skip skips the file and moves to the next one
	Skip IfExistsAction = iota

	// Error returns an error and stops processing
	Error

	// Overwrite truncates and overwrites the existing file
	Overwrite
)

// Input is the input for scaffolding a file
type Input struct {
	// Path is the file to write
	Path string

	// IfExistsAction determines what to do if the file exists
	IfExistsAction IfExistsAction

	// TemplateBody is the template body to execute
	TemplateBody string

	// Boilerplate is the contents of a Boilerplate go header file
	Boilerplate string

	// BoilerplatePath is the path to a Boilerplate go header file
	BoilerplatePath string

	// Version is the project version
	Version string

	// Domain is the domain for the APIs
	Domain string

	// Repo is the go project package
	Repo string

	// ProjectPath is the relative path to the project root
	ProjectPath string

	// MultiGroup is the multi-group boolean from the PROJECT file
	MultiGroup bool
}

// Domain allows a domain to be set on an object
type Domain interface {
	// SetDomain sets the domain
	SetDomain(string)
}

// SetDomain sets the domain
func (i *Input) SetDomain(d string) {
	if i.Domain == "" {
		i.Domain = d
	}
}

// Repo allows a repo to be set on an object
type Repo interface {
	// SetRepo sets the repo
	SetRepo(string)
}

// SetRepo sets the repo
func (i *Input) SetRepo(r string) {
	if i.Repo == "" {
		i.Repo = r
	}
}

// Boilerplate allows boilerplate text to be set on an object
type Boilerplate interface {
	// SetBoilerplate sets the boilerplate text
	SetBoilerplate(string)
}

// SetBoilerplate sets the boilerplate text
func (i *Input) SetBoilerplate(b string) {
	if i.Boilerplate == "" {
		i.Boilerplate = b
	}
}

// BoilerplatePath allows boilerplate file path to be set on an object
type BoilerplatePath interface {
	// SetBoilerplatePath sets the boilerplate file path
	SetBoilerplatePath(string)
}

// SetBoilerplatePath sets the boilerplate file path
func (i *Input) SetBoilerplatePath(bp string) {
	if i.BoilerplatePath == "" {
		i.BoilerplatePath = bp
	}
}

// Version allows the project version to be set on an object
type Version interface {
	// SetVersion sets the project version
	SetVersion(string)
}

// SetVersion sets the project version
func (i *Input) SetVersion(v string) {
	if i.Version == "" {
		i.Version = v
	}
}

// MultiGroup allows the project version to be set on an object
type MultiGroup interface {
	// SetVersion sets the project version
	SetMultiGroup(value bool)
}

// SetVersion sets the MultiGroup value
func (i *Input) SetMultiGroup(v bool) {
	i.MultiGroup = v
}

// ProjecPath allows the project path to be set on an object
type ProjecPath interface {
	// SetProjectPath sets the project file location
	SetProjectPath(string)
}

// SetProjectPath sets the project path
func (i *Input) SetProjectPath(p string) {
	if i.ProjectPath == "" {
		i.ProjectPath = p
	}
}

// File is a scaffoldable file
type File interface {
	// GetInput returns the Input for creating a scaffold file
	GetInput() (Input, error)
}

// RequiresValidation is a file that requires validation
type RequiresValidation interface {
	File
	// Validate returns true if the template has valid values
	Validate() error
}

// Options are the options for executing scaffold templates
type Options struct {
	// BoilerplatePath is the path to the boilerplate file
	BoilerplatePath string

	// Path is the path to the project
	ProjectPath string
}

// ProjectFile is deserialized into a PROJECT file
type ProjectFile struct {
	// Version is the project version - defaults to "1"
	Version string `json:"version,omitempty"`

	// Domain is the domain associated with the project and used for API groups
	Domain string `json:"domain,omitempty"`

	// Repo is the go package name of the project root
	Repo string `json:"repo,omitempty"`

	// Resources tracks scaffolded resources in the project. This info is
	// tracked only in project with version 2.
	Resources []Resource `json:"resources,omitempty"`

	// Multigroup tracks if the project has more than one group
	MultiGroup bool `json:"multigroup,omitempty"`
}

// ResourceGroups returns unique groups of scaffolded resources in the project.
func (pf *ProjectFile) ResourceGroups() []string {
	groupSet := map[string]struct{}{}
	for _, r := range pf.Resources {
		groupSet[r.Group] = struct{}{}
	}

	groups := []string{}
	for g := range groupSet {
		groups = append(groups, g)
	}
	return groups
}

// Resource contains information about scaffolded resources.
type Resource struct {
	Group   string `json:"group,omitempty"`
	Version string `json:"version,omitempty"`
	Kind    string `json:"kind,omitempty"`
}
