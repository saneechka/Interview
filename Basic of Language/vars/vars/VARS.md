# Variables in Go

## Introduction
Variables are a fundamental aspect of any programming language, serving as named storage that our programs can manipulate. Each variable in Go is associated with a specific data type, which determines the size and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

## Default Zero Values
In Go, variables declared without an explicit initial value are given their zero value. The zero value is dependent on the type:

- `int`: 0
- `float64`: 0.0
- `bool`: false
- `string`: "" (the empty string)
- `pointer`: nil

## Declaration
For examples of variable declaration, please refer to the `vars.go` file within this directory. This file includes various ways to declare and initialize variables, showcasing the syntax and best practices of variable usage in Go.

## Unused Variables
The Go compiler enforces that every declared variable must be used in a program. Declaring a variable without using it results in a compilation error, ensuring that the code remains clean and free from clutter. To see how to handle unused variables, refer to the examples in `vars.go`.

## Scope
Variables in Go have a certain scope within which they are accessible. The scope is determined by the location where the variable is declared. Explore the `vars.go` file to see examples of different variable scopes.

## Naming Conventions
Go has a few conventions regarding variable names:
- Names can begin with a letter or an underscore, followed by any combination of letters, digits, or underscores.
- Go is case-sensitive: `varName` and `VarName` are different variables.
- By convention, variable names should be written in camelCase for local variables and in PascalCase for exported variables.

For more details on variable naming and conventions in Go, see the `vars.go` file.


# (РУ) Переменные в Go

## Введение
Переменные — это основа любого языка программирования. Они представляют собой именованные области памяти, которыми можно управлять в процессе выполнения программы. В Go каждая переменная связана с определенным типом, что определяет, какой объем памяти она занимает, какие значения может принимать и какие операции с ней можно выполнять.

## Нулевые значения по умолчанию
В Go переменные, объявленные без явного начального значения, принимают нулевое значение. Нулевое значение зависит от типа:

- `int`: 0
- `float64`: 0.0
- `bool`: false
- `string`: "" (пустая строка)
- `pointer`: nil

## Объявление
Примеры объявления переменных в Go можно посмотреть в файле `vars.go`. Этот файл включает различные способы объявления и инициализации переменны.

## Неиспользуемые переменные
Компилятор Go требует, чтобы каждая объявленная переменная была использована в программе. Объявление переменной без её использования приводит к ошибке компиляции, что обеспечивает чистоту и свободу кода от лишнего. Как обойти требования компилятора для неиспользуемых переменных, см. в `vars.go`.

## Область видимости
Переменные в Go имеют определённую область видимости, в пределах которой они доступны. Область видимости определяется местом объявления переменной. Примеры - `vars.go`.

## Конвенции именования
В Go существуют определённые конвенции касательно имен переменных:
- Имена могут начинаться с буквы или подчёркивания, за которыми следуют любые комбинации букв, цифр или подчёркиваний.
- Go чувствителен к регистру: `varName` и `VarName` - это разные переменные.
- По конвенции, имена переменных должны быть написаны в стиле camelCase для локальных переменных и в стиле PascalCase для экспортируемых переменных.

