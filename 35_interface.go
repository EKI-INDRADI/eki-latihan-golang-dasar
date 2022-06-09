package main

import "fmt"

// ========================= HasName Intreface =========================

type HasName interface { // nama kontrak
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
	// value apapun yang menggunakan SayHallo dan schema nya sama dengan interface HasName maka
	// dia berhak mengikuti kontrak SayHello
}

// ========================= HasName Intreface =========================

// ========================= PERSON =========================

type Person struct {
	Name string
}

func (person Person) GetName() string { // untuk return struct person namanya doang (anggap aja struct mirip object json)

	// function ini udah otomatis ngikutin kontrak HasName
	// agar struct bisa mengimplementasikan
	// pada main()

	//  var Eki Person
	//  Eki.Name = "Eki"

	return person.Name // jari hanya return nama nya doang
}

// ========================= PERSON =========================

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface data secara abstract, ga bisa buat di dalam interface
// mirip interface typescript
func main() {

	// ========================= PERSON =========================
	var Eki Person
	Eki.Name = "Eki"
	// Eki.Name = 12321 // gak bakalan bisa di pake dan langsung error karena bukan string (syarat interface)
	// ========================= PERSON =========================

	SayHello(Eki) // return nama nya doang dari struct Eki, maka bisa di pakai interface ,
	//sama aja kaya SayHello(Eki.Name) untuk versi manual nya tanpa  func (person Person) GetName() string {

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}
