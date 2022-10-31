package cmd

type GHCConfig struct {
	Export struct {
		Github struct {
			EnterpriseDomain string `yaml:"enterpriseDomain,omitempty" json:"enterpriseDomain,omitempty"`
			Token            string `yaml:"token,omitempty" json:"token,omitempty"`
			Organization     string `yaml:"organization" json:"organization" validate:"required"`
			Repository       string `yaml:"repository" json:"repository" validate:"required"`
		} `yaml:"github"`
	} `yaml:"export"`
}
