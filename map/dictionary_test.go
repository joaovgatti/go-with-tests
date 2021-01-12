package _map

import "testing"


func TestSearch(t *testing.T){
	dictionary := Dictionary{"test":"this is just a test"}

	t.Run("known word", func(t *testing.T){
		got, _ := dictionary.Search("test")
		want  := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func (t *testing.T){
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T){
	dictionary := Dictionary{}

	t.Run("new word", func (t *testing.T){
		key := "test"
		value := "this is just a test"

		dictionary.Add(key,value)
		assertDefinition(t,dictionary,key,value)
	})

	t.Run("existing word", func (t *testing.T){
		key := "test"
		value := "this is just a test"
		dictionary = Dictionary{key:value}

		err := dictionary.Add(key, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T){
		key := "test"
		value := "this is just a test"
		dictionary:= Dictionary{key: value}
		newValue := "new value"
		dictionary.Update(key, newValue)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new word", func (t *testing.T){
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T){
	t.Run("existing word", func (t *testing.T){
		key := "test"
		dictionary := Dictionary{key:"test value"}
		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		if err != ErrNotFound{
			t.Errorf("expected %q to be deleted",key)
		}

	})
}


func assertDefinition(t testing.TB, dictionary Dictionary,key, value string){
	t.Helper()

	got,err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}
	if value != got{
		t.Errorf("got %q want %q", got, value)
	}
}



func assertError(t testing.TB, got, want error){
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q",got , want)
	}
}