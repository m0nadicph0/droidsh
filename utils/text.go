package utils

import "github.com/jaswdr/faker"

func PromptGenerator(s string) []string {
	result := make([]string, 0)
	fake := faker.New()
	for i := 0; i < 10; i++ {
		result = append(result, fake.Internet().Domain())
	}
	return result
}
