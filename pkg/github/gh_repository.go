package github

import (
	"gopkg.in/yaml.v2"
)

type GHRepository struct {
	Organisation GHOrganization
	Repository   string
}

func (ghRepo GHRepository) GetConfig() (string, error) {

	client, ctx := GHBase{
		ghToken:   ghRepo.Organisation.GhToken,
		gheDomain: ghRepo.Organisation.GhEnterpriseDomain,
	}.getCient()
	repository, _, listError := client.Repositories.Get(ctx, ghRepo.Organisation.Organisation, ghRepo.Repository)
	if listError != nil {
		return "", listError
	}

	yamlMarshal, yamlMarshalErr := yaml.Marshal(repository)
	if yamlMarshalErr != nil {
		return "", yamlMarshalErr
	}

	profile := &GHRepositoryProfile{}
	yaml.Unmarshal(yamlMarshal, profile)
	profileMarshal, profileMarshalErr := yaml.Marshal(profile)
	if profileMarshalErr != nil {
		return "", profileMarshalErr
	}

	return string(profileMarshal), nil
}

type GHRepositorySettings struct {
	Profile GHRepositoryProfile `json:"profile"`
}

type GHRepositoryProfile struct {
	Name                      string        `json:"name"`
	Fullname                  string        `json:"fullname"`
	Description               string        `json:"description"`
	Homepage                  interface{}   `json:"homepage"`
	Codeofconduct             interface{}   `json:"codeofconduct"`
	Defaultbranch             string        `json:"defaultbranch"`
	Masterbranch              interface{}   `json:"masterbranch"`
	Mirrorurl                 interface{}   `json:"mirrorurl"`
	Autoinit                  interface{}   `json:"autoinit"`
	Parent                    interface{}   `json:"parent"`
	Source                    interface{}   `json:"source"`
	Templaterepository        interface{}   `json:"templaterepository"`
	Permissions               Permissions   `json:"permissions"`
	Allowrebasemerge          bool          `json:"allowrebasemerge"`
	Allowupdatebranch         bool          `json:"allowupdatebranch"`
	Allowsquashmerge          bool          `json:"allowsquashmerge"`
	Allowmergecommit          bool          `json:"allowmergecommit"`
	Allowautomerge            bool          `json:"allowautomerge"`
	Allowforking              bool          `json:"allowforking"`
	Deletebranchonmerge       bool          `json:"deletebranchonmerge"`
	Usesquashprtitleasdefault bool          `json:"usesquashprtitleasdefault"`
	Squashmergecommittitle    string        `json:"squashmergecommittitle"`
	Squashmergecommitmessage  string        `json:"squashmergecommitmessage"`
	Mergecommittitle          string        `json:"mergecommittitle"`
	Mergecommitmessage        string        `json:"mergecommitmessage"`
	Topics                    []interface{} `json:"topics"`
	Archived                  bool          `json:"archived"`
	Disabled                  bool          `json:"disabled"`
	Private                   bool          `json:"private"`
	Hasissues                 bool          `json:"hasissues"`
	Haswiki                   bool          `json:"haswiki"`
	Haspages                  bool          `json:"haspages"`
	Hasprojects               bool          `json:"hasprojects"`
	Hasdownloads              bool          `json:"hasdownloads"`
	Istemplate                bool          `json:"istemplate"`
	Licensetemplate           interface{}   `json:"licensetemplate"`
	Gitignoretemplate         interface{}   `json:"gitignoretemplate"`
	Securityandanalysis       interface{}   `json:"securityandanalysis"`
	Teamid                    interface{}   `json:"teamid"`
	Visibility                string        `json:"visibility"`
	Rolename                  interface{}   `json:"rolename"`
}
type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Pull     bool `json:"pull"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
}
