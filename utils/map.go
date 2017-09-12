package utils

// Map map[string]interface{} 的快捷写法
type Map map[string]interface{}

// NewMap 创建一个 Map
func NewMap(args ...interface{}) Map {
	m := Map{}
	for i := 0; i < len(args); i = i + 2 {
		m[args[i].(string)] = args[i+1]
	}
	return m
}

// Set 设置一个值，并返回自己
func (m Map) Set(key string, value interface{}) Map {
	m[key] = value
	return m
}
