package state

var curstate map[string]interface{}

func init() {
	if curstate == nil {
		curstate = make(map[string]interface{})
	}
}

func Get(key string) (interface{}, bool) {
	v, ok := curstate[key]
	return v, ok
}

func Set(key string, val interface{}) {
	curstate[key] = val
}
