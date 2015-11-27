/* Se intentó emular la interfaz requerida por el profesor, pero no se logró. Es en base al tiempo 
entonces que se subió solo la implementación.
*/

package main

type AVL struct {
	root    *node
	n_nodes int
}

type node struct {
	father *node
	left   *node
	right  *node
	value  int
	FE     int
}

func (a *AVL) Insert(value int) bool {
	if a.root == nil {
		a.root.value = value
		a.n_nodes++
		return true
	}
	actual := a.root
	father := new(node)
	for actual != nil && actual.value != value {
		father = actual
		if value < actual.value {
			actual = actual.left
		} else {
			actual = actual.right
		}
	}
	if actual != nil {
		return false //the element exists
	}
	actual = &node{value: value, father: father}
	if value < father.value {
		father.left = actual
	} else {
		father.right = actual
	}
	plusactual := true
	if father == actual {
		father.FE--
	} else {
		father.FE++
	}
	actual = a.root.father
	a.root.father = actual.father
	switch actual.FE {
	case 0:
		plusactual = false
		break
	case -2: //rotation to the right
		if actual.left.FE == -1 {
			a.RSD(actual)
		} else {
			a.RDD(actual)
			plusactual = false
		}
		break
	case 2: //rotation to the right
		if actual.right.FE == 1 {
			a.RDI(actual)
		} else {
			a.RDI(actual)
			plusactual = false
		}
		break
	}
	for plusactual == true && father != nil {
		a.n_nodes++
		return true
	}
	return false
}

func (a *AVL) RSD(n *node) {
	//assistants
	father := n.father
	q := n.left
	// link 1
	n.left = q.right
	if n.left != nil {
		n.left.father = q
	}
	//link 2
	q.right = n
	n.father = q
	//begin rotation
	if father != nil {
		if father.right == n {
			father.right = q
		} else {
			father.left = q
		}
	} else {
		a.root = q
		q.father = father
		if q.FE == -1 {
			n.FE = 0
			q.FE = 0
		} else {
			n.FE = -1
			q.FE = 1
		}
	}
}

func (a *AVL) RDI(n *node) {
	//assistants
	father := n.father
	q := n.right
	r := q.left
	//link 1
	n.right = r.left
	if n.right != nil {
		n.right.father = n
	}
	//link 2
	q.left = r.right
	if q.left != nil {
		q.left.father = q
	}
	//link 3
	r.left = n
	n.father = r
	//link 4
	r.right = q
	q.father = r
	//begin rotation
	if father != nil {
		if father.left == n {
			father.left = r
		} else {
			father.right = r
		}
	} else {
		a.root = r
		r.father = father
	}

	switch r.FE {
	case -1:
		n.FE = 0
		q.FE = 1
		break
	case 0:
		n.FE = 0
		q.FE = 0
		break
	case 1:
		n.FE = -1
		q.FE = 0
		break
	}
	r.FE = 0
}

func (a *AVL) RDD(n *node) {
	//assistants
	father := n.father
	q := n.left
	r := q.right
	// link 1
	n.left = r.right
	if n.left != nil {
		n.left.father = n
	}
	//link 2
	q.right = r.left
	if q.right != nil {
		q.right.father = q
	}
	//link 3
	r.left = q
	q.father = r
	//link 4
	r.right = n
	n.father = r
	//begin of rotation
	if father != nil {
		if father.right == n {
			father.right = r
		} else {
			a.root = r
			r.father = father
		}
	}

	switch r.FE {
	case -1:
		n.FE = 1
		q.FE = 0
		break
	case 0:
		n.FE = 0
		q.FE = 0
		break
	case 1:
		n.FE = 0
		q.FE = -1
		break
	}
	r.FE = 0
}

func (a *AVL) search(value int) bool {
	pnode := a.root
	for pnode != nil && pnode.value != value {
		if pnode.value < value {
			pnode = pnode.left
		} else {
			pnode = pnode.right
		}
	}
	if pnode == nil {
		return false
	}
	return true
}
