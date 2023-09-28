tipo_combustivel = input("Informe o combustível (G - gasolina | A - álcool):  ")
quantidade_combustivel = float(input("Entre com a quantidade de combustivel em litros: "))

if tipo_combustivel.lower() == "g":
    total_a_ser_pago = quantidade_combustivel * 3.30
    if (quantidade_combustivel <= 20.0 and tipo_combustivel.lower() == "g"):
        total_com_desconto = total_a_ser_pago - (total_a_ser_pago * 0.04)
    else: 
        total_com_desconto = total_a_ser_pago - (total_a_ser_pago * 0.06)

if tipo_combustivel.lower() == "a":
    total_a_ser_pago = quantidade_combustivel * 2.90
    if (quantidade_combustivel <= 20.0 and tipo_combustivel.lower() == "a"):
        total_com_desconto = total_a_ser_pago - (total_a_ser_pago * 0.03)
    else: 
        total_com_desconto = total_a_ser_pago - (total_a_ser_pago * 0.05)

print(f"Total da ser pago: {total_a_ser_pago}")

print(f"\nA qunatidade total a ser paga com desconto é {total_com_desconto}")