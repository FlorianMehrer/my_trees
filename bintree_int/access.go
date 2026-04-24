package bintree_int

import "fmt"

// SetValue setzt den Wert des Elements auf den gegebenen Wert.
// Falls das Element vorher leer war, werden leere Kind-Elemente erstellt.
func (e *Element) SetValue(value int) {
	// Hinweis:
	// Setzen Sie den Wert des Elements auf den gegebenen Wert.
	// Prüfen Sie außerdem, ob das Element vorher leer war, z.B. mittels e.IsEmpty().

	// begin:solution
	e.value = value
	if e.IsEmpty() {
		e.left = Empty()
		e.right = Empty()
	}
	// end:solution
}

// GetValueAt gibt den Wert des Elements zurück, das durch den gegebenen Pfad erreicht wird.
// Beispiele:
// - "LR": Gehe einmal nach links und dann einmal nach rechts.
// - "LLR": Gehe zwei Mal nach links und dann einmal nach rechts.
// - "": Gib den Wert der Wurzel zurück.
//
// Als zweiter Rückgabewert wird ein Fehlerwert geliefert.
// Wenn der Pfad zu einem leeren Element führt oder ein Element auf dem Pfad nicht existiert,
// liefert die Funktion 0 und einen Fehler, andernfalls den Wert des Elements und nil.
func (e *Element) GetValueAt(path string) (int, error) {
	// Hinweis:
	// Gehen Sie rekursiv vor:
	// - Wenn der Pfad leer ist, geben Sie den Wert des aktuellen Elements zurück.
	// - Ansonsten prüfen Sie das erste Zeichen des Pfads und gehen entweder
	//   nach links oder rechts weiter, indem Sie die Funktion rekursiv aufrufen.
	// - Wenn das aktuelle Element leer ist, prüfen Sie, ob der Pfad noch weitere
	//   Schritte entält. Wenn ja, liefern Sie einen Fehler zurück, andernfalls 0.

	// begin:solution
	if path == "" {
		if e.IsEmpty() {
			return 0, fmt.Errorf("Element is empty")
		}
		return e.value, nil
	}
	if e.IsEmpty() {
		return 0, fmt.Errorf("Element is empty")
	}
	switch path[0] {
	case 'L':
		return e.left.GetValueAt(path[1:])
	case 'R':
		return e.right.GetValueAt(path[1:])
	default:
		return 0, fmt.Errorf("Invalid path character: %c", path[0])
	}
	// end:solution
}

// InsertValueAt fügt einen neuen Wert in den Baum ein.
// Die Stelle, an der eingefügt wird, wird durch den gegebenen Pfad bestimmt.
// Beispiele:
// - "LR": Gehe einmal nach links und dann einmal nach rechts.
// - "LLR": Gehe zwei Mal nach links und dann einmal nach rechts.
// - "": Füge direkt in der Wurzel ein.
//
// Wenn der Pfad zu einem leeren Element führt, wird dort ein neues Element
// mit dem gegebenen Wert erstellt.
// Wenn ein Element auf dem Pfad nicht existiert (außer dem letzten),
// liefert die Funktion einen Fehler, andernfalls true.
func (e *Element) InsertValueAt(path string, value int) error {
	// Hinweis:
	// Gehen Sie rekursiv vor:
	// - Wenn der Pfad leer ist, setzen Sie den Wert des aktuellen Elements.
	// - Ansonsten prüfen Sie das erste Zeichen des Pfads und gehen entweder
	//   nach links oder rechts weiter, indem Sie die Funktion rekursiv aufrufen.
	// - Wenn das aktuelle Element leer ist, prüfen Sie, ob der Pfad noch weitere
	//   Schritte entält.

	// begin:solution
	if path == "" {
		e.SetValue(value)
		return nil
	}
	if e.IsEmpty() {
		return fmt.Errorf("Element is empty")
	}
	switch path[0] {
	case 'L':
		return e.left.InsertValueAt(path[1:], value)
	case 'R':
		return e.right.InsertValueAt(path[1:], value)
	default:
		return fmt.Errorf("Invalid path character: %c", path[0])
	}
	// end:solution
}
