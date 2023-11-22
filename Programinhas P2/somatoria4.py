#Escreva um algoritmo que calcule e escreve a soma dos dez primeiros termos da seguionte série: 2/500 - 5/450 + 2/400 - 5/350 + ...

soma = 0
i = 0
vetor = []
while i <= 9:
    if i%2 != 0:
        num = -5
    else:
        num = 2
    valor = num/(500-50*i)
    vetor.append(valor)
    soma = soma + valor
    i+=1
print(vetor)
print(soma)

#Escreva um algoritmo que calcule e escreve a soma dos dez primeiros termos da seguionte série: 2/500 - 5/450 + 2/400 - 5/350 + ...

soma = 0
i = 0
vetor = []
while i <= 9:
    valor = (-6+(8**((i+1)%2)))/(500-50*i)
    vetor.append(valor)
    soma = soma + valor
    i+=1
print(vetor)
print(soma)