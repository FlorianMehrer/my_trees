package bintree_int

// InOrder liefert die Werte der Elemente in In-Order-Reihenfolge zurück.
//
// In-Order-Reihenfolge bedeutet, dass zuerst die Werte des linken Teilbaums, dann der
// Wert der Wurzel und schließlich die Werte des rechten Teilbaums zurückgegeben werden.
func (e *Element) InOrder() []int {
	// Hinweis:
	// - Wenn das aktuelle Element leer ist, geben Sie eine leere Liste zurück.
	// - Ansonsten rufen Sie `InOrder` rekursiv ür den linken und rechten Teilbaum auf
	//   und hängen die Ergebnisse zusammen mit dem Wurzelwert in der passenden
	//   Reihenfolge zusammen.

	// begin:solution
	if e.IsEmpty() {
		return []int{}
	}
	result := []int{}
	result = append(result, e.left.InOrder()...)
	result = append(result, e.value)
	result = append(result, e.right.InOrder()...)
	return result
	// end:solution
}

// PreOrder liefert die Werte der Elemente in Pre-Order-Reihenfolge zurück.
//
// Pre-Order-Reihenfolge bedeutet, dass zuerst der Wert der Wurzel, dann die Werte
// des linken Teilbaums und schließlich die Werte des rechten Teilbaums zurückgegeben werden.
func (e *Element) PreOrder() []int {
	// Hinweis:
	// - Wenn das aktuelle Element leer ist, geben Sie eine leere Liste zurück.
	// - Ansonsten rufen Sie `PreOrder` rekursiv ür den linken und rechten Teilbaum auf
	//   und hängen die Ergebnisse zusammen mit dem Wurzelwert in der passenden
	//   Reihenfolge zusammen.

	// begin:solution
	if e.IsEmpty() {
		return []int{}
	}
	result := []int{}
	result = append(result, e.value)
	result = append(result, e.left.PreOrder()...)
	result = append(result, e.right.PreOrder()...)
	return result
	// end:solution
}

// PostOrder liefert die Werte der Elemente in Post-Order-Reihenfolge zurück.
//
// Post-Order-Reihenfolge bedeutet, dass zuerst die Werte des linken Teilbaums, dann
// die Werte des rechten Teilbaums und schließlich der Wert der Wurzel zurückgegeben werden.
func (e *Element) PostOrder() []int {
	// Hinweis:
	// - Wenn das aktuelle Element leer ist, geben Sie eine leere Liste zurück.
	// - Ansonsten rufen Sie `PostOrder` rekursiv ür den linken und rechten Teilbaum auf
	//   und hängen die Ergebnisse zusammen mit dem Wurzelwert in der passenden
	//   Reihenfolge zusammen.

	// begin:solution
	if e.IsEmpty() {
		return []int{}
	}
	result := []int{}
	result = append(result, e.left.PostOrder()...)
	result = append(result, e.right.PostOrder()...)
	result = append(result, e.value)
	return result
	// end:solution
}
