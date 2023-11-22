#Prepare um algoritmo que calcule o valor de S, em que S = 1/1 - 2/4 + 3/9 - 4/16 + 5/25 - 6/36...-10/100

S = 0.0
i = 0
num = i

while i < 10:
    num = (i+1)*(-1)**i
    deno = num**2
    S = S + num/deno
    i+=1

print(f"{S}")