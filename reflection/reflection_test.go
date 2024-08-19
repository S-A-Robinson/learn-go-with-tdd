package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Scott"},
			[]string{"Scott"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Scott", "Leeds"},
			[]string{"Scott", "Leeds"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Scott", 27},
			[]string{"Scott"},
		},
		{
			"nested fields",
			Person{
				"Scott",
				Profile{27, "Leeds"},
			},
			[]string{"Scott", "Leeds"},
		},
		{
			"pointers to things",
			&Person{
				"Scott",
				Profile{27, "Leeds"},
			},
			[]string{"Scott", "Leeds"},
		},
		{
			"slices",
			[]Profile{
				{27, "Leeds"},
				{55, "Manchester"},
			},
			[]string{"Leeds", "Manchester"},
		},
		{
			"arrays",
			[2]Profile{
				{27, "Leeds"},
				{55, "Manchester"},
			},
			[]string{"Leeds", "Manchester"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, wanted	 %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		testMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		testChannel := make(chan Profile)

		go func() {
			testChannel <- Profile{27, "Manchester"}
			testChannel <- Profile{28, "Leeds"}
			close(testChannel)
		}()

		var got []string
		want := []string{"Manchester", "Leeds"}

		Walk(testChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		testFunction := func() (Profile, Profile) {
			return Profile{27, "Manchester"}, Profile{28, "Leeds"}
		}

		var got []string
		want := []string{"Manchester", "Leeds"}

		Walk(testFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
