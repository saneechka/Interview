# Loops and Conditions in Go

## Introduction
Loops and conditional statements are fundamental for controlling the flow of execution in a program. In Go, `for` is the only looping construct, and it can be adapted for a variety of use cases.

## For Loop
The `for` loop in Go can work like a traditional C-style for-loop, a while-loop, or an infinite loop.

Examples of `for` loop usage can be seen in `loops-n-conditions.go`.

## If Statement
The `if` statement in Go is used to test a condition and execute a block of code if the condition is true.

For `if` statement examples, refer to `loops-n-conditions.go`.

## Switch Statement
The `switch` statement in Go is a multi-way branch statement that allows a variable to be tested for equality against a list of values.

See `loops-n-conditions.go` for switch statement examples.

## Nested Loops and Conditions
Go allows for loops and conditions to be nested inside one another for more complex control flow.

Nested loops and conditions can be found in `loops-n-conditions.go`.

## Range Loop
The `range` form of the `for` loop iterates over a slice or a map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

Examples of `range` loop can be found in `loops-n-conditions.go`.

## Breaking Out of Loops
The `break` statement in Go is used to terminate the innermost loop. The `continue` statement skips the current iteration and continues with the next iteration.

Usage of `break` and `continue` is demonstrated in `loops-n-conditions.go`.


# Циклы и Условные Операторы в Go

## Введение
Циклы и условные операторы являются основой для управления потоком выполнения программы. В Go `for` — это единственная конструкция цикла, которая может быть адаптирована для различных случаев использования.

## Цикл For
Цикл `for` в Go может функционировать как традиционный C-подобный for-цикл, while-цикл или бесконечный цикл.

Примеры в `loops-n-conditions.go`.

## Оператор If
Оператор `if` в Go используется для проверки условия и выполнения блока кода, если условие истинно.

Примеры в `loops-n-conditions.go`.

## Оператор Switch
Оператор `switch` в Go — это оператор многозначительного ветвления, который позволяет проверить переменную на равенство с списком значений.

Примеры оператора `switch` - `loops-n-conditions.go`.

## Вложенные Циклы и Условия
Go позволяет вкладывать циклы и условные операторы друг в друга для более сложного управления потоком.

Примеры вложенных циклов - `loops-n-conditions.go`.

## Цикл Range
Цикл `range` в `for` используется для итерации по элементам структур данных, таких как слайсы, массивы, строки и мапы. При итерации по слайсу или массиву в каждой итерации возвращаются два значения: индекс и копия элемента по этому индексу. По итерации по мапе нет строго определенного порядка ключей. При итерации по строке, итерация будет происходить по символам, а не по байтам. 

Примеры использования цикла `range` см. в `loops-n-conditions.go`.

## Выход из Циклов
Оператор `break` в Go используется для завершения самого внутреннего цикла. Оператор `continue` пропускает текущую итерацию и продолжает со следующей.

Использование `break` и `continue` демонстрируется в `loops-n-conditions.go`.

