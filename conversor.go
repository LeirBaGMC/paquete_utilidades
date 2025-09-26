package utilidades

/*El usuario ingresará un valor en dólares y la moneda a la que quiere convertir y
el programa mostrará ese valor en dicha moneda. Las monedas permitidas son:
 Euros
LB (Libras Esterlinas)
Won (Sur Koreano)
BTC
*/
// Conversor... convierte un valor en dólares a la moneda indicada
// Opción: 1=EUR, 2=LB, 3=WON, 4=BTC
func Conversor(valor float64, opcion int) float64 {

	switch opcion {
	case 1:
		return valor * 0.86
	case 2:
		return valor * 0.75
	case 3:
		return valor * 1409.91
	case 4:
		return valor * 0.0000092
	default:
		return 0
	}

}
