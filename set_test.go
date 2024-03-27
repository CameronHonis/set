package set_test

import (
	. "github.com/CameronHonis/set"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Set", func() {
	Describe("::Add", func() {
		It("adds the element to the set", func() {
			s := EmptySet[int]()
			s.Add(12)
			Expect(s.Size()).To(Equal(1))
		})
		When("the element already exists in the set", func() {
			var s *Set[int]
			BeforeEach(func() {
				s = EmptySet[int]()
				s.Add(12)
			})
			It("does not error", func() {
				s.Add(12)
			})
		})
	})
	Describe("::Has", func() {
		When("the set does not have the item", func() {
			It("returns false", func() {
				s := EmptySet[string]()
				Expect(s.Has("asdf")).To(BeFalse())
			})
		})
		When("the set contains the item", func() {
			It("returns true", func() {
				s := EmptySet[string]()
				s.Add("asdf")
				Expect(s.Has("asdf")).To(BeTrue())
			})
		})
	})
	Describe("::Remove", func() {
		var s *Set[string]
		BeforeEach(func() {
			s = EmptySet[string]()
			s.Add("asdf")
			Expect(s.Size()).To(Equal(1))
			Expect(s.Has("asdf")).To(BeTrue())
		})
		It("removes the item from the list", func() {
			s.Remove("asdf")
			Expect(s.Has("asdf")).To(BeFalse())
			Expect(s.Size()).To(Equal(0))
		})
		When("the item does not exist in the set", func() {
			It("does not panic", func() {
				s.Remove("a")
			})
		})
	})
	Describe("::Flatten", func() {
		It("returns a slice of all the items in the set", func() {
			s := EmptySet[int]()
			s.Add(1)
			s.Add(2)
			s.Add(3)
			expFlattenedSet := []int{1, 2, 3}
			realFlattenedSet := s.Flatten()
			Expect(len(realFlattenedSet)).To(Equal(len(expFlattenedSet)))
			for _, item := range realFlattenedSet {
				foundMatch := false
				for _, expItem := range expFlattenedSet {
					if item == expItem {
						foundMatch = true
						break
					}
				}
				Expect(foundMatch).To(BeTrue())
			}
		})
		When("the set is mutated after Flatten is called", func() {
			var s *Set[int]
			BeforeEach(func() {
				s = EmptySet[int]()
				s.Add(1)
				s.Add(2)
				s.Add(3)
				Expect(s.Size()).To(Equal(3))
				flatS := s.Flatten()
				Expect(len(flatS)).To(Equal(3))
				s.Remove(3)
				Expect(s.Size()).To(Equal(2))
			})
			It("does not keep a stale flattened array in memory", func() {
				flatS := s.Flatten()
				Expect(flatS).To(HaveLen(2))
			})
		})
	})
	Describe("::Union", func() {
		var s1, s2 *Set[int]
		BeforeEach(func() {
			s1 = EmptySet[int]()
			s2 = EmptySet[int]()
		})
		When("the sets are disjoint", func() {
			BeforeEach(func() {
				s1.Add(1)
				s1.Add(2)
				s2.Add(10)
				s2.Add(20)
			})
			It("returns a set that exactly contains both sets", func() {
				s3 := s1.Union(s2)
				Expect(s3.Flatten()).To(HaveLen(4))
				Expect(s3.Has(1)).To(BeTrue())
				Expect(s3.Has(2)).To(BeTrue())
				Expect(s3.Has(10)).To(BeTrue())
				Expect(s3.Has(20)).To(BeTrue())
			})
		})
		When("the sets share some elements", func() {
			BeforeEach(func() {
				s1.Add(1)
				s1.Add(2)
				s2.Add(2)
				s2.Add(3)
			})
			It("returns a set that exactly contains both sets", func() {
				s3 := s1.Union(s2)
				Expect(s3.Flatten()).To(HaveLen(3))
				Expect(s3.Has(1)).To(BeTrue())
				Expect(s3.Has(2)).To(BeTrue())
				Expect(s3.Has(3)).To(BeTrue())
			})
		})
	})
	Describe("::Intersect", func() {
		var s1, s2 *Set[int]
		BeforeEach(func() {
			s1 = EmptySet[int]()
			s2 = EmptySet[int]()
		})
		When("the sets are disjoint", func() {
			BeforeEach(func() {
				s1.Add(1)
				s1.Add(3)
				s2.Add(20)
				s2.Add(40)
			})
			It("returns an empty set", func() {
				s3 := s1.Intersect(s2)
				Expect(s3.Flatten()).To(HaveLen(0))
			})
		})
		When("the sets share some elements", func() {
			BeforeEach(func() {
				s1.Add(1)
				s1.Add(2)
				s1.Add(3)

				s2.Add(2)
				s2.Add(3)
				s2.Add(4)
			})
			It("returns the intersection set", func() {
				s3 := s1.Intersect(s2)
				Expect(s3.Flatten()).To(HaveLen(2))
				Expect(s3.Has(2))
				Expect(s3.Has(3))
			})
		})
	})
	Describe("::Copy", func() {
		var set *Set[string]
		BeforeEach(func() {
			set = EmptySet[string]()
			set.Add("asdf")
		})
		It("Returns a shallow copy of the set", func() {
			cpSet := set.Copy()
			cpSet.Add("jkl")
			Expect(set.Has("asdf")).To(BeTrue())
			Expect(set.Has("jkl")).To(BeFalse())
			Expect(cpSet.Has("asdf")).To(BeTrue())
			Expect(cpSet.Has("jkl")).To(BeTrue())
		})
	})
	Describe("#EmptySet", func() {
		It("returns a set of the specified generic", func() {
			s := EmptySet[uint16]()
			Expect(reflect.TypeOf(*s)).To(Equal(reflect.TypeOf(Set[uint16]{})))
		})
		It("returns an empty set", func() {
			s := EmptySet[string]()
			Expect(s.Size()).To(Equal(0))
		})
	})
})
