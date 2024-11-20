# Arrays in Go

## Introduction
Arrays are a collection of elements of the same type, stored in contiguous memory locations. In Go, the size of an array is fixed at the time of declaration.

## Declaration and Initialization
Arrays can be declared and initialized in multiple ways:

- Declaring an array with a specified size: `var a [5]int`
- Initializing an array with elements: `b := [5]int{1, 2, 3, 4, 5}`
- Letting Go infer the array size: `c := [...]int{1, 2, 3, 4, 5}`

See `arrays.go` for examples of declaration and initialization.

## Accessing Elements
Elements of an array are accessed using the index, which starts at 0.

For examples of accessing elements, refer to `arrays.go`.

## Iterating Over Arrays
You can iterate over an array using a `for` loop or a `for range` loop.

Iteration examples are provided in `arrays.go`.

## Limitations and Alternatives
Arrays have a fixed size, which can be limiting. Slices are a more flexible alternative to arrays in Go.

For more on slices, see the `slices` section.

# Массивы в Go

## Введение
Массивы — это набор элементов одного типа, хранящихся в смежных ячейках памяти. В Go размер массива фиксирован при объявлении. По этой причине они не так популярны, как, например, срезы. 

## Объявление и инициализация
Массивы можно объявлять и инициализировать несколькими способами:

- Объявление массива с указанным размером: `var a [5]int`
- Инициализация массива элементами: `b := [5]int{1, 2, 3, 4, 5}`
- Позволяя Go определить размер массива: `c := [...]int{1, 2, 3, 4, 5}`

Примеры объявления и инициализации см. в `arrays.go`.

## Доступ к элементам
Доступ к элементам массива осуществляется по индексу, который начинается с 0.

См. `arrays.go`.

## Итерация по массивам
По массиву можно итерироваться, используя цикл `for` или цикл `for range`.
См. `arrays.go`.

## Ограничения и альтернативы
Массивы имеют фиксированный размер, что может быть ограничивающим. Срезы являются более гибкой альтернативой массивам в Go.

Для получения дополнительной информации о срезах см. раздел `slices`.

