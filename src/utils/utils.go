package utils

import (
	"errors"
	"strconv"
	"strings"
)

func IsValidCPF(cpf string) error {
	// Remova caracteres não numéricos
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifique o tamanho e caracteres
	if len(cpf) != 11 {
		return errors.New("cpf deve ter 11 dígitos")
	}

	for _, char := range cpf {
		if char < '0' || char > '9' {
			return errors.New("cpf deve conter apenas dígitos")
		}
	}

	// Verifique se todos os dígitos são iguais
	if allEqual(cpf) {
		return errors.New("cpf inválido")
	}

	// Calcule os dígitos verificadores
	firstDigit := calculateDigit(cpf[:9])
	secondDigit := calculateDigit(cpf[:10])

	// Verifique se os dígitos verificadores são iguais aos fornecidos
	if cpf[9] != firstDigit || cpf[10] != secondDigit {
		return errors.New("cpf inválido")
	}

	return nil
}

func allEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func calculateDigit(cpf string) byte {
	sum := 0
	multiplier := len(cpf) + 1

	for _, char := range cpf {
		digit, _ := strconv.Atoi(string(char))
		sum += digit * multiplier
		multiplier--
	}

	remainder := sum % 11
	if remainder < 2 {
		return '0'
	}

	return byte(11 - remainder + '0')
}