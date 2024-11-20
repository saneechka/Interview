# Error Handling in Go

Error handling in Go is accomplished by returning errors as an explicit value. Errors are considered a normal outcome of a function, not as exceptional occurrences.

## Returning an Error

Functions in Go can return an error value along with other return values. If a function can cause an error, it should return the error as its last value.

## Checking for Errors

After calling a function, you should check if it returned an error. This is done by checking the error value against `nil`.

## Error Chains

As of Go 1.13, error chains can be processed using `errors.As` and `errors.Is` functions. This allows for checking if an error contains a certain type or value.

## Custom Error Types

You can create custom error types by implementing the `Error()` method of the `error` interface. This is useful for more precise error handling.

## Error Handling Recommendations

- Explicit error checking is mandatory.
- Define custom error types for complex situations.
- Use `errors.Is` and `errors.As` for working with error chains.

For examples of error handling, see `errors-examples.go`.


# Обработка ошибок в Go

Обработка ошибок в Go предполагает, что функции возвращают ошибку в качестве одного из возвращаемых значений. Это подход, при котором ошибки считаются нормальным результатом работы функции, а не исключительным событием.

## Возвращение ошибки

Функция может возвращать ошибку вместе с другими значениями. Если есть возможность возникновения ошибки, функция должна возвращать ошибку как последнее значение.

## Проверка на ошибку

После вызова функции следует проверить, не вернул ли она ошибку. Это делается проверкой значения ошибки на `nil`.

## Цепочки ошибок

С версии Go 1.13 можно обрабатывать цепочки ошибок с помощью функций `errors.As` и `errors.Is`. Это позволяет проверить, содержит ли ошибка определенный тип или значение.

## Создание пользовательских ошибок

Можно создать собственный тип ошибки, реализовав метод `Error()` интерфейса `error`. Это полезно для более точной обработки ошибок.

## Обработка ошибок

- Явная проверка на ошибку обязательна.
- Создавайте пользовательские типы ошибок для сложных ситуаций.
- Используйте `errors.Is` и `errors.As` для работы с цепочками ошибок.

Примеры обработки ошибок можно найти в файле `errors-examples.go`.

