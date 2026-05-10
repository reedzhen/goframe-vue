package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDw0JAOZUACwgycDLmZ6UWJJZn5ecX6eumZJdmpqQWhIawMjBFLOZMZGRj+/w/wZucAKWaFakIYZyAaiGKcEqpxBiBgGJ+Zl1kSn5Kkl5JfnqdXXJgDNv3uAZ/kIs5P/AEchSWfCkM1OvW9fH079T3P6m5gY4HZyvVpV4csAwODOIqtj0RQbVXAa2tpAdxOT5CdfEWlPIWhEFs9kKz1CLjwIOgAv5BRsJCI0WcvEx5eE1Gvz13Glz8HBYv8+eMs8septkZExtn5z2Pbv3OaY2pj/zqVcMyaNm3mrMlC8+dPn7UgYFrElDkTI5YkpmStWvng0atnEq+YTa/cvC+eMqvp1+Wtl+evvZ21wm/NTtWnP1VmZLvs2SnpunLJjhszMiZf/uTd7WFwo6Wi2fi0R+FGzYts4ulhExLWSSYWMzfZ1Mo2K4nuiVUWNXE46TvZO/jg1fN3L5+/es1YX//+UU7dqTdnHVhb8q2Ah6uMd6OurufmE75GjTfzVv8r+rZtevfz/RE7nO5KCcCCtKzX9vEfBgaGf0yIIGVgCFL0C0MOUi6UIAUH3l7O18mwlMAATwmMTCLMqEmLFSlpwcCSRhCJL6EhmwRKVcgxr4Rikj+qSYTSGLLBoISD7H8FFIOP4jUYLRkhjMUeeBAgwPDfUYKJAXtQsrKBlLAwsDBYMzIwRDKBeIAAAAD//zchHs2qAwAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
