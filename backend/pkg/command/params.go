package command

type CommandParameterInterface interface {
	GetValue(key string, defaultValue any) any
	SetValue(key string, val any)
	Validate() bool
}

type BaseParameters struct {
	Data map[string]any
}

func (b *BaseParameters) GetValue(key string, defaultValue any) any {
	if val, ok := b.Data[key]; ok {
		return val
	}
	return defaultValue
}

func (b *BaseParameters) SetValue(key string, val any) {
	b.Data[key] = val
}

func (b *BaseParameters) Validate() bool {
	return true
}