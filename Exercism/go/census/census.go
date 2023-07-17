// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	} else if street, ok := r.Address["street"]; !ok {
		return false
	} else if street == "" {
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{
		Name:    "",
		Age:     0,
		Address: nil,
	}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	counter := 0
	for _, el := range residents {
		if el.HasRequiredInfo() {
			counter = counter + 1
		}
	}
	return counter
}
