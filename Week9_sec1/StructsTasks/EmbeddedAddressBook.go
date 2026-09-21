package main

import (
	"fmt"
	"strings"
)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Contact struct {
	Name  string
	Email string
	Phone string
	Address
}

type BusinessContact struct {
	Contact
	Company  string
	JobTitle string
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

func (c Contact) String() string {
	return fmt.Sprintf("Name: %s | Email: %s | Phone: %s | Address: [%s]",
		c.Name, c.Email, c.Phone, c.Address.String())
}

func (bc BusinessContact) String() string {
	return fmt.Sprintf("Business Contact: %s (%s at %s) | Email: %s | Phone: %s | Address: [%s]",
		bc.Name, bc.JobTitle, bc.Company, bc.Email, bc.Phone, bc.Address.String())
}

type AddressBook struct {
	Contacts         []Contact
	BusinessContacts []BusinessContact
}

func (ab *AddressBook) SearchByCity(city string) ([]string, []string) {
	var matchedContacts []string
	var matchedBusiness []string

	query := strings.ToLower(city)

	for _, c := range ab.Contacts {
		if strings.ToLower(c.City) == query {
			matchedContacts = append(matchedContacts, c.String())
		}
	}

	for _, bc := range ab.BusinessContacts {
		if strings.ToLower(bc.City) == query {
			matchedBusiness = append(matchedBusiness, bc.String())
		}
	}

	return matchedContacts, matchedBusiness
}

func (ab *AddressBook) SearchByPartialName(partialName string) ([]string, []string) {
	var matchedContacts []string
	var matchedBusiness []string

	query := strings.ToLower(partialName)

	for _, c := range ab.Contacts {
		if strings.Contains(strings.ToLower(c.Name), query) {
			matchedContacts = append(matchedContacts, c.String())
		}
	}

	for _, bc := range ab.BusinessContacts {
		if strings.Contains(strings.ToLower(bc.Name), query) {
			matchedBusiness = append(matchedBusiness, bc.String())
		}
	}

	return matchedContacts, matchedBusiness
}

func main() {
	ab := AddressBook{
		Contacts: []Contact{
			{
				Name:  "Ahmed Ali",
				Email: "ahmed@example.com",
				Phone: "+201000000001",
				Address: Address{
					Street:  "123 El-Gesh St",
					City:    "Cairo",
					State:   "Cairo",
					ZipCode: "11511",
				},
			},
			{
				Name:  "Sara Hassan",
				Email: "sara@example.com",
				Phone: "+201000000002",
				Address: Address{
					Street:  "45 Corniche St",
					City:    "Alexandria",
					State:   "Alexandria",
					ZipCode: "21500",
				},
			},
		},
		BusinessContacts: []BusinessContact{
			{
				Contact: Contact{
					Name:  "Mohamed Ahmed",
					Email: "m.ahmed@company.com",
					Phone: "+201000000003",
					Address: Address{
						Street:  "78 Smart Village",
						City:    "Giza",
						State:   "Giza",
						ZipCode: "12511",
					},
				},
				Company:  "Tech Solutions",
				JobTitle: "Senior Developer",
			},
			{
				Contact: Contact{
					Name:  "Ahmed Mahmoud",
					Email: "a.mahmoud@corp.com",
					Phone: "+201000000004",
					Address: Address{
						Street:  "12 Fouad St",
						City:    "Alexandria",
						State:   "Alexandria",
						ZipCode: "21501",
					},
				},
				Company:  "Global Trade",
				JobTitle: "Sales Manager",
			},
		},
	}

	fmt.Println("--- Field Promotion Example ---")
	fmt.Println("Contact Name:", ab.Contacts[0].Name)
	fmt.Println("Promoted City Field:", ab.Contacts[0].City)

	fmt.Println("\n--- Search By City: 'Alexandria' ---")
	contacts, business := ab.SearchByCity("Alexandria")
	for _, c := range contacts {
		fmt.Println(c)
	}
	for _, b := range business {
		fmt.Println(b)
	}

	fmt.Println("\n--- Search By Partial Name: 'ahmed' ---")
	contacts, business = ab.SearchByPartialName("ahmed")
	for _, c := range contacts {
		fmt.Println(c)
	}
	for _, b := range business {
		fmt.Println(b)
	}
}
