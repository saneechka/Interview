# Data Types in Go

## Introduction
Data types in Go specify the type of data that a variable can hold. From simple built-in types to complex user-defined types, Go offers a range of options to handle data efficiently.

## Basic Types
Go's basic types include:

- `bool`: A boolean value that is either `true` or `false`.
- `string`: A sequence of characters representing textual data.
- Numeric types:
  - `int`, `int8`, `int16`, `int32`, `int64`: Integer types with varying sizes.
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`: Unsigned integer types.
  - `float32`, `float64`: Floating-point numbers.
  - `complex64`, `complex128`: Complex numbers.

For examples of variable declarations with different data types, see `data-types.go`.

## Composite Types
Go also supports composite types, such as:

- `array`: Fixed-size list of elements of the same type.
- `slice`: A dynamically-sized, flexible view into the elements of an array.
- `struct`: A collection of fields.
- `map`: A collection of key-value pairs.
- `channel`: A conduit through which you can send and receive values with the channel operator.

For usage examples of composite data types, refer to `data-types.go`.

## Type Conversion
In Go, type conversion is explicit. This means that you need to convert variables to another type using a conversion expression if they are not of the same type.

## Type Aliases
Go allows the creation of type aliases, enabling programmers to give a new name to an existing type.

For more information on type conversion and type aliases, see `data-types.go`.


# Типы данных в Go

## Введение
В Go каждая переменная связана с типом данных, который определяет, какого рода значения она может хранить. Язык предлагает разнообразные типы данных: от базовых, встроенных в систему, до сложных (составных), определяемых пользователем, что позволяет разработчикам выбирать наиболее подходящий способ представления и обработки информации.

## Базовые типы
К базовым типам Go относятся:

- `bool`: Логическое значение, которое может быть `true` или `false`.
- `string`: Последовательность символов, представляющая текстовые данные.
- `rune`: int32.
- `byte`: uint8.
- Числовые типы:
  - `int`, `int8`, `int16`, `int32`, `int64`: Целочисленные типы различного размера.
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`: Беззнаковые целочисленные типы.
  - `float32`, `float64`: Числа с плавающей запятой. Точность ~6 знаков после запятой, и 15. float32 - быстро копятся ошибки
  - `complex64`, `complex128`: Комплексные числа.

Примеры объявления переменных с различными типами данных см. в `data-types.go`.

## Составные типы
Go также поддерживает составные типы, такие как:

- `array`: Список элементов фиксированного размера одного типа.
- `slice`: Динамически изменяемый, гибкий срез элементов массива.
- `struct`: Коллекция полей.
- `map`: Коллекция пар ключ-значение.
- `channel`: Канал, через который можно отправлять и получать значения с помощью оператора канала.

Примеры использования составных типов данных см. в `data-types.go`.

## Преобразование типов
В Go преобразование типов является явным. Это означает, что для преобразования переменных в другой тип, если они не одного типа, необходимо использовать выражение преобразования.


package main

import (
"fmt"
)

func main() {
var x []int
var y []int
var z []int

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y = append(x, 4)
	z = append(x, 5)

	fmt.Println(y, z)
}
