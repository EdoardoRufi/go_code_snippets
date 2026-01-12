package snippets

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func LaunchCheckDuplicates() {
	devs := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Elliot"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
		{Name: "Graham"},
	}

	uniqueNames := FilterUniqueV1(devs)
	fmt.Println("devs:", uniqueNames)
}

func FilterUniqueV1(developers []Developer) []string {
	seen := make(map[string]struct{}, len(developers))
	uniqueNames := make([]string, 0, len(developers))

	for _, dev := range developers {
		if _, exists := seen[dev.Name]; exists {
			continue
		}
		seen[dev.Name] = struct{}{}
		uniqueNames = append(uniqueNames, dev.Name)
	}

	return uniqueNames
}
