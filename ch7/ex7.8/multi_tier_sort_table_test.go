package multi_tier_sort_table

import (
	"sort"
	"testing"
	"time"
)

func genTable() []*Citizen {
	return []*Citizen{
		{Name: "Bred", Born: time.Date(1984, 01, 01, 0, 0, 0, 0, time.UTC)},
		{Name: "Anna", Born: time.Date(1985, 01, 01, 0, 0, 0, 0, time.UTC)},
		{Name: "Clara", Born: time.Date(1987, 01, 01, 0, 0, 0, 0, time.UTC)},
	}
}

func genTable2() []*Citizen {
	return []*Citizen{
		{Name: "Clara", Born: time.Date(1987, 01, 01, 0, 0, 0, 0, time.UTC)},
		{Name: "Anna", Born: time.Date(1984, 01, 01, 0, 0, 0, 0, time.UTC)},
		{Name: "Clara", Born: time.Date(1980, 01, 01, 0, 0, 0, 0, time.UTC)},
	}
}

func TestCitizenTable(t *testing.T) {

	table := genTable()
	sort.Sort(CitizenTableSort{Cs: table})
	if table[0].Name != "Anna" {
		t.Error("invalid default sort")
	}
	
	table = genTable()
	cts := CitizenTableSort{Cs: table}
	cts.Add(CitizenTableSortByName)
	sort.Sort(cts)

	if table[0].Name != "Anna" {
		t.Error("invalid CitizenTableSortByName sort")
	}
	
	cts.Add(CitizenTableSortByBorn)
	sort.Sort(cts)
	if table[0].Name != "Bred" {
		t.Error("invalid CitizenTableSortByBorn sort")
	}


	table = genTable2()
	cts = CitizenTableSort{Cs: table}
	cts.Add(CitizenTableSortByBorn)
	cts.Add(CitizenTableSortByName)
	sort.Sort(cts)

	if table[0].Name != "Anna" {
		t.Error("invalid complex sort 1")
	}
	
	if table[1].Born.Year() != 1980 {
		t.Error("invalid complex sort 2")
	}
}
