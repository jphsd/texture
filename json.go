package texture

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJSON(v any, name string) error {
	fDst, err := os.Create(fmt.Sprintf("%s.json", name))
	if err != nil {
		return err
	}
	defer fDst.Close()
	enc := json.NewEncoder(fDst)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
