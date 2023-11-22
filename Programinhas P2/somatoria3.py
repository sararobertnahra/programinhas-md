#Calcule um algoritmo que calcule o valor dos dez primeiros termos da séria H, em que: H = 1/pot(1,3) - 1/pot(3,3) + 1/pot(5,3) - 1/pot(7,3) + 1/pot(9,3) - 1/pot(11,3)+ ...

soma = 0
i = 1

while i <= 19:
    if i % 4 == 1:
        soma += 1 / (i ** 3)
    else:
        soma -= 1 / (i ** 3)
    i += 2

print(f'O valor dos dez primeiros termos da série H é: {soma}')

