package php_serialize

import (
    "testing"
)

func TestMap(t *testing.T) {
    v := make(map[string]interface{})
    v["width"] = 300
    v["file"] = "file.jpg"

    w, err := Encode(v)
    if err != nil {
        t.Errorf("Error PHP Serialize Encode generate, %q", err)
    }
    if w != `a:2:{s:5:"width";i:300;s:4:"file";s:8:"file.jpg";}` && w != `a:2:{s:4:"file";s:8:"file.jpg";s:5:"width";i:300;}` {
        t.Errorf("Error PHP Serizliae Encode string, %q", w)
    }
}
