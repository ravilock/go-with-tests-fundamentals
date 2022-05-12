package Dictionary

import "testing"

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if got != definition {
		t.Errorf("got %q want %q given", got, definition)
	}
	assertNoError(t, err)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Got an error but did'nt want one")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")

		if err == nil {
			t.Fatal("Expected to get an error.")
		}

		assertError(t, err, ErrorWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just another test"
		err := dictionary.Add(word, definition)

		wantedDefinition := "this is just a test"

		assertDefinition(t, dictionary, word, wantedDefinition)
		assertError(t, err, ErrorWordAlreadyDefined)
	})
}

func TestUpdate(t *testing.T) {

}
