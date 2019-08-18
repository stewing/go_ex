package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("missing")
		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")

	t.Run("duplicate word", func(t *testing.T) {
		err := dictionary.Add("test", "duplicate")
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is the definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is the new definition"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "This is a test."
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete existing", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("expected not to get an error.")
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected to get an error.")
	}
	if err != want {
		t.Fatal("got the wrong error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
