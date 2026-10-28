package registry

import "main/model"

func Success(name string) bool {
	command, ok := registry[name]
	return ok && command.Success
}

func buildRegistry(list []model.Command) map[string]model.Command {
	result := make(map[string]model.Command, len(list))
	for _, command := range list {
		if command.Name == "" {
			panic("invalid command: empty name")
		}
		if _, exists := result[command.Name]; exists {
			panic("duplicate command: " + command.Name)
		}
		result[command.Name] = command
	}
	return result
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func ignoreEvent(fn func() error) model.Handler {
	return func(_ model.Event) error {
		return fn()
	}
}
