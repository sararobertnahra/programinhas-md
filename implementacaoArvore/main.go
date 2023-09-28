package main

type No struct {
	valor int
	noesquerda *No
	nodireita *No
}

func imprimirno (no No){
	var valornoesquerda int
	var valornodireita int

	if no.noesquerda != nil {
		valornoesquerda = no.noesquerda.valor
	}
	
	if no.nodireita != nil {
		valornodireita = no.nodireita.valor
	}

	println(no.valor, "\t", valornoesquerda, "\t", valornodireita)
}

func insercao (valor int, noraiz *No) *No {
	if valor > noraiz.valor {
		var novonoraiz No = No{valor,nil,nil}
		novonoraiz.valor = valor
		if valor < noraiz.valor/2 {
			novonoraiz.noesquerda = noraiz
		} else {
			novonoraiz.nodireita = noraiz
		}
		imprimirno(novonoraiz)
		return &novonoraiz
	} else {
		if valor <= noraiz.valor/2 {
			if noraiz.noesquerda == nil {
				novonoesquerda := No{valor,nil,nil}
				noraiz.noesquerda = &novonoesquerda
				imprimirno(*noraiz)
				println("novo")
				imprimirno(novonoesquerda)
			} else if valor > noraiz.noesquerda.valor {
				novonoesquerda := No{valor, nil, nil}
				pivot := noraiz.noesquerda
				noraiz.noesquerda = &novonoesquerda
				if pivot.valor <= valor/2 {
					novonoesquerda.noesquerda = pivot
				} else {
					novonoesquerda.nodireita = pivot 
				}
				imprimirno(*noraiz)
				imprimirno(novonoesquerda)
			} else {
				insercao(valor, noraiz.noesquerda)
			}
		} else {
			if noraiz.nodireita == nil {
				novonodireita := No{valor,nil,nil}
				noraiz.nodireita = &novonodireita
				imprimirno(*noraiz)
				println("novo")
				imprimirno(novonodireita)
			} else if valor > noraiz.nodireita.valor {
				novonodireita := No{valor, nil, nil}
				pivot := noraiz.nodireita
				noraiz.nodireita = &novonodireita
				if pivot.valor <= valor/2 {
					novonodireita.noesquerda = pivot
				} else {
					novonodireita.nodireita = pivot 
				}
				imprimirno(*noraiz)
				imprimirno(novonodireita)
			} else {
				insercao(valor, noraiz.nodireita)
			}
		}
		return noraiz
	} 
}

func main(){
	noraiz := No{100,nil,nil}
	noraiz = *insercao(45, &noraiz)
	noraiz = *insercao(21, &noraiz)
	noraiz = *insercao(11, &noraiz)
	noraiz = *insercao(7, &noraiz)
	noraiz = *insercao(2, &noraiz)
	noraiz = *insercao(5, &noraiz)
	noraiz = *insercao(10, &noraiz)
	noraiz = *insercao(38, &noraiz)
	noraiz = *insercao(25, &noraiz)
	noraiz = *insercao(68, &noraiz)
	noraiz = *insercao(60, &noraiz)
	noraiz = *insercao(51, &noraiz)
	buscar(5, &noraiz)
}

func buscar(valor int, noraiz *No) *No {
	if noraiz == nil {
		println("nulo")
		return nil
	}
	if valor == noraiz.valor {
		println("encontrei")
		return noraiz
	} else if valor <= noraiz.valor/2 {
		println(noraiz.valor)
		return buscar(valor, noraiz.noesquerda)
	} else {
		println(noraiz.valor)
		return buscar(valor, noraiz.nodireita)
	}
}