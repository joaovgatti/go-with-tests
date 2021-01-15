package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestTalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"John"},
			[]string{"John"},
		}, {
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"John", "London"},
			[]string{"John", "London"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"John", 23},
			[]string{"John"},
		}, {
			"Nested field",
			Person{
				"John",
				Profile{23, "London"},
			},
			[]string{"John", "London"},
		}, {
			"Pointer to thing",
			&Person{
				"John",
				Profile{23, "London"},
			},
			[]string{"John", "London"},
		}, {
			"Slices",
			[]Profile{
				{33, "London"},
				{32, "Rio"},
			},
			[]string{"London", "Rio"},
		}, {
			"Arrays",
			[2]Profile{
				{33, "London"},
				{31, "Rio"},
			},
			[]string{"London", "Rio"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})


}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack{
		if x == needle{
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didnt", haystack, needle)
	}
}



