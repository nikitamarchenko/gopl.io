/*
Exercise 7.8: Many GUIs provide a table widget with a stateful multi-tier sort:
the primary sort key is the most recently clicked column head, the secondary
sort key is the second-most recently clicked column head, and so on. Define an
implementation of sort.Interface for use by such a table. Compare that approach
with repeated sorting using sort.Stable.
*/

package multi_tier_sort_table

import "time"

type Citizen struct {
	Name string
	Born time.Time
}

type CitizenTableSort struct {
	Cs   []*Citizen
	comp []func(i, j *Citizen) int
}

func (ct CitizenTableSort) Len() int      { return len(ct.Cs) }
func (ct CitizenTableSort) Swap(i, j int) { ct.Cs[i], ct.Cs[j] = ct.Cs[j], ct.Cs[i] }
func (ct CitizenTableSort) Less(i, j int) bool {
	if len(ct.comp) > 0 {
		for _, comp := range ct.comp {
			r := comp(ct.Cs[i], ct.Cs[j])
			if r == 0 {
				continue
			}
			return r < 0
		}
	}
	return ct.Cs[i].Name < ct.Cs[j].Name
}

func (ct *CitizenTableSort) Add(f func(i, j *Citizen) int) {
	ct.comp = append([]func(i, j *Citizen) int {f}, ct.comp...)
}

func CitizenTableSortByName(i, j *Citizen) int {
	if i.Name == j.Name {
		return 0
	} else if i.Name < j.Name {
		return -1
	}
	return 1
}

func CitizenTableSortByBorn(i, j *Citizen) int {
	if i.Born == j.Born {
		return 0
	} else if i.Born.Before(j.Born) {
		return -1
	}
	return 1
}