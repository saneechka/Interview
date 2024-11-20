# Structures in Go

## Introduction
Structs in Go are collections of fields that are used to group together data that is meant to be treated as a single unit. They are similar to classes in other languages but do not support inheritance.

## Defining Structs
Structs are defined using the `type` keyword, followed by the name of the struct and the `struct` keyword.

Example struct definitions can be found in `structs.go`.

## Accessing and Modifying Fields
Fields of a struct can be accessed and modified using the dot notation.

For examples of how to access and modify struct fields, see `structs.go`.

## Pointers to Structs and the new() Function
Pointers to structs are used to access struct fields without copying the struct. This can be more efficient for large structs.

Go provides the `new()` function to create a pointer to a new zero value of any type, including structs. When you use `new()` with a struct, Go allocates memory for a new instance of that struct, initializes it to the zero value, and returns a pointer to it.

Example can be found in `structs.go`.

## Embedding Pointers to Structs
In Go, you can embed pointers to structs within other structs. This allows for the creation of more complex data structures where one struct can inherit the fields and methods of another.

For examples of embedding pointers to structs, see `structs.go`.

## Methods
Methods in Go are functions that are associated with a struct type. They can be defined with a receiver type that is a struct.

For method examples, refer to `structs.go`.

## Anonymous and Nested Structs
Go supports anonymous structs, which do not require a named type, and nested structs, where one struct contains another struct as a field.

Examples of anonymous and nested structs are in `structs.go`.

## Tags
Struct fields can have tags, which are string literals that provide metadata about the field. These are often used with encoding/decoding libraries.

Examples of using struct tags can be found in `structs.go`.


# Структуры в Go

## Введение
Структуры в Go — это собрание полей, которые используются для группировки данных, предназначенных для обработки как единая сущность. Они похожи на классы в других языках, но не поддерживают наследование.

## Определение Структур
Структуры определяются с использованием ключевого слова `type`, за которым следует имя структуры и ключевое слово `struct`.

Примеры определения структур можно найти в `structs.go`.

## Доступ и Модификация Полей
Поля структуры доступны и могут быть изменены с использованием точечной нотации.

Примеры доступа и изменения полей структуры смотрите в `structs.go`.


## Указатели на Структуры и Функция new()
Указатели на структуры используются для доступа к полям структуры без копирования самой структуры. Это может быть более эффективным для больших структур.

В Go функция `new()` позволяет создать указатель на новый экземпляр любого типа со значением по умолчанию, включая структуры. При использовании `new()` со структурой, Go выделяет память для нового экземпляра этой структуры, инициализирует его значением по умолчанию и возвращает указатель на него.

Примеры использования функции `new()` со структурами можно найти в файле `structs.go`.

## Встраивание Указателей на Структуры
В Go можно встраивать указатели на структуры в другие структуры. Это позволяет создавать более сложные структуры данных, где одна структура может наследовать поля и методы другой.

Примеры встраивания указателей на структуры смотрите в файле `structs.go`.


## Методы
Методы в Go — это функции, которые ассоциированы с типом структуры. Они могут быть определены с типом приемника, который является структурой.

Примеры методов смотрите в `structs.go`.

## Анонимные и Вложенные Структуры
Go поддерживает анонимные структуры, которые не требуют именованного типа, и вложенные структуры, где одна структура содержит другую структуру в качестве поля.

Примеры анонимных и вложенных структур находятся в `structs.go`.

## Теги
Поля структуры могут иметь теги, которые являются строковыми литералами, предоставляющими метаданные о поле. Они часто используют

