#Prepare um algoritmo que calcule o valor de H, sendo que ele é determinado pela série H = 1/1 + 3/2 + 5/3 + 7/4 + ... + 99/50

H = 0.0
m = 50.0
i = 1

while i <= m:
    numerador = 1 + 2*(i-1)
    denominador = i
    H = H + numerador/denominador
    i+=1

print(f"{H}")