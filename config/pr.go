package config

type PRConfig struct {
	Github *GithubConfig
	Title string
	Description string
}

func (c *PRConfig) Validate() error {
	return nil
}

type GithubConfig struct {
	Org string
	Repo string
}

func (c *GithubConfig) Validate() error {
	return nil
}
