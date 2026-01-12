package snippets

import (
	"fmt"
	"strings"
)

func ensureMaxOpenAPI30(version string) error {
	var major, minor int

	version = strings.TrimSpace(version)
	version = strings.TrimPrefix(version, "v")

	if _, err := fmt.Sscanf(version, "%d.%d", &major, &minor); err != nil {
		return fmt.Errorf("invalid openapi version %q", version)
	}

	if major != 3 || minor > 0 {
		return fmt.Errorf("max openapi version (max 3.0.x): got %s", version)
	}
	return nil
}

func EnsureMaxVersionInString() {
	inputs := []string{
		"3.0",
		"3.0.0",
		"3.0.7",
		"3.1.0",
		"4.0.0",
		"2.0",
		" 3.0 ",
		"3",
		"3.a",
		"v3.0.0",
		"3.0.0-rc1",
		"",
	}

	for _, in := range inputs {
		err := ensureMaxOpenAPI30(in)
		if err != nil {
			fmt.Printf("%-12q -> ERROR: %v\n", in, err)
			continue
		}
		fmt.Printf("%-12q -> OK\n", in)
	}
}
