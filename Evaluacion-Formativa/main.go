package main

import "fmt"

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate string
	shelfLife       int
}

func NewMedicine(name, manufacturer, date string, shelfLife int) Medicine {
	return Medicine{
		name:            name,
		manufacturer:    manufacturer,
		manufactureDate: date,
		shelfLife:       shelfLife,
	}
}

func (m Medicine) GetName() string {
	return m.name
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m Medicine) GetManufactureDate() string {
	return m.manufactureDate
}

func (m Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(newName string) {
	if newName == "" {
		fmt.Println("Error esta vacio el nombre")
	} else {
		m.name = newName
	}
}

func (m *Medicine) SetManufacturer(newManufacturer string) {
	if newManufacturer == "" {
		fmt.Println("Error esta vacio el fabricante")
	} else {
		m.manufacturer = newManufacturer
	}
}

func (m *Medicine) SetShelLife(newShelLife int) {
	if newShelLife > 0 {
		m.shelfLife = newShelLife
	}
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(m Medicine, dose float64, prescription bool) *Tablet {
	return &Tablet{
		Medicine:       m,
		dosePerTablet:  dose,
		isPrescription: prescription,
	}
}

func (t Tablet) MostrarInfo() {
	fmt.Println("DETALLE DE TABLET")
	fmt.Println("El nombre es: ", t.name)
	fmt.Println("El manufacturador es: ", t.manufacturer)
	fmt.Println("La fecha de fabricación es: ", t.manufactureDate)
	fmt.Println("La vida útil en meses es: ", t.shelfLife, " meses")
	fmt.Println("La dosis por tableta es: ", t.dosePerTablet, " mg")
	fmt.Println("Requiere receta: ", t.isPrescription)
}

func (t *Tablet) SetDose(newDose float64) {
	if newDose > 0 {
		t.dosePerTablet = newDose
	} else {
		fmt.Println("Error La dosis debe ser mayor a 0")
	}
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(m Medicine, volume float64, flavor string) *Syrup {
	return &Syrup{
		Medicine: m,
		volume:   volume,
		flavor:   flavor,
	}
}

func (s Syrup) MostrarInfo() {
	fmt.Println("DETALLE DE SYRUP")
	fmt.Println("El nombre es: ", s.name)
	fmt.Println("El fabricante es: ", s.manufacturer)
	fmt.Println("La fecha de fabricación es: ", s.manufactureDate)
	fmt.Println("La vida útil en meses es: ", s.shelfLife, " meses")
	fmt.Println("Su volúmen es ", s.volume)
	fmt.Println("Su sabor es: ", s.flavor)
}

func (s *Syrup) SetVolume(newVolume float64) {
	if newVolume > 0 {
		s.volume = newVolume
	} else {
		fmt.Println("Error el volumen debe ser mayor a 0")
	}
}

func main() {
	listaTabletas := []*Tablet{}
	listaJarabes := []*Syrup{}

	basePastilla := NewMedicine("Paracetamol", "Pepito", "2026-01-01", 24)
	miTableta := NewTablet(basePastilla, 500.0, false)
	listaTabletas = append(listaTabletas, miTableta)

	baseJarabe := NewMedicine("Abrilar", "Jorgito", "2026-05-10", 12)
	miJarabe := NewSyrup(baseJarabe, 120.0, "Cereza")
	listaJarabes = append(listaJarabes, miJarabe)

	fmt.Println("     INVENTARIO DE TABLETAS")
	for _, t := range listaTabletas {
		t.MostrarInfo()
	}

	fmt.Println("     INVENTARIO DE JARABES")
	for _, s := range listaJarabes {
		s.MostrarInfo()
	}

	fmt.Println("     PRUEBA DE SETTER Y GETTER ")
	fmt.Println("Nombre original:", miTableta.GetName())

	miTableta.SetName("Paracetamol Forte")

	fmt.Println("Nombre actualizado:", miTableta.GetName())

}
