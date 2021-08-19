package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(got, want, t)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		want := ErrNotFound

		assertErrorExists(err, want, t)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(err, nil, t)

		assertDefinition(dictionary, word, definition, t)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(err, ErrWordExists, t)
		assertDefinition(dictionary, word, definition, t)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("updates existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(dictionary, word, newDefinition, t)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(err, ErrWordDoesNotExist, t)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deletes existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("expected %q to be deleted", word)
		}
	})

	t.Run("delete a word that does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		assertError(err, ErrCannotDeleteNonExistantWord, t)
	})
}

func assertDefinition(dictionary Dictionary, word, definition string, t testing.TB) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(got, want error, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrorExists(got, want error, t testing.TB) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertStrings(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
