a = int(input("Entre com o lado 1: "))
b = int(input("Entre com o lado 2: "))
c = int(input("Entre com o lado 3: "))

if a < (b + c) and b < (a + c) and c < (a + b):
    print("É um triângulo!")
    if a == b and a == c and b == c:
        print("É um triângulo equilátero!")
    elif (a == b and a!=c) or (b == c and b != a) or (a == c and b != c): 
        print ("É um triângulo isósceles!")
    else: print("É um triângulo escaleno!")
else: print("Não é um triângulo!")