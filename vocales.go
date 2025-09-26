package utilidades

// contadorVocales recibe una palabra y devuelve el número total de vocales
func contadorVocales(palabra string) int {
	contador := 0

	for _, letra := range palabra {
		if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' ||
			letra == 'A' || letra == 'E' || letra == 'I' || letra == 'O' || letra == 'U' {
			contador += 1
		}
	}

	return contador
}
