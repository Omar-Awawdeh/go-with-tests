package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Omar"},
			[]string{"Omar"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Omar", "Irbid"},
			[]string{"Omar", "Irbid"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Omar", 21},
			[]string{"Omar"},
		},
		{
			"nested fields",
			Person{"Omar", Profile{21, "Irbid"}},
			[]string{"Omar", "Irbid"},
		},
		{
			"pointer to things",
			&Person{"Omar", Profile{21, "Irbid"}},
			[]string{"Omar", "Irbid"},
		},
		{
			"slices",
			[]Profile{
				{21, "Irbid"},
				{30, "Amman"},
			},
			[]string{"Irbid", "Amman"},
		},
		{
			"arrays",
			[2]Profile{
				{21, "Irbid"},
				{30, "Amman"},
			},
			[]string{"Irbid", "Amman"},
		},
		{
			"maps",
			map[string]string{
				"Name": "Omar",
				"Age":  "21",
			},
			[]string{"Omar", "21"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "MOO",
			"Sheep": "BAAA",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "MOO")
		assertContains(t, got, "BAAA")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{31, "Berlin"}
			aChannel <- Profile{21, "Irbid"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Irbid"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with a function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{21, "Berlin"}, Profile{30, "Irbid"}
		}

		var got []string
		want := []string{"Berlin", "Irbid"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, hayStack []string, needle string) {
	t.Helper()

	contains := false

	for _, val := range hayStack {
		if val == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", hayStack, needle)
	}
}
