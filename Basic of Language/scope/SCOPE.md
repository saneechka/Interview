## Scope in Go

### Local Scope
A variable declared inside a function or a block is only accessible within that function or block and cannot be accessed from outside.

### Global Scope
A variable declared outside of all functions is global and can be accessed from any function within the same package.

### Block Scope
A variable declared within braces `{}` is limited to that block and cannot be accessed from outside of it.

## Scope of Structs in Go

### Package Level Structs
Structs declared at the package level are accessible within the entire package. If the name of the struct starts with a capital letter, it is exported and can be accessed from other packages.

### Struct Fields Visibility
The visibility of struct fields follows the same naming conventions:
- Fields starting with a capital letter are exported and can be accessed from other packages.
- Fields starting with a lowercase letter are unexported and can only be accessed from within the same package.


## Область Видимости в Go

### Локальная Область
Переменная, объявленная внутри функции или блока кода `{}`, доступна только внутри этой функции или блока и не может быть доступна снаружи.

### Глобальная Область
Переменная, объявленная вне всех функций, является глобальной и может быть доступна из любой функции в том же пакете.

### Область Блока
Переменная, объявленная внутри фигурных скобок `{}`, ограничена этим блоком и не может быть доступна снаружи.

## Область Видимости Структур в Go

### Структуры на Уровне Пакета
Структуры, объявленные на уровне пакета, доступны в пределах всего пакета. Если имя структуры начинается с заглавной буквы, она экспортируема и может быть доступна из других пакетов.

### Видимость Полей Структур
Видимость полей структур подчиняется тем же правилам именования:
- Поля, начинающиеся с заглавной буквы, экспортируемы и могут быть доступны из других пакетов.
- Поля, начинающиеся с маленькой буквы, не экспортируемы и доступны только внутри того же пакета.


