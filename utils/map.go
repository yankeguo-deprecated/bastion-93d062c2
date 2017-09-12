package utils

// WildMap map[string]interface{} 的快捷写法
type WildMap map[string]interface{}

// Map 创建一个 WildMap
func Map(args ...interface{}) WildMap {
	m := WildMap{}
	for i := 0; i < len(args); i = i + 2 {
		m[args[i].(string)] = args[i+1]
	}
	return m
}

// Set 设置一个值，并返回自己
func (m WildMap) Set(key string, value interface{}) WildMap {
	m[key] = value
	return m
}
