package collection

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

type Foo struct {
	A string
	B int
}

func TestObjCollection_Insert(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}
	a3 := Foo{A: "a3"}

	{
		objColl := NewObjCollection([]Foo{a1, a2})
		objColl2 := objColl.Insert(0, a3)
		objColl2.DD()
		i, err := objColl2.Index(0).ToInterface()
		if err != nil {
			t.Fatal(err)
		}
		f1 := i.(Foo)
		if !reflect.DeepEqual(f1, a3) {
			t.Fatal("insert 0 error 1")
		}
		i1, err := objColl2.Index(1).ToInterface()
		if err != nil {
			t.Fatal(err)
		}
		f2 := i1.(Foo)
		if !reflect.DeepEqual(f2, a1) {
			t.Fatal("insert 0 error 2")
		}
	}

	{

		objColl := NewObjCollection([]Foo{a1, a2})
		objColl2 := objColl.Insert(1, a3)
		i, err := objColl2.Index(1).ToInterface()
		if err != nil {
			t.Fatal(err)
		}
		f1 := i.(Foo)
		if !reflect.DeepEqual(f1, a3) {
			t.Fatal("insert 0 error")
		}
		objColl2.DD()
	}
}

func TestObjCollection_Pluck(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})

	strColl := objColl.Pluck("A")

	strColl.DD()

	str, err := strColl.Index(0).ToString()
	if err != nil {
		t.Fatal(err)
	}

	if str != "a1" {
		t.Fatal(errors.New("Pluck error"))
	}
}

func TestObjCollection_SortBy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.SortBy("B")
	count := newObjColl.Count()
	if count != 3 {
		t.Fatal("sort By count error")
	}
	newObjColl.DD()
	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal(err)
	}

	foo := obj.(Foo)
	if foo.B != 2 {
		t.Fatal("SortBy error")
	}

	new2 := newObjColl.Slice(0, 2)
	new2.DD()
	if new2.Count() != 2 {
		t.Fatal("slice error")
	}
}
func TestObjCollection_Sort(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})
	compare := func(m interface{}, n interface{}) int {
		m1, _ := m.(Foo)
		n1, _ := n.(Foo)
		return m1.B - n1.B
	}
	objColl.SetCompare(compare)

	newObjColl := objColl.Sort()
	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal(err)
	}

	foo := obj.(Foo)
	if foo.B != 2 {
		t.Fatal("Sort error")
	}
}

func TestObjCollection_Sort_EmptyColl(t *testing.T) {

	objColl := NewObjCollection([]Foo{})
	compare := func(m interface{}, n interface{}) int {
		m1, _ := m.(Foo)
		n1, _ := n.(Foo)
		return m1.B - n1.B
	}
	objColl.SetCompare(compare)

	newObjColl := objColl.Sort()
	if newObjColl.Count() != 0 {
		t.Fatalf("Sort Empty Error")
	}
}

func TestObjCollection_EmptyColl(t *testing.T) {
	// not panic is ok
	compare := func(m interface{}, n interface{}) int {
		m1, _ := m.(Foo)
		n1, _ := n.(Foo)
		return m1.B - n1.B
	}
	NewObjCollection([]Foo{}).SetCompare(compare).Sort()
	NewObjCollection([]Foo{}).SetCompare(compare).SortDesc()
	NewObjCollection([]Foo{}).SetCompare(compare).SortByDesc("A")
	NewObjCollection([]Foo{}).SetCompare(compare).SortBy("A")
	NewObjCollection([]Foo{}).SetCompare(compare).Search(Foo{A: "a"})
	NewObjCollection([]Foo{}).SetCompare(compare).Insert(0, Foo{A: "a"})
	NewObjCollection([]Foo{}).SetCompare(compare).Unique()
	NewObjCollection([]Foo{}).SetCompare(compare).ForPage(0, 2)
	NewObjCollection([]Foo{}).SetCompare(compare).Pop()
	NewObjCollection([]Foo{}).SetCompare(compare).Reverse()
	NewObjCollection([]Foo{}).SetCompare(compare).Shuffle()
	NewObjCollection([]Foo{}).SetCompare(compare).Pluck("A")

	NewIntCollection([]int{}).Random()
}

func TestObjCollection_Copy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 5}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.Copy()

	newObjColl.DD()

	if newObjColl.Count() != 3 {
		t.Fatal("Copy count error")
	}
	inewA1, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal("Copy get first element error" + err.Error())
	}
	newA1 := inewA1.(Foo)
	if newA1.B != 3 {
		t.Fatal("Copy get first element error")
	}
}

func TestObjCollection_SortByDesc(t *testing.T) {
	a1 := Foo{A: "a1", B: 2}
	a2 := Foo{A: "a2", B: 3}
	a3 := Foo{A: "a3", B: 1}

	objColl := NewObjCollection([]Foo{a1, a2, a3})

	newObjColl := objColl.SortByDesc("B")

	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal(err)
	}

	foo := obj.(Foo)
	if foo.B != 3 {
		t.Fatal("SortBy error")
	}
}

func TestObjCollection(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})
	objColl.DD()

	if objColl.IsNotEmpty() != true {
		t.Fatal("Is Not Empty error")
	}

	if objColl.Count() != 2 {
		t.Fatal("Count error")
	}

	a3 := Foo{A: "a3"}
	a4 := Foo{A: "a4"}
	objColl.Append(a3).Append(a4)
	if objColl.Count() != 4 {
		t.Fatal("Append Error")
	}

	objColl.SetCompare(func(a interface{}, b interface{}) int {
		aObj := a.(Foo)
		bObj := b.(Foo)
		if aObj.A > bObj.A {
			return 1
		}
		if aObj.A == bObj.A {
			return 0
		}
		if aObj.A < bObj.A {
			return -1
		}
		return 0
	})

	objColl.DD()
	if objColl.Search(Foo{A: "a3"}) != 2 {
		t.Fatal("Search error")
	}

	objColl2 := objColl.Filter(func(obj interface{}, index int) bool {
		foo := obj.(Foo)
		if foo.A == "a3" {
			return true
		}
		return false
	})
	if objColl2.Count() != 1 {
		t.Fatal("Filter Error")
	}

	obj, _ := objColl.Last().ToInterface()
	if foo, ok := obj.(Foo); !ok || foo.A != "a4" {
		t.Fatal("Last error")
	}

	ret, err := objColl.Map(func(item interface{}, key int) interface{} {
		foo := item.(Foo)
		return foo.A
	}).Reduce(func(carry IMix, item IMix) IMix {
		ret, _ := carry.ToString()
		join, _ := item.ToString()
		return NewMix(ret + join)
	}).ToString()
	if err != nil {
		t.Fatal("Map error")
	}
	if ret != "a1a2a3a4" {
		t.Fatal("Reduce error")
	}

	objColl.ForPage(1, 2).DD()

	aColl := objColl.Pluck("A")
	aColl.DD()

	a0 := Foo{A: "a0"}
	objColl.Append(a0)

	objColl.Sort().DD()
	o, err := objColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal(err)
	}
	fooOut := o.(Foo)
	if fooOut.A != "a0" {
		t.Fatal("sort result error")
	}

	objColl.DD()
	objColl.SetCompare(func(a interface{}, b interface{}) int {
		aFoo := a.(Foo)
		bFoo := b.(Foo)
		return strings.Compare(aFoo.A, bFoo.A)
	})

	objColl3 := objColl.SortBy("A")
	objColl3.DD()

	objColl3.SortByDesc("A")
	objColl3.DD()

	objColl3.Remove(2)
	objColl3.DD()
}

func TestObjCollection_ToJson(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})
	byt, err := objColl.ToJson()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(byt))

	out := NewEmptyMixCollection()
	for i := 0; i < objColl.Count(); i++ {
		if i == 1 {
			out.Append(objColl.Index(i).SetField("d", 12).RemoveFields("B"))
			continue
		}
		out.Append(objColl.Index(i).SetField("c", "test").RemoveFields("A"))
	}
	byt, err = out.ToJson()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(byt))
}

func TestObjCollection_FromJson(t *testing.T) {
	data := `[{"A":"a1","B":1},{"A":"a2","B":2}]`
	dataByte := []byte(data)
	objColl := NewObjCollection([]Foo{})
	err := objColl.FromJson(dataByte)
	if err != nil {
		t.Fatal(err)
	}
	objColl.DD()
}

func TestObjCollection_Marshal(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})
	byt, err := json.Marshal(objColl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(byt))

	out := NewEmptyMixCollection()
	for i := 0; i < objColl.Count(); i++ {
		if i == 1 {
			out.Append(objColl.Index(i).SetField("d", 12).RemoveFields("B"))
			continue
		}
		out.Append(objColl.Index(i).SetField("c", "test").RemoveFields("A"))
	}
	byt, err = json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(byt))
}

func TestObjCollection_map(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})
	maps := objColl.Map(func(item interface{}, key int) interface{} {
		foo := item.(Foo)
		return foo.A
	})
	maps.DD()
	ret, err := maps.Reduce(func(carry IMix, item IMix) IMix {
		ret, _ := carry.ToString()
		join, _ := item.ToString()
		return NewMix(ret + join)
	}).ToString()
	if err != nil {
		t.Fatal("Map error " + err.Error())
	}
	if ret != "a1a2" {
		t.Fatal("Reduce error")
	}
}

func TestObjCollection_Remove(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 3}

	objColl := NewObjCollection([]Foo{a1, a2, a3})
	r := objColl.Remove(0)
	if r.Err() != nil {
		t.Fatal(r.Err())
	}
	r.DD()
}
