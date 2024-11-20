# Constants in Go

## Introduction
Constants are immutable values that are known at compile time and do not change for the life of the program. In Go, constants are declared like variables, but with the `const` keyword.

## Declaring Constants
Constants can be character, string, boolean, or numeric values. Constants are declared with the `const` keyword, which can be used in place of the `var` keyword when the value does not change.

For examples of constant declarations, see `const.go`.

## Typed and Untyped Constants
Constants in Go can be typed or untyped. Typed constants work like immutable variables, can only be combined with same type values, and are subject to type conversion rules. Untyped constants, on the other hand, are flexible and can be used in any context that requires a value of their kind, without the need for explicit conversion.

## Enumerated Constants
Go does not have an enum type, but it supports enumerated constants with the `iota` enumerator. `iota` starts at 0 in each `const` block and increments by one for each constant declared.

For examples of enumerated constants using `iota`, refer to `const.go`.

## Benefits of Using Constants
Using constants has several benefits:
- It prevents changes to the value.
- It gives meaning to otherwise obscure literal values.
- It can improve performance, as the compiler can optimize the code.

For practical uses of constants and their benefits, see `const.go`.


# Константы в Go

## Введение
Константы — это неизменяемые (иммутабельные) значения, которые известны во время компиляции и не изменяются в течение жизни программы. В Go константы объявляются подобно переменным, но с использованием ключевого слова `const`.

## Объявление констант
Константы могут быть символьными, строковыми, булевыми или числовыми значениями. Константы объявляются с помощью ключевого слова `const`, которое используется вместо `var`.
Примеры объявления констант см. в `const.go`.

## Типизированные и нетипизированные константы
Константы в Go могут быть типизированными или нетипизированными. Типизированные константы действуют как неизменяемые переменные, могут быть соединены только с значениями того же типа и подчиняются правилам преобразования типов. Нетипизированные константы, в свою очередь, являются гибкими и могут использоваться в любом контексте, который требует значения их вида, без необходимости явного преобразования.

## Перечисляемые константы
Go не имеет типа enum, но поддерживает перечисляемые константы с помощью перечислителя `iota`. `iota` начинается с 0 в каждом блоке `const` и увеличивается на один с каждой объявленной константой.

Примеры перечисляемых констант с использованием `iota` см. в `const.go`.

## Преимущества использования констант
Использование констант имеет несколько преимуществ:
- Это предотвращает изменения значения.
- Это придаёт смысл иначе непонятным литеральным значениям.
- Это может улучшить производительность, так как компилятор может оптимизировать код.

