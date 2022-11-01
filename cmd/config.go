package cmd

type GHCConfig struct {
	Github struct {
		EnterpriseDomain string `yaml:"enterpriseDomain,omitempty" json:"enterpriseDomain,omitempty"`
	} `yaml:"github" json:"github"`
	Export struct {
		Github struct {
			EnterpriseDomain string `yaml:"enterpriseDomain,omitempty" json:"enterpriseDomain,omitempty"`
			Token            string `yaml:"token,omitempty" json:"token,omitempty"`
			Organization     string `yaml:"organization" json:"organization" validate:"required"`
			Repository       string `yaml:"repository" json:"repository" validate:"required"`
		} `yaml:"github" json:"github"`
	} `yaml:"export" json:"export"`
}
