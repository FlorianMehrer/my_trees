package bintree_int

// IsEmpty liefert true zurück, wenn das Element leer ist.
//
// Ein Element gilt als leer, wenn es keine Kinder hat.
// D.h. wenn die Kind-Pointer Left und Right beide nil sind.
// Der Wert des Elements spielt dabei keine Rolle.
func (e *Element) IsEmpty() bool {
	// Hinweis:
	// Prüfen Sie, ob beide Kind-Pointer nil sind.

	// begin:solution
	return e.left == nil && e.right == nil
	// end:solution
}

// IsLeaf liefert true zurück, wenn das Element ein Blatt ist.
//
// Ein Element gilt als Blatt, wenn beide Kinder leer sind.
func (e *Element) IsLeaf() bool {
	// Hinweis:
	// Nutzen Sie die IsEmpty-Methode, um zu prüfen, ob beide Kinder leer sind,
	// aber stellen Sie sicher, dass das Element e selbst nicht leer ist.

	// begin:solution
	return !e.IsEmpty() && e.left.IsEmpty() && e.right.IsEmpty()
	// end:solution
}

// Count zählt die Anzahl der Elemente im Baum, beginnend bei diesem Element.
// Dabei werden nur nicht-leere Elemente gezählt.
func (e *Element) Count() int {
	// Hinweis:
	// Prüfen Sie zunächst, ob das aktuelle Element leer ist. Wenn ja, geben Sie 0 zurück.
	// Ansonsten nutzen Sie Count rekursiv, um die Anzahl der Elemente in den linken
	// und rechten Teilbäumen zu zählen.

	// begin:solution
	if e.IsEmpty() {
		return 0
	}
	return 1 + e.left.Count() + e.right.Count()
	// end:solution
}

// Height berechnet die Höhe des Baums.
// Ein leeres Element hat die Höhe 0, ein Blatt hat die Höhe 1, usw.
func (e *Element) Height() int {
	// Hinweis:
	// Prüfen Sie zunächst, ob das aktuelle Element leer ist. Wenn ja, geben Sie 0 zurück.
	// Ansonsten berechnen Sie die Höhe der linken und rechten Teilbäume rekursiv.
	// Die Höhe von e ist dann 1 plus die größere der beiden Höhen.

	// begin:solution
	if e.IsEmpty() {
		return 0
	}
	leftHeight := e.left.Height()
	rightHeight := e.right.Height()
	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
	// end:solution
}
