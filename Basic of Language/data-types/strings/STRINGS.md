# Strings in Go

## Introduction
A string in Go is a sequence of bytes that represents text. Strings are immutable, which means once a string is created, it is not possible to change its contents.

## String Initialization
Strings can be initialized in different ways:
- Directly with double quotes `"Hello"`
- With backticks for multi-line strings or to include special characters without escaping `\n`
- Using the `+` operator for concatenation

For string initialization examples, see `strings.go`.

## Common Operations
Some common operations on strings include:
- Getting the length of a string with `len()`
- Accessing individual bytes or characters
- Comparing strings
- Searching and replacing substrings

For more on string operations, refer to `strings.go`.

## Strings and Runes
In Go, strings are made up of bytes, but they can also be thought of as a sequence of runes. A rune is an integer that represents a Unicode code point.

## String Iteration
To iterate over each character of a string, ranging over the string will give you access to each rune.

To see string iteration in action, check `strings.go`.

## Converting to and from Strings
The `strconv` package provides functions for converting to and from strings.

Conversion examples can be found in `strings.go`.


# Строки в Go

## Введение
Строка в Go — это последовательность байтов, представляющих текст. Строки неизменяемы, то есть после создания строки изменить её содержимое невозможно.

## Инициализация строк
Строки можно инициализировать различными способами:
- Непосредственно с использованием двойных кавычек `"Hello"`
- С помощью обратных кавычек для многострочных строк или включения специальных символов без экранирования `\n`
- Используя оператор `+` для конкатенации

Примеры в `strings.go`.

## Распространенные операции
Некоторые распространенные операции со строками включают:
- Получение длины строки с помощью `len()`
- Доступ к отдельным байтам или символам
- Сравнение строк
- Поиск и замена подстрок

См. `strings.go`.

## Строки и руны
В Go строки состоят из байтов, но их также можно рассматривать как последовательность рун. Руна — это целое число, представляющее кодовую точку Unicode.

## Итерация по строкам
Для итерации по каждому символу строки можно использовать цикл `range`, который предоставит доступ к каждой руне.

Примеры в `strings.go`.

## Преобразование в строки и из строк
Пакет `strconv` предоставляет функции для преобразования в строки и из строк.

Примеры в `strings.go`.

