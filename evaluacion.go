package main

import (
	"fmt"
	"time"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func (m *Medicine) GetName() string {
	return m.name
}

func (m *Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetName(name string) bool {
	if name == "" {
		fmt.Println("Error: el nombre no puede estar vacío")
		return false
	}
	m.name = name
	return true
}

func (m *Medicine) SetManufacturer(manufacturer string) bool {
	if manufacturer == "" {
		fmt.Println("Error: el fabricante no puede estar vacío")
		return false
	}
	m.manufacturer = manufacturer
	return true
}

func (m *Medicine) SetManufactureDate(date string) bool {
	fecha, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error: fecha inválida (use el formato AAAA-MM-DD):", date)
		return false
	}
	m.manufactureDate = fecha
	return true
}

func (m *Medicine) SetShelfLife(months int) bool {
	if months <= 0 {
		fmt.Println("Error: la vida útil debe ser un número positivo")
		return false
	}
	m.shelfLife = months
	return true
}

func NewMedicine(name, manufacturer, date string, shelfLife int) (Medicine, bool) {
	m := Medicine{}
	if !m.SetName(name) {
		return m, false
	}
	if !m.SetManufacturer(manufacturer) {
		return m, false
	}
	if !m.SetManufactureDate(date) {
		return m, false
	}
	if !m.SetShelfLife(shelfLife) {
		return m, false
	}
	return m, true
}

func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

type Tablet struct {
	Medicine
	dosePerTablet float64
}

func NewTablet(name, manufacturer, date string, shelfLife int, dose float64) *Tablet {
	base, ok := NewMedicine(name, manufacturer, date, shelfLife)
	if !ok {
		return nil
	}
	if dose <= 0 {
		fmt.Println("Error: la dosis por tableta debe ser mayor a 0")
		return nil
	}
	return &Tablet{Medicine: base, dosePerTablet: dose}
}

func (t *Tablet) ShowDetails() {
	fmt.Println("[Tableta]")
	fmt.Println("Nombre:", t.name)
	fmt.Println("Fabricante:", t.manufacturer)
	fmt.Println("Fecha de fabricación:", t.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida útil:", t.shelfLife, "meses")
	fmt.Println("Dosis por tableta:", t.dosePerTablet, "mg")
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(name, manufacturer, date string, shelfLife int, volume float64, flavor string) *Syrup {
	base, ok := NewMedicine(name, manufacturer, date, shelfLife)
	if !ok {
		return nil
	}
	if volume <= 0 {
		fmt.Println("Error: el volumen debe ser mayor a 0")
		return nil
	}
	return &Syrup{Medicine: base, volume: volume, flavor: flavor}
}

func (s *Syrup) ShowDetails() {
	fmt.Println("[Jarabe]")
	fmt.Println("Nombre:", s.name)
	fmt.Println("Fabricante:", s.manufacturer)
	fmt.Println("Fecha de fabricación:", s.manufactureDate.Format("2006-01-02"))
	fmt.Println("Vida útil:", s.shelfLife, "meses")
	fmt.Println("Volumen:", s.volume, "ml")
	fmt.Println("Sabor:", s.flavor)
}

func main() {
	tab1 := NewTablet("Paracetamol", "Genfar", "2025-03-15", 24, 500)
	tab2 := NewTablet("Amoxicilina", "Sanofi", "2025-06-01", 18, 875)
	syr1 := NewSyrup("Ambroxol", "Bayer", "2025-09-10", 12, 120, "Fresa")
	syr2 := NewSyrup("Ibuprofeno Pediátrico", "Pfizer", "2025-01-20", 36, 100, "Naranja")

	if tab1 == nil || tab2 == nil || syr1 == nil || syr2 == nil {
		fmt.Println("No se pudo crear algún medicamento.")
		return
	}

	tablets := []*Tablet{}
	tablets = append(tablets, tab1)
	tablets = append(tablets, tab2)

	syrups := []*Syrup{}
	syrups = append(syrups, syr1)
	syrups = append(syrups, syr2)

	fmt.Println("===== INVENTARIO DE TABLETAS =====")
	for _, t := range tablets {
		t.ShowDetails()
		fmt.Println()
	}

	fmt.Println("===== INVENTARIO DE JARABES =====")
	for _, s := range syrups {
		s.ShowDetails()
		fmt.Println()
	}

	fmt.Println("===== ACTUALIZACIÓN CON SETTER =====")
	fmt.Println("Vida útil antes:", tab1.GetName(), "-", tab1.GetShelfLife(), "meses")
	tab1.SetShelfLife(36)
	fmt.Println("Vida útil después:", tab1.GetName(), "-", tab1.GetShelfLife(), "meses")
	fmt.Println()

	fmt.Println("===== FECHAS DE CADUCIDAD =====")
	for _, t := range tablets {
		fmt.Println(t.GetName(), "->", t.ExpirationDate().Format("2006-01-02"))
	}
	for _, s := range syrups {
		fmt.Println(s.GetName(), "->", s.ExpirationDate().Format("2006-01-02"))
	}
}