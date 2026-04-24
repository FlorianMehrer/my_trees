package binsearchtree_int

// GetElement liefert das Element mit dem gegebenen Wert zurück, oder nil, wenn es nicht gefunden wird.
func (e *Element) GetElement(value int) *Element {
	// Hinweis:
	// Gehen Sie rekursiv vor:
	// - Wenn das aktuelle Element leer ist, geben Sie nil zurück.
	// - Wenn der Wert des aktuellen Elements gleich dem gesuchten Wert ist, geben Sie das aktuelle Element zurück.
	// - Wenn der gesuchte Wert kleiner als der Wert des aktuellen Elements ist, suchen Sie im linken Teilbaum weiter.
	// - Wenn der gesuchte Wert größer als der Wert des aktuellen Elements ist, suchen Sie im rechten Teilbaum weiter.
	//
	// Verwenden Sie die Methoden Value(), Left() und Right(), um auf die Werte und Kinder des Elements zuzugreifen.

	if e.IsEmpty() {
		return nil
	}
	if value == e.Value() {
		return e
	} else if value < e.Value() {
		return e.Left().GetElement(value)
	} else {
		return e.Right().GetElement(value)
	}
}

// InsertValue fügt einen Wert nach den Regeln eines binären Suchbaums in den Baum ein.
func (e *Element) InsertValue(value int) {
	// Hinweis:
	// Gehen Sie rekursiv vor:
	// - Wenn das aktuelle Element leer ist, setzen Sie den Wert des aktuellen Elements auf den neuen Wert.
	// - Wenn der neue Wert kleiner oder gleich dem Wert des aktuellen Elements ist, fügen Sie den neuen Wert in den linken Teilbaum ein.
	// - Wenn der neue Wert größer als der Wert des aktuellen Elements ist, fügen Sie den neuen Wert in den rechten Teilbaum ein.
	//
	// Verwenden Sie die Methoden Value(), Left() und Right(), um auf die Werte und Kinder des Elements zuzugreifen.

	if e.IsEmpty() {
		e.SetValue(value)
		return
	}

	if value <= e.Value() {
		e.Left().InsertValue(value)
		return
	}
	e.Right().InsertValue(value)
}
