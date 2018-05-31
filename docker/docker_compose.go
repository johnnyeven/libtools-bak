package docker

func NewDockerCompose() *DockerCompose {
	return &DockerCompose{
		Version: "2",
	}
}

type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services,omitempty"`
}

func (dc DockerCompose) AddService(name string, s *Service) *DockerCompose {
	if dc.Services == nil {
		dc.Services = map[string]Service{}
	}
	dc.Services[name] = *s
	return &dc
}
