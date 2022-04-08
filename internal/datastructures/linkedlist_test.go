package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedListCreation(t *testing.T) {
	ll := NewLinkedList()
	require.EqualValues(t, ll, &LinkedList{})
}

func TestLinkedListAdd(t *testing.T) {
	ll := NewLinkedList()
	testAddingToLinkedList(t, ll)
	testAddingSameKeyToLinkedList(t, ll)
}

func testAddingToLinkedList(t *testing.T, ll *LinkedList) {
	t.Helper()
	ll.Add(3, 4)
	require.NotNil(t, ll.head)
	require.Equal(t, ll.head.key, uint64(3))
	require.Equal(t, ll.head.value, uint64(4))
	require.Nil(t, ll.head.next)
}

func testAddingSameKeyToLinkedList(t *testing.T, ll *LinkedList) {
	t.Helper()
	ll.Add(3, 5)
	require.Equal(t, ll.head.key, uint64(3))
	require.Equal(t, ll.head.value, uint64(5))
}

func TestLinkedListSearch(t *testing.T) {
	ll := NewLinkedList()
	testCantSearchEmptyLinkedList(t, ll)
	ll.Add(3, 4)
	ll.Add(2, 3)
	ll.Add(4, 2)
	ll.Add(5, 6)
	testCantFindUnknownKey(t, ll)
	testCanFindKnownKey(t, ll)
}

func TestLinkedListDelete(t *testing.T) {
	ll := NewLinkedList()
	ll.Add(3, 4)
	ll.Add(2, 3)
	ll.Add(4, 2)
	ll.Add(5, 6)
	ll.Delete(3)
	ll.Delete(4)
	_, err := ll.Search(5)
	require.NoError(t, err)
	_, err = ll.Search(4)
	require.Error(t, err)
	_, err = ll.Search(3)
	require.Error(t, err)
}

func testCantSearchEmptyLinkedList(t *testing.T, ll *LinkedList) {
	t.Helper()
	_, err := ll.Search(0)
	require.Error(t, err)
}

func testCantFindUnknownKey(t *testing.T, ll *LinkedList) {
	t.Helper()
	_, err := ll.Search(32)
	require.Error(t, err)
}

func testCanFindKnownKey(t *testing.T, ll *LinkedList) {
	t.Helper()
	v, err := ll.Search(5)
	require.NoError(t, err)
	require.Equal(t, v, uint64(6))
}
