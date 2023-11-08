package main

import "fmt"

func printarNaturais(n int) {
	if n > 50 {
		return
	}
	fmt.Println(n)
	printarNaturais(n + 1)
}
func contadorDigitos(n int) int {
	if n == 0 {
		fmt.Printf("A contagem foi de ")
		return 0
	}

	return 1 + contadorDigitos(n/10)
}
func encontreMax(arr []int, n int) int {
	if n == 1 {
		return arr[0]
	}
	return max(encontreMax(arr, n-1), arr[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func mmcTeste(a, b int) int {
	if b == 0 {
		return a
	}
	return mmcTeste(b, a%b)
}

func mmc(a, b int) int {
	return (a * b) / mmcTeste(a, b)
}
func hailstone(hailston int) {
	if hailston == 1 {
		fmt.Println(hailston)
		return
	}
	fmt.Println(hailston)
	if hailston%2 == 0 {
		hailstone(hailston / 2)
	} else {
		hailstone(3*hailston + 1)
	}
}
func menu(arr []int) int {
	opcao := 0
	hailston := 0
	for {
		fmt.Println("\n")
		fmt.Println("_______INICIO_DO_PROGRAMA___________")
		fmt.Println("1 - NumerosNaturais ")
		fmt.Println("2 - Contador de Digitos")
		fmt.Println("3 - Encontrar o Numero Maximo em uma array")
		fmt.Println("4 - MMC ")
		fmt.Println("5 - Seqüência de Hailstone ")
		fmt.Println("0 - Sair ")
		fmt.Print("Digite: ")
		fmt.Scan(&opcao)
		fmt.Println("\n")

		switch opcao {
		case 1:
			fmt.Println("-------------------")
			fmt.Println("1 - Nº naturais ")
			printarNaturais(1)
			fmt.Println("-------------------")
		case 2:
			fmt.Println("-------------------")
			fmt.Println("2 - Contador de Digitos ")
			fmt.Println(contadorDigitos(123456789))
			fmt.Println("-------------------")
		case 3:
			fmt.Println("-------------------")
			fmt.Println("3 - Encontrar o Numero Maximo em uma array ")
			fmt.Println("A sequencia do array foi", arr)
			print("o maior numero encontrado foi: ")
			fmt.Println(encontreMax(arr, len(arr)))
			fmt.Println("-------------------")

		case 4:
			fmt.Println("-------------------")
			fmt.Println("4 - MMC ")
			print("O MMC de 15 e 20 é: ")
			fmt.Println(mmc(15, 20))

			fmt.Println("-------------------")

		case 5:
			fmt.Println("-------------------")
			fmt.Println("5 - Seqüência de Hailstone ")
			print("Digite o numero para calcular ")
			fmt.Scan(&hailston)
			fmt.Println("Digite o numero calculado foi ",hailston)
			hailstone(hailston)

			fmt.Println("-------------------")

		case 0:
			fmt.Println("Saindo do Programa")

			fmt.Println("-------------------")
			fmt.Println("__PROGRAMA_EM_GOlang_DESENVOLVIDO_POR_ANDREWDEVOS-72409_____")
			fmt.Println("-------------------")
			return opcao
		}
	}

}

func main() {
	arr := []int{1, 4, 5, 3, 2, 8, 6}

	menu(arr)

}
