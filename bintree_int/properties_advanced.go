package bintree_int

// IsPerfect liefert true zurück, wenn der Baum perfekt ist.
//
// Ein Baum ist perfekt, wenn alle Blätter die gleiche Tiefe haben und
// alle inneren Knoten genau zwei Kinder haben.
func (e *Element) IsPerfect() bool {
	// Hinweis:
	// Überlegen Sie sich, wie viele Knoten ein Baum haben muss, um perfekt zu sein.
	// Dies ist eine mathematische Eigenschaft, die man nutzen kann, um die Perfektheit
	// zu überprüfen. Ggf. kann auch die Höhe des Baums eine Rolle spielen.

	// begin:solution
	if e.IsEmpty() {
		return true
	}
	height := e.Height()
	expectedCount := (1 << height) - 1 // 2^height - 1
	return e.Count() == expectedCount
	// end:solution
}

// IsComplete liefert true zurück, wenn der Baum vollständig ist.
//
// Ein Baum ist vollständig, wenn alle Ebenen bis auf die letzte vollständig gefüllt sind
// und alle Knoten der letzten Ebene so weit wie möglich nach links verschoben sind.
func (e *Element) IsComplete() bool {
	// Hinweis:
	// Ein vollständiger Baum hat die Eigenschaft, dass es keine Lücken in den Ebenen
	// gibt. D.h. sobald ein Element fehlt, darf in dieser Ebene kein weiteres Element
	// mehr folgen und alle folgenden Ebenen müssen leer sein.
	//
	// Diese Eigenschaft, kann man z.B. mit einer "Breitensuche" überprüfen.
	// Verwenden Sie dafür eine Liste von Knoten, die am Anfang nur die Wurzel enthält.
	// Verwenden Sie außerdem ein Flag, um zu verfolgen, ob bereits ein leeres Element
	// gefunden wurde. Dieses Flag ist am Anfang `false`.
	//
	// Gehen Sie dann wie folgt vor:
	// 1. Nehmen Sie das erste Element aus der Liste.
	// 2. Wenn dieses Element leer ist, setzen Sie das Flag auf `true`.
	// 3. Wenn dieses Element nicht leer ist, fügen Sie seine Kinder in die Liste ein.
	//
	// Fahren Sie mit diesem Prozess fort, bis die Liste leer ist oder bis Sie
	// ein nicht leeres Element finden, nachdem das Flag bereits `true` ist.

	// begin:solution
	nodes := []*Element{e}
	foundEmpty := false

	for len(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		if current.IsEmpty() {
			foundEmpty = true
		} else {
			if foundEmpty {
				return false // Ein nicht-leeres Element nach einem leeren Element gefunden
			}
			nodes = append(nodes, current.left, current.right)
		}
	}

	return true
	// begin:solution
}

// IsBalanced liefert true zurück, wenn der Baum balanciert ist.
//
// Ein Baum ist balanciert, wenn für jedes Element die Höhen seiner linken und rechten Teilbäume
// sich um höchstens 1 unterscheiden.
func (e *Element) IsBalanced() bool {
	// Hinweis:
	// Berechnen Sie die Höhe der linken und rechten Teilbäume und prüfen Sie, ob sie sich um mehr als 1 unterscheiden.
	// Wiederholen Sie diesen Prozess rekursiv für alle Elemente im Baum.

	// begin:solution
	if e.IsEmpty() {
		return true
	}
	leftHeight := e.left.Height()
	rightHeight := e.right.Height()
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		return false
	}
	return e.left.IsBalanced() && e.right.IsBalanced()
	// end:solution
}
