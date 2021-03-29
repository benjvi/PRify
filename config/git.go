package config

type BranchConfig struct {
	Name string `yaml:"name"`
	Base *string `yaml:"base"`
	Rebase bool `yaml:"rebase"`
}

func (c *BranchConfig) Validate() error {
	return nil
}

type CommitConfig struct {
	Message string `yaml:"message"`
	Author *string `yaml:"author"`
	Email *string `yaml:"email"`
}

func (c *CommitConfig) Validate() error {
	return nil
}



