## Object-Oriented Programming in Go

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects", which can contain data and code: data in the form of fields (often known as attributes), and code, in the form of procedures (often known as methods).

Go is not a traditional object-oriented language. It does not support classes and inheritance as most OOP languages do. However, Go has types and methods, and it allows polymorphism and encapsulation, which enables it to achieve many of the same goals as OOP.

### Encapsulation
In Go, encapsulation is achieved through the use of packages and visibility rules. If a type or a function starts with an uppercase letter, it is exported (visible outside the package). If it starts with a lowercase letter, it is unexported (not visible outside the package).

### Polymorphism
Go supports polymorphism through interfaces. An interface is a type defined by a set of method signatures. A Go type implements an interface by implementing the methods. There is no explicit declaration of intent; if the methods are implemented, the type satisfies the interface.

### Composition over Inheritance
Go does not have classes or inheritance. It uses composition instead. You can embed one type within another to extend its behavior.


## Объектно-ориентированное программирование в Go

Объектно-ориентированное программирование (ООП) — это парадигма программирования, основанная на концепции "объектов", которые могут содержать данные в форме полей (часто известных как атрибуты) и код в форме процедур (часто известных как методы).

Go не является традиционным объектно-ориентированным языком. Он не поддерживает классы и наследование, как большинство языков ООП. Однако в Go есть типы и методы, и он поддерживает полиморфизм и инкапсуляцию, что позволяет достигать многих целей ООП.

### Инкапсуляция
В Go инкапсуляция достигается с помощью пакетов и правил видимости. Если тип или функция начинаются с заглавной буквы, они экспортируются (видны за пределами пакета). Если с маленькой буквы — не экспортируются (не видны за пределами пакета).

### Полиморфизм
Go поддерживает полиморфизм через интерфейсы. Интерфейс — это тип, определённый набором сигнатур методов. Тип в Go реализует интерфейс, реализуя его методы. Явного объявления намерений нет; если методы реализованы, тип удовлетворяет интерфейсу.

### Композиция вместо наследования
В Go нет классов или наследования. Вместо этого используется композиция. Можно встраивать один тип в другой для расширения его поведения.

